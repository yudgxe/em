package cmd

import (
	"em/pkg/utils"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "em",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		configFile, err := rootCmd.PersistentFlags().GetString("env")
		if err != nil {
			utils.Panicf("error on Parse env - %v", err)
		}

		viper.SetConfigFile(configFile)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			utils.Panicf("error on ReadInConfig - %v", err)
		}

		// TODO: setup default value
	})

	rootCmd.PersistentFlags().String("env", ".env", "env file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
