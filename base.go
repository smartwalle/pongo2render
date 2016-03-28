package pongo2render

import (
	"path"
	"net/http"
	"github.com/flosch/pongo2"
)

	//var render = pongo2render.NewRender("./templates")
	//
	//http.HandleFunc("/m", func(w http.ResponseWriter, req *http.Request) {
	//	render.Template("index.html").ExecuteWriter(w, pongo2.Context{"aa": "eeeeeee"})
	//})
	//http.ListenAndServe(":9005", nil)

var htmlContentType = []string{"text/html; charset=utf-8"}

type Render struct {
	TemplateDir string
	Cache       bool
}

type HTMLTemplate struct {
	Template *pongo2.Template
	context  pongo2.Context
}

func NewRender(templateDir string) *Render {
	var r = &Render{}
	r.TemplateDir = templateDir
	return r
}

func (this *Render) Template(name string) *HTMLTemplate {
	var template *pongo2.Template
	var filename string
	if len(this.TemplateDir) > 0 {
		filename = path.Join(this.TemplateDir, name)
	} else {
		filename = name
	}

	if this.Cache {
		template, _ = pongo2.FromCache(filename)
	} else {
		template, _ = pongo2.FromFile(filename)
	}

	if template == nil {
		panic("template " + name + " not exists")
		return nil
	}

	var r = &HTMLTemplate{}
	r.Template = template
	return r
}

func (this *HTMLTemplate) ExecuteWriter(w http.ResponseWriter, c pongo2.Context) (err error) {
	WriteContentType(w, htmlContentType)
	this.context = c
	err = this.Template.ExecuteWriter(this.context, w)
	return err
}

func WriteContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}