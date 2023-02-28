package main

import (
	"path"

	"github.com/example/{{ cookiecutter.project_name }}/views"
	"github.com/gin-gonic/gin"
)

const (
	TemplatePath = "templates"
	StaticPath   = "/static"
)

func initRouter(router *gin.Engine) {
	router.GET("/", views.Homepage)
	router.SetFuncMap(views.FuncMap)
	router.LoadHTMLGlob(path.Join(TemplatePath, "/**/*"))
	router.Static(StaticPath, StaticPath)
}
