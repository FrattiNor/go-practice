package ginSample

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix     string
	engine     *Engine
	middleware []*HandleFunc
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	currentGroup := &RouterGroup{prefix: g.prefix + prefix, engine: g.engine, middleware: g.middleware}
	g.engine.group = append(g.engine.group, currentGroup)
	return currentGroup
}

func (g *RouterGroup) Use(handler HandleFunc) {
	g.middleware = append(g.middleware, &handler)
}

func (g *RouterGroup) Static(relativePath string, root string) {
	fs := http.Dir(root)
	fileServer := http.StripPrefix(path.Join(g.prefix, relativePath), http.FileServer(fs))

	g.GET(path.Join(relativePath, "/*filepath"), func(c *Context) {
		file := c.Param("filepath")

		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	})
}

func (g *RouterGroup) GET(pattern string, handler HandleFunc) {
	g.engine.router.addRoute("GET", g.prefix+pattern, handler, g.middleware)
}

func (g *RouterGroup) POST(pattern string, handler HandleFunc) {
	g.engine.router.addRoute("POST", g.prefix+pattern, handler, g.middleware)
}

func (g *RouterGroup) PUT(pattern string, handler HandleFunc) {
	g.engine.router.addRoute("PUT", g.prefix+pattern, handler, g.middleware)
}

func (g *RouterGroup) DELETE(pattern string, handler HandleFunc) {
	g.engine.router.addRoute("DELETE", g.prefix+pattern, handler, g.middleware)
}
