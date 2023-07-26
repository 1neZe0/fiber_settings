package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	// New fiber_serv app
	app := fiber.New()

	// http middleware -> fiber_serv.Handler
	app.Use(adaptor.HTTPMiddleware(logMiddleware))

	// Listen on port 3000
	app.Listen(":3000")
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("log middleware")
		next.ServeHTTP(w, r)
	})
}
