package frontend

import "embed"

// Content holds our static web server content.
//go:embed public
var Content embed.FS
