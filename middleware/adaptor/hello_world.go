package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	// New fiber_serv app
	app := fiber.New()

	// http.Handler -> fiber_serv.Handler
	app.Get("/", adaptor.HTTPHandler(handler(Greet)))

	// http.HandlerFunc -> fiber_serv.Handler
	app.Get("/func", adaptor.HTTPHandlerFunc(Greet))

	// Listen on port 3000
	app.Listen(":3000")
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
