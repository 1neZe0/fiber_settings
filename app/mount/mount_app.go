package mount

import (
	"fiber_settings/fiber_serv"
	"github.com/gofiber/fiber/v2"
)

func StartMountApp() {
	app := fiber_serv.GetFiber()
	micro := fiber.New()
	app.Mount("/john", micro) // GET /john/doe -> 200 OK
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	fiber_serv.StartFiber(&app)

}
