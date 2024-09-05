package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mongodb MongodbConfig `mapstructure:"mongodb"`
	JWT     JWTConfig     `mapstructure:"jwt"`
}

type (
	MongodbConfig struct {
		Host string `mapstructure:"host"`
		DB   string `mapstructure:"db"`
	}

	JWTConfig struct {
		Secret string `mapstructure:"secret"`
	}
)

func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}
