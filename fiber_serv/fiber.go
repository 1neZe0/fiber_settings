package fiber_serv

import (
	configs "fiber_settings/fiber_serv/config"
	"fiber_settings/fiber_serv/utils"
	"github.com/gofiber/fiber/v2"
	"os"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func GetFiber() (configured_fiber fiber.App) {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	//// Middlewares.
	//middleware.FiberMiddleware(app) // Register Fiber's  middleware for app.
	//
	//// Routes.
	//routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	//routes.PublicRoutes(app)  // Register a public routes for app.
	//routes.PrivateRoutes(app) // Register a private routes for app.
	//routes.NotFoundRoute (app) // Register route for 404 Error.
	//
	//// Start server (with or without graceful shutdown).

	return *app
}

func StartFiber(app *fiber.App) (started_fiber fiber.App) {
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

	return *app
}
