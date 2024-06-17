package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal("Invalid type assertion")
	}

	return value
}

