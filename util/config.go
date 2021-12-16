package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RabbitUrl    string `mapstructure:"RABBIT_URL"`
	ExchangeName string `mapstructure:"EXCHANGE_NAME"`
	MessageBody  string `mapstructure:"MESSAGE_BODY"`
}

func assertConfigSet(name string) {
	if !viper.IsSet(name) {
		log.Panicf("Missing environment variable %s", name)
	}
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read from `.env` and override with specificed environment variables
	viper.ReadInConfig()

	viper.BindEnv("RABBIT_URL")
	viper.BindEnv("EXCHANGE_NAME")
	viper.BindEnv("MESSAGE_BODY")

	// Assert each required variable is set
	assertConfigSet("RABBIT_URL")
	assertConfigSet("EXCHANGE_NAME")
	assertConfigSet("MESSAGE_BODY")

	err = viper.Unmarshal(&config)
	return
}
