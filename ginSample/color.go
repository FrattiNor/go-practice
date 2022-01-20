package ginSample

import (
	"github.com/fatih/color"
)

type methodColor struct {
	Fg color.Attribute
	Bg color.Attribute
}

var (
	Colors = make(map[string]methodColor)
)

func getMethodColor(method string) (string, methodColor) {
	switch method {
	case "GET":
		return " GET    ", methodColor{color.FgHiWhite, color.BgRed}
	case "POST":
		return " POST   ", methodColor{color.FgHiWhite, color.BgBlue}
	case "PUT":
		return " PUT    ", methodColor{color.FgHiWhite, color.BgGreen}
	case "DELETE":
		return " DELETE ", methodColor{color.FgHiWhite, color.BgCyan}
	default:
		return " " + method + " ", methodColor{color.FgHiWhite, color.BgYellow}
	}
}

func colorPrint(method string, pattern string) {
	methodStr, methodColors := getMethodColor(method)

	c := color.New(methodColors.Fg)
	c.Add(methodColors.Bg)
	c.Print(methodStr)

	c2 := color.New(color.FgHiWhite)
	c2.Println(" ", "/"+pattern, " ")
}
