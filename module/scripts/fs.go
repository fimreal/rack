package scripts

import (
	"embed"
	"io/fs"
)

var (
	//go:embed static
	FSstatic       embed.FS
	StaticFiles, _ = fs.Sub(FSstatic, "static")
)
