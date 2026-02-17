package web

import "embed"

// StaticFS embeds the static frontend assets (index.html, css, js).
// The "static" subdirectory is the root of the embedded filesystem.
//
//go:embed all:static
var StaticFS embed.FS
