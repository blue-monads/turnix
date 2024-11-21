package frontend

import "embed"

//go:embed all:output/build/*
var BuildProd embed.FS
