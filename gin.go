package pongo2render

import (
	"net/http"
	"github.com/gin-gonic/gin/render"
	"github.com/flosch/pongo2"
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
}

func NewGinRender(templateDir string) *GinRender {
	var r = &GinRender{}
	r.TemplateDir = templateDir
	return r
}

func (this GinRender) Instance(name string, data interface{}) render.Render {
	var gHtml = &GinHTML{}
	var h = this.HTML(name)

	if data != nil {
		h.context = data.(pongo2.Context)
	}

	gHtml.HTML = h
	return gHtml
}

func (this *GinHTML) Render(w http.ResponseWriter) (err error) {
	return this.HTML.ExecuteWriter(w, this.context)
}

