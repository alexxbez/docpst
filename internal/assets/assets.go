package assets

import "embed"

//go:embed template.typ config.toml main.typ
var Assets embed.FS
