package sqlstore

import (
	"context"

	// "github.com/olivere/elastic"
)

// SQLStore implements the storage interface for the customer module
type sqlStore struct {
	db *sql.DB
	recordTableName string
}

// New returns a new sqlstore service
func New(cfg *Config) (*SQLStore, error) {
	_ = pq.Efatal
	connString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port
	)

	client, err := sql.open(cfg.DriverName, connString)
	if err != nil 

	return dbConn, err
}
