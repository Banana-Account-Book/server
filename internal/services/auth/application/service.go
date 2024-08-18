package application

import (
	"fmt"
	"time"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/jwt"
	"banana-account-book.com/internal/libs/oauth"
	accountModel "banana-account-book.com/internal/services/accountBooks/domain"
	accountBookInfra "banana-account-book.com/internal/services/accountBooks/infrastructure"
	dto "banana-account-book.com/internal/services/auth/dto/_provider"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	roleInfra "banana-account-book.com/internal/services/roles/infrastructure"
	userModel "banana-account-book.com/internal/services/users/domain"
	userInfra "banana-account-book.com/internal/services/users/infrastructure"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository        userInfra.UserRepository
	accountBookRepository accountBookInfra.AccountBookRepository
	roleRepository        roleInfra.RoleRepository
	oauthProvider         *oauth.OAuthProvider
	db                    *gorm.DB
}

func NewAuthService(
	userRepository userInfra.UserRepository,
	accountBookRepository accountBookInfra.AccountBookRepository,
	roleRepository roleInfra.RoleRepository,
	oauthProvider *oauth.OAuthProvider,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		userRepository:        userRepository,
		oauthProvider:         oauthProvider,
		db:                    db,
		accountBookRepository: accountBookRepository,
		roleRepository:        roleRepository,
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

		err = s.createAccountBook(tx, user)
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

// TODO: 이벤트 소싱으로 변경
func (s *AuthService) createAccountBook(tx *gorm.DB, user *userModel.User) error {
	accountBook, err := accountModel.New(user.Id, fmt.Sprintf("%s의 가계부", user.Name))

	if err != nil {
		return err
	}

	if err := s.accountBookRepository.Save(tx, accountBook); err != nil {
		return err
	}

	role, err := roleModel.New(user.Id, accountBook.Id, "owner")
	if err != nil {
		return err
	}

	if err := s.roleRepository.Save(tx, role); err != nil {
		return err
	}

	return err
}
