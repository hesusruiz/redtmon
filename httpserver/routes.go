package httpserver

import (
	"bytes"
	"fmt"
	"text/template"

	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/hesusruiz/redtmon/client"
	"github.com/hesusruiz/redtmon/types"
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

	// Connect to the Blockchain node at the specified URL
	redtNode := srv.cfg.String("redtnode")
	fmt.Println("Target Node URL", redtNode)
	qc, err := client.NewQuorumClient(redtNode)
	if err != nil {
		zlog.Error().Err(err).Msgf("connecting to %v", redtNode)
		return
	}
	defer qc.Stop()

	// Subscribe to receive notifications when new blocks are added to the blockchain
	// Each notification is received as a Header in the inputCh channel
	inputCh := make(chan types.RawHeader)
	err = qc.SubscribeChainHead(inputCh)
	if err != nil {
		zlog.Err(err).Msgf("connecting to %v", redtNode)
		return
	}

	// Initialize the timestamp to calculate elapsed time between blocks
	latestTimestamp := uint64(0)

	// We will start calculating statistics (eg. elapsed time) on the second block
	isFirst := true

	// Loop forever until an error (probably the client disconnects)
	var rendered bytes.Buffer
	var data map[string]any

	for {

		// We block here until a new header is received on the channel
		rawheader := <-inputCh

		var currentHeader *ethertypes.Header

		// Get the full header, because the raw one does not have the info we need
		currentHeader, err = srv.rt.HeaderByNumber(int64(rawheader.Number))
		if err != nil {
			// Log the error and retry with next block
			zlog.Err(err).Msgf("connecting to %v", redtNode)
			return
		}

		if isFirst {
			// Do not display, we just get its timestamp to start statistics
			latestTimestamp = currentHeader.Time
			isFirst = false

			// Wait for the next block
			continue
		}

		// Get the signer data and accumulated statistics
		data, latestTimestamp = srv.rt.SignersForHeader(currentHeader, latestTimestamp)

		// Format the data into an HTML table
		rendered.Reset()
		err = t.ExecuteTemplate(&rendered, "table.html", data)
		if err != nil {
			zlog.Err(err).Msgf("connecting to %v", redtNode)
			return
		}

		// Send the HTML table to the client via the WebSocket connection
		err = c.WriteMessage(websocket.TextMessage, rendered.Bytes())
		if err != nil {
			zlog.Err(err).Msgf("connecting to %v", redtNode)
			return
		}

	}

}
