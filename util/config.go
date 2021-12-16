package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	RabbitUrl    string `mapstructure:"RABBIT_URL"`
	ExchangeName string `mapstructure:"EXCHANGE_NAME"`
	MessageBody  string `mapstructure:"MESSAGE_BODY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read and overwrite from the environment if variables exist
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
