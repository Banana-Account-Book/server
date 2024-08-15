package application

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/oauth"
	userModel "banana-account-book.com/internal/services/users/domain"
	userInfra "banana-account-book.com/internal/services/users/infrastructure"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository userInfra.UserRepository
	kakaoClient    *oauth.KakaoClient
	naverClient    *oauth.NaverClient
	googleClient   *oauth.GoogleClient
	db             *gorm.DB
}

func NewAuthService(
	userRepository userInfra.UserRepository,
	kakaoClient *oauth.KakaoClient,
	naverClient *oauth.NaverClient,
	googleClient *oauth.GoogleClient,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		kakaoClient:    kakaoClient,
		naverClient:    naverClient,
		googleClient:   googleClient,
		db:             db,
	}
}

func (s *AuthService) GetAuthUrl(provider string) (string, error) {
	switch provider {
	case "kakao":
		return s.kakaoClient.GetUrl(), nil
	case "naver":
		return s.naverClient.GetUrl(), nil
	case "google":
		return s.googleClient.GetUrl(), nil
	default:
		message := fmt.Sprintf("Invalid provider: %s", provider)
		return "", appError.New(httpCode.BadRequest, message, message)
	}
}

func (s *AuthService) OAuth(code, provider string) (string, error) {
	var (
		accessToken string
		userInfo    *oauth.OauthInfo
		err         error
	)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// OAuth 정보를 가져오는 로직을 별도 함수로 분리
		userInfo, err = s.getOAuthInfo(provider, code)
		if err != nil {
			return appError.Wrap(err)
		}

		// 유저 정보를 찾고, 없다면 생성 후 토큰 발행
		accessToken, err = s.generateAccessToken(tx, userInfo)
		if err != nil {
			return appError.Wrap(err)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) getOAuthInfo(provider, code string) (*oauth.OauthInfo, error) {
	switch provider {
	case "kakao":
		return s.kakaoClient.OAuth(code)
	case "naver":
		return s.naverClient.OAuth(code)
	case "google":
		return s.googleClient.OAuth(code)
	default:
		message := fmt.Sprintf("Invalid provider: %s", provider)
		return nil, appError.New(httpCode.BadRequest, message, message)
	}
}

func (s *AuthService) generateAccessToken(tx *gorm.DB, userInfo *oauth.OauthInfo) (string, error) {
	user, exists, err := s.userRepository.FindByEmail(tx, userInfo.Email)
	if err != nil {
		return "", err
	}

	if exists {
		return user.EncodeAccessToken()
	}

	newUser, err := userModel.New(userInfo.Email, userInfo.Name, []string{"kakao"})
	if err != nil {
		return "", err
	}

	if err := s.userRepository.Save(tx, newUser); err != nil {
		return "", err
	}

	return newUser.EncodeAccessToken()
}
