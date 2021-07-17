package main

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func HealthHandler(c *routing.Context) error {
	fmt.Fprintf(c, "{\"status\": \"OK\"}")
	return nil
}

func GreeterHandler(c *routing.Context) error {
	fmt.Fprintf(c, "Hello %v!", c.Param("name"))
	return nil
}

func main() {

	router := routing.New()

	router.Get("/health", HealthHandler)
	router.Get("/greeting/<name>", GreeterHandler)

	panic(fasthttp.ListenAndServe(":8000", router.HandleRequest))
}
