package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func Token() string {
	return viper.GetString("token")
}

func ApiHost() string {
	api := viper.GetString("api")
	if api == "" {
		return "api.engineyard.com"
	}

	return api
}

func CheckToken() error {
	if len(Token()) == 0 {
		return fmt.Errorf(
			`This operation requires Engine Yard API authentication.

This should be listed as :token: in /etc/scaley/config.yml`,
		)
	}

	return nil
}
