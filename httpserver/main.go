package httpserver

import (
	"fmt"
	"os"

	"github.com/hesusruiz/vcutils/yaml"

	"flag"
	"log"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

const defaultConfigFile = "configs/config.yaml"

var (
	listen     = flag.String("listen", ":8000", "Port to listen on")
	configFile = flag.String("config", defaultConfigFile, "path to configuration file")
	// password   = flag.String("pass", defaultPassword, "admin password for the server")
)

func main() {

	// Initialize the log and its format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zlog.Logger = zlog.With().Caller().Logger()

	// Parse command-line flags
	flag.Parse()

	// Read configuration file
	cfg := readConfiguration(*configFile)

	// Create a new server
	srv := NewServer(cfg)

	// Listen on specified port
	log.Fatal(srv.Listen(cfg.String("server.listenAddress", *listen)))
}

// readConfiguration reads a YAML file and creates an easy-to navigate structure
func readConfiguration(configFile string) *yaml.YAML {
	var cfg *yaml.YAML
	var err error

	cfg, err = yaml.ParseYamlFile(configFile)
	if err != nil {
		fmt.Printf("Config file not found, using defaults\n")
		panic(err)
	}
	return cfg
}
