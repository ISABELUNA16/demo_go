package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func Hello (c echo.Context) error {
	name := c.Param("name")

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	return c.JSON(http.StatusOK, "Hola " +name)

}

func main(){
	//Create construct
	e := echo.New()
	e.GET("/hello/:name", Hello)
	e.Logger.Fatal(e.Start(":9082"))
}