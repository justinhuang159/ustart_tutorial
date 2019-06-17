package storage 

import {
	"github.com/justinhuang159/ustart_tutorial/record/storage"
}
func NewSQL(config *Config) (Storage, error){
	strg, err := sqlstore.new(Config.SQLConfig)
	return strg, err
}