package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sub-rat/contactbook-api/config"
	"github.com/sub-rat/contactbook-api/internal/model"
	"github.com/sub-rat/contactbook-api/pkg/db/postgres"
	"gorm.io/gorm"
)

var Version = "1.0.0"

type server struct {
	Router *mux.Router
	DB     *gorm.DB
}

func GetServer() *server {
	return &server{}
}

func (server *server) Run() {
	var err error
	// Load Config
	config := config.GetConfig()
	// Initialize database
	server.DB, err = postgres.NewPsqlDB(config)
	if err != nil {
		log.Fatal(err)
	}
	// server initialize
	server.initialize()
	// Run Server
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server.Router))
}

func (server *server) initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	server.Migrate()
}

func (server *server) Migrate() {
	server.DB.AutoMigrate(
		model.Contact{},
	)
}
