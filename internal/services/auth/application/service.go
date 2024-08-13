package application

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/oauth"
	userModel "banana-account-book.com/internal/services/users/domain"
	userInfra "banana-account-book.com/internal/services/users/infrastructure"
)

type AuthService struct {
	userRepository userInfra.UserRepository
	kakaoClient    *oauth.KakaoClient
}

func NewAuthService(userRepository userInfra.UserRepository, kakaoClient *oauth.KakaoClient) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		kakaoClient:    kakaoClient,
	}
}

func (s *AuthService) GetAuthUrl(provider string) (string, error) {
	if provider == "kakao" {
		return s.kakaoClient.GetUrl(), nil
	}

	return "", appError.New(httpCode.BadRequest, fmt.Sprintf("Invalid provider: %s", provider), "")

}

// TODO: refactoring
func (s *AuthService) OAuth(code, provider string) (string, error) {
	if provider == "kakao" {
		userInfo, err := s.kakaoClient.OAuth(code)
		if err != nil {
			return "", appError.Wrap(err)
		}

		user, ok, err := s.userRepository.FindByEmail(userInfo.Email)
		if err != nil {
			return "", appError.Wrap(err)
		}

		accessToken := ""

		if ok {
			token, err := user.EncodeAccessToken()
			if err != nil {
				return "", appError.Wrap(err)
			}
			accessToken = token
		} else {

			newUser, err := userModel.New(userInfo.Email, userInfo.Name, []string{"kakao"})
			if err != nil {
				return "", appError.Wrap(err)
			}

			if err := s.userRepository.Save(newUser); err != nil {
				return "", appError.Wrap(err)
			}
			token, err := newUser.EncodeAccessToken()

			if err != nil {
				return "", appError.Wrap(err)
			}
			accessToken = token
		}

		return accessToken, nil
	}

	return "", appError.New(httpCode.BadRequest, fmt.Sprintf("Invalid provider: %s", provider), "")

}
