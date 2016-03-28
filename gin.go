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

type GinHTMLTemplate struct {
	*HTMLTemplate
}

func NewGinRender(templateDir string) *GinRender {
	var r = &GinRender{}
	r.TemplateDir = templateDir
	return r
}

func (this GinRender) Instance(name string, data interface{}) render.Render {
	var ginTemplate = &GinHTMLTemplate{}
	var t = this.Template(name)

	if data != nil {
		t.context = data.(pongo2.Context)
	}

	ginTemplate.HTMLTemplate = t
	return ginTemplate
}

func (this *GinHTMLTemplate) Render(w http.ResponseWriter) (err error) {
	return this.HTMLTemplate.ExecuteWriter(w, this.context)
}

