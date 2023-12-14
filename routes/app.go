package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	app.Start()
}
