package configs

import (
	"github.com/bytedance/sonic"
	"time"

	"github.com/gofiber/fiber/v2"
)

// https://docs.gofiber.io/api/fiber#config
// Config fields:
// BodyLimit: int
// CaseSensitive: bool
// Concurrency: int
// DisableDefaultContentType
// DisableDefaultDate
// DisableHeaderNormalizing
// DisableKeepalive
// EnableIPValidation
// EnablePrintRoutes
// EnableTrustedProxyCheck
// ErrorHandler
// GETOnly
// IdleTimeout
// Immutrable
// Network : NetworkTCP4
// Prefork https://github.com/gofiber/fiber/issues/180
// ProxyHeader
// TrustedProxies
// ViewsLayout

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	//readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	readTimeoutSecondsCount := 60
	//Prefork_on := tre
	//StrictRouting_on := true
	//CaseSensitive := true
	// Return Fiber configuration.
	return fiber.Config{
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		CaseSensitive: true,
		StrictRouting: true,
		Prefork:       true,
		ReadTimeout:   time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
