package sqldb

import "database/sql"

//MSDBInterface is for all the mysql db operations
type MSDBInterface interface {
	Init(confg *SQLDBConfg) error
	Query(string, ...interface{}) (*sql.Rows, error)
	Execute(string, ...interface{}) (sql.Result, error)
	Ping() error
	Close() error
}
