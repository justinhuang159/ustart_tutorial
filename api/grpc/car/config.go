package customerapi

import (
	"github.com/justinhuang159/ustart_tutorial/car"
)

// Config for auth api
type Config struct {
	CarCfg *car.Config
	Port       int //Recomended 5101
}
