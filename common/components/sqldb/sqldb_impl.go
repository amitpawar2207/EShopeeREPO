package sqldb

import (
	"EShopeeREPO/common/factory"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlDriver sql driver
type MysqlDriver struct {
	db *sql.DB
}

//Init intializes and create a mysql connection
func (obj *MysqlDriver) Init(conf *SQLDBConfg) error {

	var err error

	str := conf.UserName + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DBName

	//open connection
	obj.db, err = sql.Open(conf.DriverName, str)
	if err != nil {
		return fmt.Errorf("Error while connecting mysql database ")
	} else {
		obj.db.SetMaxOpenConns(conf.MaxOpenCon)

		obj.db.SetMaxIdleConns(conf.MaxIdleCon)

		err = obj.db.Ping()
	}
	return nil
}

//Query executes the query in mysql db and returns the pointer to rows
func (obj *MysqlDriver) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := obj.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//Execute executes the sql query and returns the pointer to the result
func (obj *MysqlDriver) Execute(query string, args ...interface{}) (sql.Result, error) {
	result, err := obj.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Ping checks the connection
func (obj *MysqlDriver) Ping() error {
	err := obj.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//Close closes the mysql connection
func (obj *MysqlDriver) Close() error {
	err := obj.db.Close()
	if err != nil {
		return err
	}
	return nil
}

//GetSQLDBConfig return the configs
func GetSQLDBConfig() SQLDBConfg {
	var sdbcf SQLDBConfg
	sdbcf.DriverName = factory.MySQLDriveName
	sdbcf.UserName = factory.MySQLDBUserName
	sdbcf.Password = factory.MySQLDBPassword
	sdbcf.Host = factory.MySQLHost
	sdbcf.Port = factory.MySQLDBPort
	sdbcf.DBName = factory.MySQLDBName
	sdbcf.MaxOpenCon = factory.MaxOpenConnections
	sdbcf.MaxIdleCon = factory.MaxIdleConnections
	return sdbcf
}
