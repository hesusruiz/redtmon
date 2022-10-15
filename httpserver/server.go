package httpserver

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hesusruiz/redtmon/redt"
	"github.com/hesusruiz/vcutils/yaml"
	"github.com/rs/zerolog/log"
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
	nodeURL := srv.cfg.String("redtnode")
	rt, err := redt.NewRedTNode(nodeURL)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
	srv.rt = rt

	// Check that we can get info about the node
	thisNode, err := srv.rt.NodeInfo()
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
	log.Info().Str("name", thisNode.Name).Msg("connected")

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
	srv.registerMiddleware()

	srv.setupRoutes()

	return srv

}

func (srv *Server) registerMiddleware() {

	srv.Use(logger.New(logger.Config{
		TimeZone: "Europe/Brussels",
	}))

	srv.Use(recover.New(recover.Config{EnableStackTrace: true}))

	// CORS
	if srv.cfg.Bool("server.cors.enabled") {
		srv.Use(cors.New(cors.Config{
			AllowOrigins:     srv.cfg.String("server.cors.alloworigins", "*"),
			AllowMethods:     srv.cfg.String("server.cors.allowmethods", "GET,POST,HEAD,PUT,DELETE,PATCH"),
			AllowHeaders:     srv.cfg.String("server.cors.allowheaders"),
			AllowCredentials: srv.cfg.Bool("server.cors.allowcredentials"),
			ExposeHeaders:    srv.cfg.String("server.cors.exposeheaders"),
			MaxAge:           srv.cfg.Int("server.cors.maxage"),
		}))
	}

	if srv.cfg.Bool("server.favicon.enabled") {
		srv.Use(favicon.New(favicon.Config{
			File:         srv.cfg.String("server.favicon.file"),
			CacheControl: srv.cfg.String("server.favicon.cachecontrol"),
		}))
	}

	// Limit the number of requests per second from a given IP
	srv.Use(limiter.New())

}
