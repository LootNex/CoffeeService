package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Kafka struct {
		Brokers []string
		Topic   string
	}
}

func ConfigLoad() (*Config, error) {

	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file err: %v", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct err: %v", err)
	}

	return &cfg, nil
}
