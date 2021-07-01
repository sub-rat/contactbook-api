package postgres

import (
	"fmt"
	"log"

	"github.com/sub-rat/contactbook-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(c *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DB.Host, c.DB.Username, c.DB.Password, c.DB.Name, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to %s database: %v", c.DB.Dialect, err)
	}
	return db, nil
}
