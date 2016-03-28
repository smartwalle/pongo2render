package pongo2render

import (
	"net/http"
	"github.com/gin-gonic/gin/render"
)

//	var router = gin.Default()
//	router.HTMLRender = pongo2render.NewGinRender("./templates")
//
//	router.GET("/m", func(c *gin.Context) {
//		c.HTML(200, "index.html", pongo2.Context{"aa": "eee"})
//	})
//	router.Run("localhost:9005")

type GinRender struct {
	Render
}

type GinHTML struct {
	*HTML
	data interface{}
}

func NewGinRender(templateDir string) *GinRender {
	var r = &GinRender{}
	r.TemplateDir = templateDir
	return r
}

func (this GinRender) Instance(name string, data interface{}) render.Render {
	var gHtml = &GinHTML{}
	var h = this.html(name)
	gHtml.HTML = h
	gHtml.data = data
	return gHtml
}

func (this *GinHTML) Render(w http.ResponseWriter) (err error) {
	return this.HTML.ExecuteWriter(w, this.data)
}

