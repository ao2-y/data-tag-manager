package injector

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Database struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

type config struct {
	Database Database
}

func newConfig() *config {
	envName := os.Getenv("ENV_NAME")
	if envName == "" {
		envName = "local"
	}
	viper.SetConfigName(envName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config read error.%w", err))
	}
	var config config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("config unmarshal error.%w", err))
	}

	return &config
}
