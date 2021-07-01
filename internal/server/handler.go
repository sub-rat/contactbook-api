package server

import "github.com/sub-rat/contactbook-api/internal/contact"

func (server *server) initializeRoutes() {
	contact.RegisterHandler(server.Router)
}
