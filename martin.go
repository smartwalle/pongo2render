package pongo2render

import (
	"net/http"
	"github.com/go-martini/martini"
)

//	m := martini.Classic()
//	m.Use(pongo2render.NewMartinRender("./templates", false))
//	m.Get("/", func(r pongo2render.MartinRender) {
//		r.HTML("index.html", pongo2.Context{"aa": "eafdsF"})
//	})
//  m.RunOnAddr(":9005")

type MartinRender interface {
	HTML(name string, data interface{})
}

type martinRender struct {
	res    http.ResponseWriter
	req    *http.Request
	render *Render
}

func (this *martinRender) HTML(name string, data interface{}) {
	this.render.html(name).ExecuteWriter(this.res, data)
}

func NewMartinRender(templateDir string, cache bool) martini.Handler {
	var render = NewRender(templateDir)
	render.Cache = cache
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		var mr = &martinRender{}
		mr.res = res
		mr.req = req
		mr.render = render
		c.MapTo(mr, (*MartinRender)(nil))
	}
}