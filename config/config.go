package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	errors []error
	Host   string
	Port   string
}

func LoadConfig() ServerConfiguration {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	c := ServerConfiguration{}

	c.Host = c.getMandatoryString("SIMILIGO_API_HOST")
	c.Port = c.getMandatoryString("SIMILIGO_API_PORT")

	if len(c.errors) != 0 {
		errorReport := "errors in config :\n"
		for _, err := range c.errors {
			errorReport += fmt.Sprintf("- %s\n", err)
		}
		panic(fmt.Errorf(errorReport))
	}
	return c
}

func (c *ServerConfiguration) getMandatoryString(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		c.errors = append(c.errors, fmt.Errorf("cannot find configuration for key %s", key))
	}
	return value
}
