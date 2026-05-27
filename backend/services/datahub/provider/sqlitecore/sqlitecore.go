package sqlitecore

import (
	"embed"
)

//go:embed migrations/*
var MigrationBox embed.FS
