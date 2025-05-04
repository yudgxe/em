package cmd

import (
	"context"
	"em/internal/app"
	"em/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(viper.GetInt("LOG_LEVEL")))

		conn := fmt.Sprintf("postgres://%v:%v@%v/%v",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_ADDR"),
			viper.GetString("DB_DATABASE"),
		)

		log.Info("start init app")

		app, err := app.New(context.TODO(), conn)
		if err != nil {
			utils.Panicf("app.New: %v", err)
		}

		log.Infof("start server on %v", viper.GetString("SERVER_ADDRES"))

		if err := app.Listen(viper.GetString("SERVER_ADDRES")); err != nil {
			utils.Panicf("app.Listen :%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
