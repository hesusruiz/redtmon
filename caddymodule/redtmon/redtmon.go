package redtmon

import (
	"bytes"
	"context"
	"net/http"
	"text/template"
	"time"

	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/hesusruiz/redtmon/redt"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("redtnode", parseCaddyfile)
}

// Middleware implements an HTTP handler that writes the
// visitor's IP address to a file or stream.
type Middleware struct {
	// The location of the configuration file
	NodeURL string `json:"nodeurl,omitempty"`

	// The connection to the RedT blockchain node
	rt *redt.RedTNode

	// The precompiled template for sending to the client
	t *template.Template

	// The logger
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.redtnode",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) (err error) {
	m.logger = ctx.Logger()

	// Prepare the template for the HTML Table with the block info
	m.t = template.Must(template.New("table.html").Parse(tableHTMLNew))

	// Connect to the RedT node
	m.rt, err = redt.NewRedTNode(m.NodeURL)
	if err != nil {
		m.logger.Error("connecting to RedT node")
		return err
	}

	// Check that we can get info about the node
	thisNode, err := m.rt.NodeInfo()
	if err != nil {
		m.logger.Error("connecting to RedT node")
		return err
	}
	m.logger.Info("connected to node", zap.String("name", thisNode.Name))

	return nil
}

// Cleanup cleans up the handler.
func (m *Middleware) Cleanup() error {
	// Close the connection to the RedT blockchain node
	m.logger.Info("closing connection to RedT node")
	m.rt.Close()
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	return m.WShandler(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.NodeURL) {
			return d.ArgErr()
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.CleanerUpper          = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)

var upgrader = websocket.Upgrader{}

func (m Middleware) WShandler(w http.ResponseWriter, r *http.Request) error {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m.logger.Error("upgrading to WebSockets", zap.Error(err))
		return err
	}
	defer c.Close()

	// We are going to call the Geth API, with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Subscribe to receive notifications when new blocks are added to the blockchain
	// Each notification is received as a Header in the inputCh channel
	notificationCh := make(chan *ethertypes.Header)
	subscription, err := m.rt.EthClient().SubscribeNewHead(ctx, notificationCh)
	if err != nil {
		m.logger.Error("subscribing", zap.Error(err))
		return err
	}
	defer subscription.Unsubscribe()

	// Create a new statistics object
	stats := redt.NewStatistics(m.rt.AllValidators(), m.rt.Validators())

	// Wait for the first block
	currentHeader := <-notificationCh
	m.logger.Info("new block", zap.Int64("number", currentHeader.Number.Int64()))

	// Initialize statistics with this block
	stats.StatisticsForHeaderNew(currentHeader)

	// Loop forever until an error (probably the client disconnects)
	var rendered bytes.Buffer

	for {

		// We block here until a new header is received on the channel
		currentHeader := <-notificationCh
		m.logger.Info("new block", zap.Int64("number", currentHeader.Number.Int64()))

		// Get the signer data and accumulated statistics
		data := stats.StatisticsForHeaderNew(currentHeader)

		// Format the data into an HTML table
		rendered.Reset()
		err = m.t.ExecuteTemplate(&rendered, "table.html", data)
		if err != nil {
			m.logger.Error("executing template", zap.Error(err))
			return err
		}

		// Send the HTML table to the client via the WebSocket connection
		err = c.WriteMessage(websocket.TextMessage, rendered.Bytes())
		if err != nil {
			m.logger.Error("sending to browser websocket", zap.Error(err))
			return err
		}

	}

}
