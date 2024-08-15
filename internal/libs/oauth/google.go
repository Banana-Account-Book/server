package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
)

type GoogleClient struct {
	httpClient *http.Client
}

type googleResult struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func NewGoogleClient() *GoogleClient {
	return &GoogleClient{httpClient: &http.Client{}}
}

func (c *GoogleClient) GetUrl() string {
	baseURL := "https://accounts.google.com/o/oauth2/auth"
	params := url.Values{}
	params.Add("client_id", config.Oauth.Google.ClientId)
	params.Add("redirect_uri", config.Oauth.Google.RedirectUri)
	params.Add("response_type", "code")
	params.Add("scope", "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email")

	fullURL := baseURL + "?" + params.Encode()
	return fullURL
}

func (c GoogleClient) OAuth(code string) (*OauthInfo, error) {
	// token 가져오기
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.Oauth.Google.ClientId)
	data.Set("client_secret", config.Oauth.Google.ClientSecret)
	data.Set("redirect_uri", config.Oauth.Google.RedirectUri)
	data.Set("state", "banana-account-book")
	data.Set("code", code)

	res, err := c.httpClient.Post("https://oauth2.googleapis.com/token", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to request kakao oauth token, %v", err), "")
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("unexpected status code: %d, body: %s", res.StatusCode, string(body)), "")
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	json.NewDecoder(res.Body).Decode(&result)

	// user 정보 가져오기
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to create request, %v", err), "")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))

	res, err = c.httpClient.Do(req)

	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to request kakao user info, %v", err), "")
	}
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("unexpected status code: %d, body: %s", res.StatusCode, string(body)), "")
	}

	var r googleResult

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Email:        r.Email,
		Name:         r.Name,
		ProfileImage: r.Picture,
	}

	return &userInfo, nil
}
