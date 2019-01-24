package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gen1us2k/service-discovery/api/config"
	"github.com/labstack/echo"
)

func main() {
	c, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		name, err := os.Hostname()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "Hello, World from "+name)
	})
	fmt.Println(c)
	e.Logger.Fatal(e.Start(c.HTTPBindAddr))
}
