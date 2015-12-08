package server

import (
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
)

func GetServer() *macaron.Macaron {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer(pongo2.Options{
		// Directory to load templates. Default is "templates".
		Directory: "templates",
		// Extensions to parse template files from. Defaults are [".tmpl", ".html"].
		Extensions: []string{".tmpl", ".html"},
		// Appends the given charset to the Content-Type header. Default is "UTF-8".
		Charset: "UTF-8",
		// Outputs human readable JSON. Default is false.
		IndentJSON: true,
		// Outputs human readable XML. Default is false.
		IndentXML: true,
		// Allows changing of output to XHTML instead of HTML. Default is "text/html".
		HTMLContentType: "text/html",
	}))

	m.Use(macaron.Static("./static", macaron.StaticOptions{Prefix: "/static"}))
	return m
}
