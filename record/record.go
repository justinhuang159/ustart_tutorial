package record

import (
	"github.com/sea350/ustart_tutorial/customer/storage"
)

//Customer is an implementation of the customer service as defined in service.proto
type Record struct {
	strg          storage.Storage
}

// New returns a new Eclient customer service
func New(cfg *Config) (*Record, error) {
	// if cfg.useDummy

	storg, err := storage.NewSQL(cfg.StorageConfig)

	cust := &Record{
		strg: storg,
	}

	return cust, err
}
