package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
}

var AppConfig *Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		hlog.Fatalf("Error reading config file: %v", err)
	}

	appConfig := &Config{}

	if err := viper.Unmarshal(appConfig); err != nil {
		hlog.Fatalf("Error unmarshaling config: %v", err)
	}

	AppConfig = appConfig
}
