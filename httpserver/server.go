package httpserver

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hesusruiz/redtmon/redt"
	"github.com/hesusruiz/vcutils/yaml"
)

type Server struct {
	*fiber.App
	cfg *yaml.YAML
	rt  *redt.RedTNode
}

func NewServer(cfg *yaml.YAML) *Server {

	srv := &Server{}
	srv.cfg = cfg

	// Connect to the RedT node
	redtNode := srv.cfg.String("redtnode")
	rt, err := redt.NewRedTNode(redtNode)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	srv.rt = rt

	// Define the configuration for Fiber
	fiberCfg := fiber.Config{
		Concurrency:  1024,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  5 * time.Minute,
	}

	// Create a Fiber instance and set it in our Server struct
	srv.App = fiber.New(fiberCfg)

	// Middleware
	srv.Use(recover.New(recover.Config{EnableStackTrace: true}))

	srv.Use(logger.New(logger.Config{
		// TimeFormat: "02-Jan-1985",
		TimeZone: "Europe/Brussels",
	}))

	// CORS (default config)
	srv.Use(cors.New())

	srv.setupRoutes()

	return srv

}
