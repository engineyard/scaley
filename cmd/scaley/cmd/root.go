package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "scaley",
	Short: "Generalized faux autoscaling for Engine Yard servers",
	Long:  `Generalized faux autoscaling for Engine Yard servers`,
}

func Execute() error {
	err := RootCmd.Execute()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/scaley")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
	}
}
