package record

import"github.com/justinhuang159/ustart_tutorial/record/storage"

type Config struct {
	StorageConfig *storage.Config
}

func New(cfg *Config) (*Customer)