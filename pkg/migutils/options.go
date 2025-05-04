package migutils

import "io/fs"

const LastVersion uint = 0

type Direction string

const (
	UP   Direction = "up"
	DOWN Direction = "down"
)

type Database struct {
	User     string
	Password string
	Addr     string
	Database string
}

type Migrations struct {
	Embed     fs.FS
	Dir       string
	Version   uint
	Direction Direction
}

type Options struct {
	Database   Database
	Migrations Migrations
}
