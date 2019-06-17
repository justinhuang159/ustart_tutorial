package recordapi

import (
	"github.com/justinhuang159/ustart_tutorial/record"
)

// Config for auth api
type Config struct {
	RecordCfg *record.Config
	Port       int //Recomended 5101
}
