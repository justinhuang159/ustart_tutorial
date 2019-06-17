package recordapi

import (
	"strconv"

	"github.com/justinhuang159/ustart_tutorial/record"
)

//GPRCAPI is the GRPC API implementation for customer
type GPRCAPI struct {
	record *record.Record
	port string
}

// New creates a new customer api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	prof, err := record.New(cfg.RecordCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		prof: prof,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
