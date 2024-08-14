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

type KakaoClient struct {
	httpClient *http.Client
}

func NewKakaoClient() *KakaoClient {
	return &KakaoClient{httpClient: &http.Client{}}
}

type kakaoResult struct {
	KakaoAccount struct {
		Profile struct {
			Nickname        string `json:"nickname"`
			ProfileImageUrl string `json:"profile_image_url"`
		} `json:"profile"`
		Email string `json:"email"`
	} `json:"kakao_account"`
}

func (c *KakaoClient) GetUrl() string {
	baseURL := "https://kauth.kakao.com/oauth/authorize"
	params := url.Values{}
	params.Add("client_id", config.Oauth.Kakao.ClientId)
	params.Add("redirect_uri", config.Oauth.Kakao.RedirectUri)
	params.Add("response_type", "code")

	fullURL := baseURL + "?" + params.Encode()
	return fullURL
}

func (c *KakaoClient) OAuth(code string) (*OauthInfo, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.Oauth.Kakao.ClientId)
	data.Set("client_secret", config.Oauth.Kakao.ClientSecret)
	data.Set("redirect_uri", config.Oauth.Kakao.RedirectUri)
	data.Set("code", code)

	res, err := c.httpClient.Post("https://kauth.kakao.com/oauth/token", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
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

	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
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

	var r kakaoResult

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("failed to decode user info response: %v", err), "")
	}

	userInfo := OauthInfo{
		Name:         r.KakaoAccount.Profile.Nickname,
		Email:        r.KakaoAccount.Email,
		ProfileImage: r.KakaoAccount.Profile.ProfileImageUrl,
	}

	return &userInfo, nil
}
