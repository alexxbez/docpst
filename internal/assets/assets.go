package assets

import "embed"

//go:embed default.typ config.toml
var Assets embed.FS
