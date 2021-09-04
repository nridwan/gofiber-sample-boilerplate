package main

import (
	"log"

	"github.com/nridwan/config"
	"github.com/nridwan/config/configutil"
	"github.com/nridwan/features"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// HostSwitch for CheckHost helper
type HostSwitch []string

// CheckHost Check which domain is it
func (hs HostSwitch) CheckHost(ctx *fiber.Ctx) error {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	for _, host := range hs {
		if host == ctx.Hostname() {
			return ctx.Next()
		}
	}
	return ctx.Status(404).SendString("Not Found")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadAllConfiguration()
	// Make a new HostSwitch and insert the router (our http handler)
	// for example.com and port 12345
	host := configutil.Getenv("APP_HOST", "localhost:8000")
	hs := HostSwitch{host}

	app := fiber.New()
	app.Use(hs.CheckHost)
	features.Register(app)

	log.Fatal(app.Listen(host))
}
