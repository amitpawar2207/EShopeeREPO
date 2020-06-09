package mongodb

import (
	"EShopeeREPO/common/factory"
	"fmt"
)

//GetMongoDriver to get mongo connection
func GetMongoDriver() (MongoDriver, error) {
	var mdbC MDBConfig
	mdbC.DBName = factory.MongoDBDatabaseName
	mdbC.URL = factory.MongoDBURL

	var prodmdb MongoDriver

	cerr := prodmdb.Init(&mdbC)
	if cerr != nil {
		return prodmdb, fmt.Errorf("Error while initializing mongo connection ", cerr)
	}
	return prodmdb, nil
}
