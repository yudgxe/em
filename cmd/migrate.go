package cmd

import (
	"em/migrations"
	"em/pkg/migutils"
	"em/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		direction := args[0]

		options := migutils.Options{
			Database: migutils.Database{
				User:     viper.GetString("DB_USER"),
				Password: viper.GetString("DB_PASSWORD"),
				Addr:     viper.GetString("DB_ADDR"),
				Database: viper.GetString("DB_DATABASE"),
			},
			Migrations: migutils.Migrations{
				Embed:     migrations.EmbedMigrations,
				Dir:       migrations.MigrationsDir,
				Version:   migutils.LastVersion,
				Direction: migutils.Direction(direction),
			},
		}

		if err := migutils.Do(options); err != nil {
			utils.Panicf("error on do migrations - %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
