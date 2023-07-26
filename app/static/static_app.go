package static

import (
	"fiber_settings/fiber_serv"
	"github.com/gofiber/fiber/v2"
	"time"
)

func static_config() {
	// Static defines configuration options when defining static assets.
	type Static struct {
		// When set to true, the server tries minimizing CPU usage by caching compressed files.
		// This works differently than the github.com/gofiber/compression middleware.
		// Optional. Default value false
		Compress bool `json:"compress"`

		// When set to true, enables byte range requests.
		// Optional. Default value false
		ByteRange bool `json:"byte_range"`

		// When set to true, enables directory browsing.
		// Optional. Default value false.
		Browse bool `json:"browse"`

		// When set to true, enables direct download.
		// Optional. Default value false.
		Download bool `json:"download"`

		// The name of the index file for serving a directory.
		// Optional. Default value "index.html".
		Index string `json:"index"`

		// Expiration duration for inactive file handlers.
		// Use a negative time.Duration to disable it.
		//
		// Optional. Default value 10 * time.Second.
		CacheDuration time.Duration `json:"cache_duration"`

		// The value for the Cache-Control HTTP-header
		// that is set on the file response. MaxAge is defined in seconds.
		//
		// Optional. Default value 0.
		MaxAge int `json:"max_age"`

		//// ModifyResponse defines a function that allows you to alter the response.
		////
		//// Optional. Default: nil
		//ModifyResponse Handler
		//
		//// Next defines a function to skip this middleware when returned true.
		////
		//// Optional. Default: nil
		//Next func(c *Ctx) bool
	}
}
func RunStatic() {
	app := fiber_serv.GetFiber()
	//app.Get("/")
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // send text
	})

	fiber_serv.StartFiber(&app)
}
