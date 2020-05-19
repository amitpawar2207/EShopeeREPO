package mongodb

import (
	"EShopeeREPO/common/factory"
	"log"
)

//GetMongoDriver to get mongo connection
func GetMongoDriver() MongoDriver {
	var mdbC MDBConfig
	mdbC.DBName = factory.MongoDBDatabaseName
	mdbC.URL = factory.MongoDBURL

	var prodmdb MongoDriver

	cerr := prodmdb.Init(&mdbC)
	if cerr != nil {
		log.Fatal(cerr)
	}
	return prodmdb
}
