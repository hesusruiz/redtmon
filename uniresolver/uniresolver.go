package uniresolver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hesusruiz/redtmon/contracts/resolver"
	"github.com/hesusruiz/redtmon/redt"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(Uniresolver{})
	httpcaddyfile.RegisterHandlerDirective("uniresolver", parseCaddyfile)
}

// Uniresolver implements an HTTP handler that resolves DIDs
// using a SmartContract deployed in Alastria RedT.
// It implements the EBSI 'Get a DID Document' API at
// https://api-pilot.ebsi.eu/docs/apis/did-registry/latest#/operations/get-did-registry-v3-identifier
type Uniresolver struct {
	// The location of the Geth node
	NodeURL string `json:"nodeurl,omitempty"`

	// The connection to the RedT blockchain node
	rt *redt.RedTNode

	// The Resolver Smart Contract wrapper used to call the actual contract in the blockchain
	res_session resolver.ResolverCallerSession

	// The logger
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (Uniresolver) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.uniresolver",
		New: func() caddy.Module { return new(Uniresolver) },
	}
}

// Provision implements caddy.Provisioner
func (m *Uniresolver) Provision(ctx caddy.Context) (err error) {
	m.logger = ctx.Logger()

	m.logger.Info("connecting to RedT node at", zap.String("url", m.NodeURL))

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
	m.logger.Info("connected to RedT node", zap.String("name", thisNode.Name))

	// Hardcoded deployment address of the Smart Contract
	var resolverContractAddr = "0x560f80EdEA19a176Bd16584C520c0e275e67C714"

	// Create a read-only instance of the Resolver pointing at the deployment address
	res, err := resolver.NewResolverCaller(common.HexToAddress(resolverContractAddr), m.rt.EthClient())
	if err != nil {
		m.logger.Error("error binding")
		return err
	}

	// Create a session instance to avoid having to specify call parameters later
	res_session := resolver.ResolverCallerSession{
		Contract: res,
	}

	// And store the session instance in the Uniresolver struct
	m.res_session = res_session

	return nil
}

// Cleanup cleans up the handler.
func (m *Uniresolver) Cleanup() error {
	// Close the connection to the RedT blockchain node
	m.logger.Info("closing connection to RedT node")
	if m.rt != nil {
		m.rt.Close()
	}
	return nil
}

// Validate implements caddy.Validator.
func (m *Uniresolver) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Uniresolver) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	return m.UniversalResolver(w, r)

}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Uniresolver) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.NodeURL) {
			return d.ArgErr()
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Uniresolver
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Uniresolver)(nil)
	_ caddy.CleanerUpper          = (*Uniresolver)(nil)
	_ caddy.Validator             = (*Uniresolver)(nil)
	_ caddyhttp.MiddlewareHandler = (*Uniresolver)(nil)
	_ caddyfile.Unmarshaler       = (*Uniresolver)(nil)
)

// ErrorResponse represents the JSON object to return when an error happens, according to
// https://api-pilot.ebsi.eu/docs/apis/did-registry/latest#/schemas/ProblemDetails
type ErrorResponse struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// UniversalResolver uses the Alastria RedT network to resolve a DID
func (m Uniresolver) UniversalResolver(w http.ResponseWriter, r *http.Request) error {

	m.logger.Sugar().Infoln(r.URL.Path)

	// Get all the components of the full path without the fragments or query parameters
	// The last one should be the DID to resolve, but we require at least a path component before the DID
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 2 {
		return fmt.Errorf("invalid path: %v", r.URL.Path)
	}

	// Get the DID. We accept any format and try to resolve.
	// The downstream routines are robust enough to resolve to a null DIDDocument if not found.
	did := segments[len(segments)-1]

	// The Smart Contract stores hashes for efficiency.
	didKecka := crypto.Keccak256([]byte(did))
	// We need to provide a [32]byte array instead of a slice
	var didHash [32]byte
	copy(didHash[:], didKecka)

	// We need two steps:
	// 1. Retrieve the Node hash in the ENS tree that 'owns' the DID
	// 2. Get the DIDDocument associated to the entity in the ENS Node

	// Perform step 1
	node, err := m.res_session.NodeFromDID(didHash)
	if err != nil {
		m.logger.Error("error in NodeFromDID")
		return err
	}

	// And now step 2
	resolved, err := m.res_session.AlaDIDPublicEntity(node)
	if err != nil {
		m.logger.Error("error in AlaDIDPublicEntity")
		return err
	}

	// Log info for debugging
	m.logger.Sugar().Infoln("Active", resolved.Active)
	m.logger.Sugar().Infoln("DomainName", resolved.DomainName)
	m.logger.Sugar().Infoln("DIDDocument", resolved.DIDDocument)

	// We should respond to caller with JSON Mimetype
	w.Header().Add("Content-Type", "application/json")

	// Respond with error if DID not found
	if !resolved.Active || len(resolved.DIDDocument) == 0 {
		result := ErrorResponse{}
		result.Type = "about:blank"
		result.Title = "Identifier Not Found"
		result.Status = http.StatusBadRequest
		result.Detail = "Identifier " + did + " not found"

		w.WriteHeader(http.StatusBadRequest)
		out, _ := json.MarshalIndent(result, "", "  ")
		_, err = w.Write(out)
		return err
	}

	// Write the DIDDocument which is already a JSON object
	_, err = w.Write([]byte(resolved.DIDDocument))

	return err
}
