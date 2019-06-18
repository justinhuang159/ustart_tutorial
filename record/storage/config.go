package record

import "github.com/justinhuang159/ustart_tutorial/record/storage"

// Config determines the runtime behavior of the SQL-backed record server
type sqlStore struct {
	db *sql.DB
	recordTableName string
}