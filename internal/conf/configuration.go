package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host			   string `mapstructure:"SERVER_HOST"`
	Port               string `mapstructure:"SERVER_PORT"`
	DbHost             string `mapstructure:"DB_HOST"`
}

func LoadConfig() (conf Config) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Error loading config file")
		return Config{}
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalln("Unable to decode into struct")
		return Config{}
	}

	return conf
}