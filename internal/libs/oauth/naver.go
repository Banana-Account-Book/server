package oauth

import (
	"encoding/json"
	"fmt"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
)

type NaverClient struct {
	config oauthConfig
}

func NewNaverClient() *NaverClient {
	return &NaverClient{config: oauthConfig{
		authURL:      config.Oauth.Naver.BaseURL + "/authorize",
		tokenURL:     config.Oauth.Naver.BaseURL + "/token",
		userInfoURL:  "https://openapi.naver.com/v1/nid/me",
		clientId:     config.Oauth.Naver.ClientId,
		clientSecret: config.Oauth.Naver.ClientSecret,
		redirectUri:  config.Oauth.Naver.RedirectUri,
		state:        "banana-account-book",
	}}
}

func (c *NaverClient) parseUserInfo(body []byte) (*OauthInfo, error) {
	var result struct {
		Response struct {
			Email        string `json:"email"`
			Nickname     string `json:"nickname"`
			ProfileImage string `json:"profile_image"`
		} `json:"response"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Name:         result.Response.Nickname,
		Email:        result.Response.Email,
		ProfileImage: result.Response.ProfileImage,
	}

	return &userInfo, nil
}

func (c *NaverClient) getConfig() oauthConfig {
	return c.config
}
