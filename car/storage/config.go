package storage

import (
	"github.com/justinhuang159/ustart_tutorial/car/storage/elastic"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed customer server
type Config struct {
	useDummy      bool
	ElasticConfig *elasticstore.Config
	// SQLConfig     *sqlstore.Config
}

// ESNewConfig returns a default config object
func ESNewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}
