package mount_path

import (
	"fiber_settings/fiber_serv"
	"github.com/gofiber/fiber/v2"
)

// The MountPath property contains one or more path patterns on which a sub-app was mounted.

func StartMountPathApp() {
	app := fiber_serv.GetFiber()
	one := fiber.New()
	two := fiber.New()
	three := fiber.New()

	two.Mount("/three", three)
	one.Mount("/two", two)
	app.Mount("/one", one)

	three.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	one.MountPath()   // "/one"
	two.MountPath()   // "/one/two"
	three.MountPath() // "/one/two/three"
	app.MountPath()   // ""

	fiber_serv.StartFiber(&app)
}

//Mounting order is important for MountPath. If you want to get mount paths properly, you should start mounting from the deepest app.
