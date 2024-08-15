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

type NaverClient struct {
	httpClient *http.Client
}

func NewNaverClient() *NaverClient {
	return &NaverClient{httpClient: &http.Client{}}
}

type naverResult struct {
	Response struct {
		Email        string `json:"email"`
		Nickname     string `json:"nickname"`
		ProfileImage string `json:"profile_image"`
	} `json:"response"`
}

func (c *NaverClient) GetUrl() string {
	baseURL := config.Oauth.Naver.BaseURL + "/authorize"
	params := url.Values{}
	params.Add("client_id", config.Oauth.Naver.ClientId)
	params.Add("redirect_uri", config.Oauth.Naver.RedirectUri)
	params.Add("response_type", "code")
	params.Add("state", "banana-account-book")

	fullURL := baseURL + "?" + params.Encode()
	return fullURL
}

func (c *NaverClient) OAuth(code string) (*OauthInfo, error) {
	// token 가져오기
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.Oauth.Naver.ClientId)
	data.Set("client_secret", config.Oauth.Naver.ClientSecret)
	data.Set("state", "banana-account-book")
	data.Set("code", code)

	res, err := c.httpClient.Post(config.Oauth.Naver.BaseURL+"/token", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
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
	req, err := http.NewRequest("GET", "https://openapi.naver.com/v1/nid/me", nil)
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

	var r naverResult

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Email:        r.Response.Email,
		Name:         r.Response.Nickname,
		ProfileImage: r.Response.ProfileImage,
	}

	return &userInfo, nil
}
