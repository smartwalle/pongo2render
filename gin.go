package pongo2render

import (
	"path"
	"net/http"
	"github.com/gin-gonic/gin/render"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

var htmlContentType = []string{"text/html; charset=utf-8"}

type Pongo2GinRender struct {
	TemplateDir string
}

type pongo2GinHTMLRender struct {
	Template *pongo2.Template
	Context  pongo2.Context
}

func NewGinRender(templateDir string) *Pongo2GinRender {
	return &Pongo2GinRender{
		TemplateDir: templateDir,
	}
}

func (this Pongo2GinRender) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	var filename string
	if len(this.TemplateDir) > 0 {
		filename = path.Join(this.TemplateDir, name)
	} else {
		filename = name
	}

	if gin.Mode() == gin.DebugMode {
		template, _ = pongo2.FromFile(filename)
	} else {
		template, _ = pongo2.FromCache(filename)
	}

	if template == nil {
		panic("template " + name + " not exists")
		return nil
	}

	var r = pongo2GinHTMLRender{}
	r.Template = template
	if data != nil {
		r.Context  = data.(pongo2.Context)
	}
	return r
}

func (this pongo2GinHTMLRender) Render(w http.ResponseWriter) (err error) {
	writeContentType(w, htmlContentType)
	err = this.Template.ExecuteWriter(this.Context, w)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}