package migrations

import "embed"

const MigrationsDir string = "migrations"

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
