package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	
	// plug in plugins here, for example:
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/captncraig/cors"
	_ "github.com/epicagency/caddy-expires"
	_ "github.com/echocat/caddy-filter"
	_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/tarent/loginsrv/caddy"
	_ "github.com/hacdias/caddy-minify"
	_ "github.com/Xumeiquer/nobots"
	_ "github.com/miekg/caddy-prometheus"
	_ "github.com/xuqingfeng/caddy-rate-limit"
	_ "github.com/freman/caddy-reauth"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}