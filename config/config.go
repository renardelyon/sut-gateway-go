package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
	StorageSvcUrl string `mapstructure:"STORAGE_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	NotifSvcUrl   string `mapstructure:"NOTIF_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	return
}
