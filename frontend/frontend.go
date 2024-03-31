package frontend

import "embed"

//go:embed all:build/*
var BuildProd embed.FS
