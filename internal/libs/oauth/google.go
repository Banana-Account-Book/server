package oauth

import (
	"encoding/json"
	"fmt"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
)

type googleClient struct {
	config oauthConfig
}

func newGoogleClient() *googleClient {
	return &googleClient{config: oauthConfig{
		authURL:      "https://accounts.google.com/o/oauth2/auth",
		tokenURL:     "https://oauth2.googleapis.com/token",
		userInfoURL:  "https://www.googleapis.com/oauth2/v2/userinfo",
		clientId:     config.Oauth.Google.ClientId,
		clientSecret: config.Oauth.Google.ClientSecret,
		redirectUri:  config.Oauth.Google.RedirectUri,
		state:        "banana-account-book",
		scope:        "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email",
	}}
}

func (c *googleClient) parseUserInfo(body []byte) (*OauthInfo, error) {
	var result struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Email:        result.Email,
		Name:         result.Name,
		ProfileImage: result.Picture,
	}

	return &userInfo, nil
}

func (c *googleClient) getConfig() oauthConfig {
	return c.config
}
