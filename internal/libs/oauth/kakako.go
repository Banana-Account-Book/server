package oauth

import (
	"encoding/json"
	"fmt"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
)

type kakaoClient struct {
	config oauthConfig
}

func newKakaoClient() *kakaoClient {
	return &kakaoClient{config: oauthConfig{
		authURL:      config.Oauth.Kakao.BaseURL + "/authorize",
		tokenURL:     config.Oauth.Kakao.BaseURL + "/token",
		userInfoURL:  "https://kapi.kakao.com/v2/user/me",
		clientId:     config.Oauth.Kakao.ClientId,
		clientSecret: config.Oauth.Kakao.ClientSecret,
		redirectUri:  config.Oauth.Kakao.RedirectUri,
	}}
}

func (c *kakaoClient) parseUserInfo(body []byte) (*OauthInfo, error) {
	var result struct {
		KakaoAccount struct {
			Profile struct {
				Nickname        string `json:"nickname"`
				ProfileImageUrl string `json:"profile_image_url"`
			} `json:"profile"`
			Email string `json:"email"`
		} `json:"kakao_account"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Name:         result.KakaoAccount.Profile.Nickname,
		Email:        result.KakaoAccount.Email,
		ProfileImage: result.KakaoAccount.Profile.ProfileImageUrl,
	}

	return &userInfo, nil
}

func (c *kakaoClient) getConfig() oauthConfig {
	return c.config
}
