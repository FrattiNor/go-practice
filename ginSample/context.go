package ginSample

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer   http.ResponseWriter
	Req      *http.Request
	Params   map[string]string
	handlers []*HandleFunc
	index    int
	code     int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
	}
}

func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
	c.code = code
}

func (c *Context) JSON(code int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "json/application; charset=utf-8")
	c.Status(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

func (c *Context) String(code int, data []byte) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Status(code)

	c.Writer.Write(data)
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	//return c.Req.Form.Get(key)
	return c.Req.FormValue(key)
}

func (c *Context) run() {
	for c.index < len(c.handlers) {
		(*c.handlers[c.index])(c)
		c.index++
	}
}

func (c *Context) Next() {
	if c.index < len(c.handlers) {
		c.index++
		(*c.handlers[c.index])(c)
	}
}
