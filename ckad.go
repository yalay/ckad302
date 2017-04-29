package main

import (
	"controllers"
	"flag"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var listenPort int

func init() {
	flag.IntVar(&listenPort, "p", 1320, "p=1320")
	flag.Parse()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ajax/dl/:id", controllers.GetDlLinks)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(listenPort)))
}
