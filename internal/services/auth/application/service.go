package application

import (
	"time"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/jwt"
	"banana-account-book.com/internal/libs/oauth"
	"banana-account-book.com/internal/services/auth/dto"
	userModel "banana-account-book.com/internal/services/users/domain"
	userInfra "banana-account-book.com/internal/services/users/infrastructure"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository userInfra.UserRepository
	oauthProvider  *oauth.OAuthProvider
	db             *gorm.DB
}

func NewAuthService(
	userRepository userInfra.UserRepository,
	oauthProvider *oauth.OAuthProvider,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		oauthProvider:  oauthProvider,
		db:             db,
	}
}

func (s *AuthService) GetAuthUrl(provider string) (string, error) {
	url, err := s.oauthProvider.GetUrl(provider)
	if err != nil {
		return "", appError.Wrap(err)
	}

	return url, nil
}

func (s *AuthService) OAuth(code, provider string) (*dto.OauthResponse, error) {
	var (
		responseDto *dto.OauthResponse
		userInfo    *oauth.OauthInfo
		err         error
	)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		userInfo, err = s.oauthProvider.OAuth(provider, code)
		if err != nil {
			return appError.Wrap(err)
		}

		// 유저 정보를 찾고, 없다면 생성 후 토큰 발행
		responseDto, err = s.generateAccessToken(tx, userInfo, provider)
		if err != nil {
			return appError.Wrap(err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return responseDto, nil
}

func (s *AuthService) generateAccessToken(tx *gorm.DB, userInfo *oauth.OauthInfo, provider string) (*dto.OauthResponse, error) {
	sync := false
	user, exists, err := s.userRepository.FindByEmail(tx, userInfo.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		if !user.HasProvider(provider) {
			user.AddProvider(provider)
			sync = true
		}
	} else {
		user, err = userModel.New(userInfo.Email, userInfo.Name, []string{provider})
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := user.EncodeAccessToken()

	if err := s.userRepository.Save(tx, user); err != nil {
		return nil, err
	}

	return &dto.OauthResponse{AccessToken: accessToken, Sync: sync, ExpiredAt: time.Now().Add(jwt.AccessTokenExpiredAfter)}, err
}
