package scripts

import (
	"embed"
	"io/fs"
)

//go:embed static
var (
	FSstatic       embed.FS
	StaticFiles, _ = fs.Sub(FSstatic, "static")
)
