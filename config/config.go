package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

type DB struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	User   string `json:"user"`
	DBName string `json:"db_name"`
}

type Config struct {
	App App `json:"app"`
	DB  DB  `json:"db"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:      viper.GetString("APP_PORT"),
			AppEnv:       viper.GetString("APP_ENV"),
			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:    viper.GetString("JWT_ISSUER"),
		},
		DB: DB{
			Host:   viper.GetString("DATABASE_HOST"),
			Port:   viper.GetString("DATABASE_PORT"),
			User:   viper.GetString("DATABASE_USER"),
			DBName: viper.GetString("DATABASE_NAME"),
		},
	}
}
