package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-code/go-hockey-shop/config"
)

// InitDB connects to the database using the configuration settings
// retrieved by calling GetConfig and returns a pointer to a gorm.DB instance
// or an error if the connection could not be established.
func InitDB() (*gorm.DB, error) {
	config := config.GetConfig()

	// Connect to the database using the configuration settings
	db, err := gorm.Open("postgres", config.DBConnect)
	if err != nil {
		return nil, err
	}

	return db, nil
}
