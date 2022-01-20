package ginSample

import (
	"net/http"
)

type Engine struct {
	*RouterGroup
	router *Router
	group  []*RouterGroup
}

type H map[string]interface{}

type HandleFunc func(c *Context)

func New() *Engine {
	e := &Engine{
		router: newRouter(),
	}

	e.RouterGroup = &RouterGroup{
		prefix: "",
		engine: e,
	}

	return e
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	handlers, params := e.router.findHandlers(r.Method, r.URL.Path)
	if handlers != nil {
		context.Params = params
		context.handlers = handlers
		context.run()
	} else {
		context.JSON(http.StatusNotFound, H{
			"data": "路由不存在",
		})
	}
}

func (e *Engine) Run(addr string) {
	e.router.logApis()
	http.ListenAndServe(addr, e)
}
