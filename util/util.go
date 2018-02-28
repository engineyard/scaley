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

This should be listed as token: in /etc/scaley/config.yml`,
		)
	}

	return nil
}

func ReportingURL() string {
	return viper.GetString("reporting_url")
}

func CheckReportingURL() error {
	if len(ReportingURL()) == 0 {
		return fmt.Errorf(
			`You must provide a reporting URL.

This should be listed as reporting_url: in /etc/scaley/config.yml`,
		)
	}

	return nil
}
