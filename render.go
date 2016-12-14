package pongo2render

import (
	"path"
	"net/http"
	"github.com/flosch/pongo2"
)

//	var render = pongo2render.NewRender("./templates")
//
//	http.HandleFunc("/m", func(w http.ResponseWriter, req *http.Request) {
//		render.HTML(w, 200, "index.html", pongo2.Context{"aa": "eeeeeee"})
//	})
//	http.ListenAndServe(":9005", nil)

var htmlContentType = []string{"text/html; charset=utf-8"}

type Render struct {
	TemplateDir string
	Cache       bool
}

type HTML struct {
	Template *pongo2.Template
	context  pongo2.Context
}

func NewRender(templateDir string) *Render {
	var r = &Render{}
	r.TemplateDir = templateDir
	return r
}

func (this *Render) GetHTML(name string) *HTML {
	var template *pongo2.Template
	var filename string
	if len(this.TemplateDir) > 0 {
		filename = path.Join(this.TemplateDir, name)
	} else {
		filename = name
	}

	if this.Cache {
		template  = pongo2.Must(pongo2.FromCache(filename))
	} else {
		template  = pongo2.Must(pongo2.FromFile(filename))
	}

	if template == nil {
		panic("template " + name + " not exists")
		return nil
	}

	var r = &HTML{}
	r.Template = template
	return r
}

func (this *Render) HTML(w http.ResponseWriter, status int, name string, data interface{}) {
	w.WriteHeader(status)
	this.GetHTML(name).ExecuteWriter(w, data)
}

func (this *HTML) ExecuteWriter(w http.ResponseWriter, data interface{}) (err error) {
	WriteContentType(w, htmlContentType)
	this.context = DataToContext(data)
	err = this.Template.ExecuteWriter(this.context, w)
	return err
}

func (this *HTML) Execute(data interface{}) (string, error) {
	this.context = DataToContext(data)
	return this.Template.Execute(data)
}

func WriteContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func DataToContext(data interface{}) pongo2.Context {
	var ctx pongo2.Context
	if data != nil {
		switch data.(type) {
		case pongo2.Context:
			ctx = data.(pongo2.Context)
		case map[string]interface{}:
			ctx = pongo2.Context(data.(map[string]interface{}))
		}
	}
	return ctx
}