package api

import (
	_ "embed"
)

//go:embed ui/index.html
var html_home []byte

//go:embed ui/preview.html
var preview_home []byte

//go:embed ui/filehelper.html
var filehelper_home []byte

//go:embed sph/index.html
var sph_home []byte

type Assets struct {
	HTMLHome       []byte
	HTMLPreview    []byte
	HTMLFilehelper []byte
	HTMLSph        []byte // SPH query page
}

var files = &Assets{
	HTMLHome:       html_home,
	HTMLPreview:    preview_home,
	HTMLFilehelper: filehelper_home,
	HTMLSph:        sph_home,
}
