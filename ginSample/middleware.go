package ginSample

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func LogTime() HandleFunc {
	return func(c *Context) {
		start := time.Now()
		fmt.Println("start", c.Req.Method, c.Req.URL.Path)
		c.Next()
		fmt.Println("end  ", c.Req.Method, c.Req.URL.Path, time.Now().Sub(start))
	}
}

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandleFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
