package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App      `yaml:"app"`
		Database Database `yaml:"database"`
		Log      Log      `yaml:"log"`
		Token    Token    `yaml:"token"`
	}

	App struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		BaseUrl string `yaml:"baseUrl"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbName"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Log struct {
		Level string `yaml:"level"`
	}

	Token struct {
		TokenTTL  string `yaml:"tokenTTL"`
		JWTSecret string `yaml:"jwtSecret"`
	}
)

func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error  reading  config file: %w", err)
		}
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}
