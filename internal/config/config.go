package config

import "os"

var (
	IsProduction = os.Getenv("APP_ENV") == "production"
	Port         = os.Getenv("PORT")
	DB_NAME      = os.Getenv("DB_NAME")
	DB_USER      = os.Getenv("DB_USER")
	DB_PASSWORD  = os.Getenv("DB_PASSWORD")
	DB_HOST      = os.Getenv("DB_HOST")
	DB_PORT      = os.Getenv("DB_PORT")
)
