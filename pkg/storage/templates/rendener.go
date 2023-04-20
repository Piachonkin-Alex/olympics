package templates

import "embed"

//go:embed sql/*
var sqlStatements embed.FS
