package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/hesusruiz/redtmon/caddymodule/redtmon"
)

func main() {

	caddycmd.Main()

}
