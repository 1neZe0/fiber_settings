package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {

	// fiber_serv.Handler -> http.Handler
	http.Handle("/", adaptor.FiberHandler(greet))

	// fiber_serv.Handler -> http.HandlerFunc
	http.HandleFunc("/func", adaptor.FiberHandlerFunc(greet))

	// Listen on port 3000
	http.ListenAndServe(":3000", nil)
}

func greet(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
