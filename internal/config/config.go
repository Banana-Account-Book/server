package config

import "os"

type OAuthProviderConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type OauthConfig struct {
	Kakao OAuthProviderConfig
	Naver OAuthProviderConfig
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
			ClientId:     os.Getenv("KAKAO_CLIENT_ID"),
			ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("KAKAO_REDIRECT_URI"),
		},
		Naver: OAuthProviderConfig{
			ClientId:     os.Getenv("NAVER_CLIENT_ID"),
			ClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("NAVER_REDIRECT_URI"),
		},
	}
)
