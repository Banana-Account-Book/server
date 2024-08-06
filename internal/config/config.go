package config

import "os"

var (
	IsProduction = os.Getenv("APP_ENV") == "production"
	Port         = os.Getenv("PORT")
	DbName       = os.Getenv("DB_NAME")
	DbUser       = os.Getenv("DB_USER")
	DbPassword   = os.Getenv("DB_PASSWORD")
	DbHost       = os.Getenv("DB_HOST")
	DbPort       = os.Getenv("DB_PORT")
	Origin       = os.Getenv("ORIGIN")
)
