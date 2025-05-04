package migutils

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Do(options Options) error {
	switch options.Migrations.Direction {
	case UP:
		return up(options)
	case DOWN:
		return down(options)
	}

	return errors.New(fmt.Sprintf("unsupported direction: %v", options.Migrations.Direction))
}

func up(options Options) error {
	db, err := sql.Open("postgres", getPostgreDBString(options))
	if err != nil {
		return err
	}

	var (
		embedMigrations = options.Migrations.Embed
		version         = options.Migrations.Version
		dir             = options.Migrations.Dir
	)

	goose.SetBaseFS(embedMigrations)

	if version == LastVersion {
		if err := goose.Up(db, dir); err != nil {
			return err
		}

		return nil
	}

	if err := goose.UpTo(db, dir, int64(version)); err != nil {
		return err
	}

	return nil
}

func down(options Options) error {
	db, err := sql.Open("postgres", getPostgreDBString(options))
	if err != nil {
		return err
	}

	var (
		embedMigrations = options.Migrations.Embed
		version         = options.Migrations.Version
		dir             = options.Migrations.Dir
	)

	goose.SetBaseFS(embedMigrations)

	if version == LastVersion {
		if err := goose.Down(db, dir); err != nil {
			return err
		}

		return nil
	}

	if err := goose.DownTo(db, dir, int64(version)); err != nil {
		return err
	}

	return nil
}

func getPostgreDBString(options Options) string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		options.Database.User, options.Database.Password, options.Database.Addr, options.Database.Database,
	)
}
