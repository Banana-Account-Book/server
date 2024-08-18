package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
)

type OauthInfo struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
	Provider     string `json:"provider"`
}

type oauthConfig struct {
	authURL      string
	tokenURL     string
	userInfoURL  string
	clientId     string
	clientSecret string
	redirectUri  string
	state        string
	scope        string
}

type oauthClient interface {
	parseUserInfo([]byte) (*OauthInfo, error)
	getConfig() oauthConfig
}
type OAuthProvider struct {
	clients    map[string]oauthClient
	httpClient *http.Client
}

func NewOAuthProvider() *OAuthProvider {
	return &OAuthProvider{
		clients: map[string]oauthClient{
			"google": newGoogleClient(),
			"kakao":  newKakaoClient(),
			"naver":  newNaverClient()},
		httpClient: &http.Client{},
	}
}

func (o *OAuthProvider) client(provider string) (oauthClient, error) {
	client, ok := o.clients[provider]
	if !ok {
		return nil, appError.New(httpCode.BadRequest, fmt.Sprintf("unknown provider: %s", provider), "")
	}
	return client, nil
}

func (o *OAuthProvider) GetUrl(provider string) (string, error) {
	client, err := o.client(provider)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("client_id", client.getConfig().clientId)
	params.Add("redirect_uri", client.getConfig().redirectUri+"?provider="+provider)
	params.Add("response_type", "code")
	params.Add("state", client.getConfig().state)
	params.Add("scope", client.getConfig().scope)

	return client.getConfig().authURL + "?" + params.Encode(), nil
}

func (o *OAuthProvider) OAuth(provider, code string) (*OauthInfo, error) {
	client, err := o.client(provider)
	if err != nil {
		return nil, err
	}
	// token 가져오기
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", client.getConfig().clientId)
	data.Set("client_secret", client.getConfig().clientSecret)
	data.Set("redirect_uri", client.getConfig().redirectUri)
	data.Set("state", client.getConfig().state)
	data.Set("code", code)

	res, err := o.httpClient.Post(client.getConfig().tokenURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
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
	req, err := http.NewRequest("GET", client.getConfig().userInfoURL, nil)
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to create request, %v", err), "")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))

	res, err = o.httpClient.Do(req)

	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to request user info, %v", err), "")
	}
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, appError.New(httpCode.Status{Code: res.StatusCode}, fmt.Sprintf("unexpected status code: %d, body: %s", res.StatusCode, string(body)), "Internal Server Error")
	}

	return client.parseUserInfo(body)
}
