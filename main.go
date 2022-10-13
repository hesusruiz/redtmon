package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hesusruiz/redtmon/httpserver"
	"github.com/hesusruiz/redtmon/redt"
	"github.com/urfave/cli/v2"

	"github.com/hesusruiz/vcutils/yaml"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

const defaultConfigFile = "configs/config.yaml"

func main() {

	// Initialize the log and its format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zlog.Logger = zlog.With().Caller().Logger()

	// The main command without any options starts the HTTP server
	app := &cli.App{
		Usage: "Monitoring of an Alastria RedT blockchain node",

		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Version:                "v0.3",
		Compiled:               time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jesus Ruiz",
				Email: "hesus.ruiz@gmail.com",
			},
		},

		Action: func(c *cli.Context) error {
			runHTTPServer(c)
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "c",
				Value:    defaultConfigFile,
				Usage:    "location for configuration file",
				Aliases:  []string{"config"},
				Required: false,
			},
		},
	}

	// The Display Peers command
	displayPeersCMD := &cli.Command{
		Name:  "peers",
		Usage: "display peers information",

		Action: func(c *cli.Context) error {
			cfg := readConfiguration(c)
			url := cfg.String("redtnode")
			redt.DisplayPeersInfo(url)
			return nil
		},
	}

	// The TEST
	testCMD := &cli.Command{
		Name:  "test",
		Usage: "test several things",

		Action: func(c *cli.Context) error {
			cfg := readConfiguration(c)
			// url := cfg.String("redtnode")
			readValidatorListFromGithub(cfg)
			return nil
		},
	}

	app.Commands = []*cli.Command{
		displayPeersCMD,
		testCMD,
	}

	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		zlog.Fatal().Err(err).Msg("")
	}

}

func runHTTPServer(c *cli.Context) {

	// Read configuration file
	cfg := readConfiguration(c)

	// Create a new server
	srv := httpserver.NewServer(cfg)

	// Listen from a different goroutine
	go func() {
		err := srv.Listen(cfg.String("server.listenAddress"))
		if err != nil {
			zlog.Panic().Err(err).Msg("unclean exit")
		}
	}()

	// Install signal handlers to perform clean shutdown
	shut := make(chan os.Signal, 1)
	signal.Notify(shut, os.Interrupt, syscall.SIGTERM)

	<-shut // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = srv.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Server was successful shutdown.")

}

// readConfiguration reads a YAML file and creates an easy-to navigate structure
func readConfiguration(c *cli.Context) *yaml.YAML {
	var cfg *yaml.YAML
	var err error

	// Get the path of the configuration file
	cfg_path := c.String("c")
	zlog.Info().Msgf("reading config file: %v", cfg_path)

	cfg, err = yaml.ParseYamlFile(cfg_path)
	if err != nil {
		fmt.Printf("Config file not found, exiting\n")
		os.Exit(1)
	}
	return cfg
}

func readValidatorListFromGithub(cgf *yaml.YAML) []*redt.NodeInfo {

	// Read the raw file, with one line per Validator
	// Each line is like: VAL_IN2 enode://0ede782b7ce6...2b8e64059adc03@15.236.56.133:21000?discport=0
	// Where the dots are an ellipsis to make the line shorter
	resp, err := http.Get("https://raw.githubusercontent.com/alastria/alastria-node-quorum-directory/main/directory.val")
	if err != nil {
		zlog.Fatal().Err(err).Msg("")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zlog.Fatal().Err(err).Msg("")
	}

	// Process each line
	tokens := bytes.Fields(body)
	num_tokens := len(tokens)

	// Some sanity checks
	// The number of tokens should be even (and greater than zero)
	if num_tokens <= 0 || num_tokens%2 != 0 {
		zlog.Panic().Msgf("invalid number of tokens: %v", len(tokens))

	}

	var enodes = make([]*redt.NodeInfo, num_tokens/2)
	for i, t := range tokens {
		if i%2 == 0 {
			enode := &redt.NodeInfo{}
			enodes[i/2] = enode
			t = bytes.TrimPrefix(t, []byte("VAL_"))
			enodes[i/2].Operator = string(t)
		} else {
			// The odd tokens should start with "enode://"
			if bytes.HasPrefix(t, []byte("enode://")) {
				enodes[i/2].Enode = string(t)
			} else {
				fmt.Println("caca")
			}
		}
	}

	return enodes

}
