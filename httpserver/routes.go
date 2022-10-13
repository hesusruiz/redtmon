package httpserver

import (
	"bytes"
	"context"
	"text/template"
	"time"

	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/hesusruiz/redtmon/redt"
	zlog "github.com/rs/zerolog/log"
)

const defaultStaticDir = "back/www"

func (srv *Server) setupRoutes() {

	// Upgraded websocket request
	srv.Get("/ws", websocket.New(srv.WSHandler))

	// Setup static files
	srv.Static("/static", srv.cfg.String("server.staticDir", defaultStaticDir))

	srv.Get("/", HandleHome)

}

func HandleHome(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(indexHTML)
}

// WSHandler is the WebSocket handler for an already upgraded connection
func (srv *Server) WSHandler(c *websocket.Conn) {

	// Prepare the template for the HTML Table with the block info
	t := template.Must(template.New("table.html").Parse(tableHTML))

	// We are going to call the Geth API, with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Subscribe to receive notifications when new blocks are added to the blockchain
	// Each notification is received as a Header in the inputCh channel
	notificationCh := make(chan *ethertypes.Header)
	subscription, err := srv.rt.EthClient().SubscribeNewHead(ctx, notificationCh)
	if err != nil {
		zlog.Err(err).Msgf("subscribing")
		return
	}
	defer subscription.Unsubscribe()

	// Initialize statistics for this client
	stats := redt.NewStatistics(srv.rt.AllValidators(), srv.rt.Validators())

	// Initialize the timestamp to calculate elapsed time between blocks
	latestTimestamp := uint64(0)

	// We will start calculating statistics (eg. elapsed time) on the second block
	isFirst := true

	// Loop forever until an error (probably the client disconnects)
	var rendered bytes.Buffer
	var data map[string]any

	for {

		// We block here until a new header is received on the channel
		currentHeader := <-notificationCh

		zlog.Info().Int64("number", currentHeader.Number.Int64()).Msg("new block")

		if isFirst {
			// Do not display, we just get its timestamp to start statistics
			latestTimestamp = currentHeader.Time
			isFirst = false

			// Wait for the next block
			continue
		}

		// Get the signer data and accumulated statistics
		// data, latestTimestamp = srv.rt.StatisticsForHeader(currentHeader, latestTimestamp)
		data, latestTimestamp = stats.StatisticsForHeader(currentHeader, latestTimestamp)

		// Format the data into an HTML table
		rendered.Reset()
		err = t.ExecuteTemplate(&rendered, "table.html", data)
		if err != nil {
			zlog.Err(err).Msg("executing template")
			return
		}

		// Send the HTML table to the client via the WebSocket connection
		err = c.WriteMessage(websocket.TextMessage, rendered.Bytes())
		if err != nil {
			zlog.Err(err).Msgf("sending to browser websocket")
			return
		}

	}

}
