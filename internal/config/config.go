package config

import "os"

type OAuthProviderConfig struct {
	BaseURL      string
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type OauthConfig struct {
	Kakao  OAuthProviderConfig
	Naver  OAuthProviderConfig
	Google OAuthProviderConfig
}

var (
	IsProduction = os.Getenv("APP_ENV") == "production"
	Port         = os.Getenv("PORT")
	DbName       = os.Getenv("DB_NAME")
	DbUser       = os.Getenv("DB_USER")
	DbPassword   = os.Getenv("DB_PASSWORD")
	DbHost       = os.Getenv("DB_HOST")
	DbPort       = os.Getenv("DB_PORT")
	Origin       = os.Getenv("ORIGIN")
	Salt         = os.Getenv("SALT")
	SecretKey    = os.Getenv("SECRET_KEY")
	Oauth        = OauthConfig{
		Kakao: OAuthProviderConfig{
			BaseURL:      "https://kauth.kakao.com/oauth",
			ClientId:     os.Getenv("KAKAO_CLIENT_ID"),
			ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("KAKAO_REDIRECT_URI"),
		},
		Naver: OAuthProviderConfig{
			BaseURL:      "https://nid.naver.com/oauth2.0",
			ClientId:     os.Getenv("NAVER_CLIENT_ID"),
			ClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("NAVER_REDIRECT_URI"),
		},
		Google: OAuthProviderConfig{
			BaseURL:      "https://accounts.google.com/o/oauth2",
			ClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("GOOGLE_REDIRECT_URI"),
		},
	}
)
