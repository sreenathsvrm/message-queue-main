package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`

	BROKER_ADDR  string `mapstructure:"BROKER_ADDR"`
	BROKER_TOPIC string `mapstructure:"BROKER_TOPIC"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "BROKER_ADDR", "BROKER_TOPIC",
}

func GetConfig() (*Config, error) {
	var config *Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Read from env
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, errors.New("error binding " + err.Error())
		}
	}

	// De serialize config values
	if err := viper.Unmarshal(&config); err != nil {
		return config, errors.New("error un marshalling " + err.Error())
	}

	fmt.Println("DB PORT ", config.DB_PORT)
	fmt.Println("DB PASSWORD ", config.DB_PASSWORD)

	return config, nil
}
