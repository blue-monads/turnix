package output

import "embed"

//go:embed all:build/*
var BuildProd embed.FS
