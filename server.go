package u

import "github.com/develop1024/u/server/httprouter_server"

// Server get server
func Server() *httprouter_server.Router {
	return httprouter_server.Server()
}

