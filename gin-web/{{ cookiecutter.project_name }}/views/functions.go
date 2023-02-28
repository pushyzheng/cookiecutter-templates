package views

import "html/template"

var FuncMap template.FuncMap

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func init() {
	FuncMap = template.FuncMap{
		"safeHTML": safeHTML,
	}
}
