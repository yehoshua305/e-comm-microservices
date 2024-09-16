package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment                string        `mapstructure:"ENVIRONMENT"`
	ServerAddress              string        `mapstructure:"ServerAddress"`
	DYNAMODB_ACCESS_KEYID      string        `mapstructure:"DYNAMODB_ACCESS_KEYID"`
	DYNAMODB_SECRET_ACCESS_KEY string        `mapstructure:"DYNAMODB_SECRET_ACCESS_KEY"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
