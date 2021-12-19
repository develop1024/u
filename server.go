package u

import (
	"github.com/develop1024/u/server/hr"
)

// Server get server
func Server() *hr.Router {
	return hr.Server()
}
