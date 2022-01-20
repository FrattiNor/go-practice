package main

import (
	"fmt"
	"goServer/ginSample"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	e := ginSample.New()
	e.Use(ginSample.Recovery())
	e.Use(ginSample.LogTime())
	e.Static("/static", "./static")

	home := e.Group("/api")
	{
		home.GET("/hello/:name", func(c *ginSample.Context) {
			panic("okk11zz")
			time.Sleep(time.Second * 10)
			c.JSON(http.StatusOK, ginSample.H{
				"data": c.Param("name"),
			})
		})

		home.POST("/hello/:name", func(c *ginSample.Context) {
			c.JSON(http.StatusOK, ginSample.H{
				"data": c.Param("name1"),
			})
		})

		home.POST("/okk/*name", func(c *ginSample.Context) {
			c.JSON(http.StatusOK, ginSample.H{
				"data": c.Param("name"),
			})
		})

		home.POST("/okk/okk", func(c *ginSample.Context) {
			c.JSON(http.StatusOK, ginSample.H{
				"data": c.PostForm("username"),
				"okk":  "okk",
			})
		})

		home.GET("/hello/okk", func(c *ginSample.Context) {
			c.JSON(http.StatusOK, ginSample.H{
				"data": "okk",
			})
		})
	}

	e.Run(":9099")
}
