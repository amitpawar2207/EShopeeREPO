package category

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/factory"
	"encoding/json"
	"log"

	"gopkg.in/mgo.v2/bson"
)

//List of categories
type List struct {
	CategoryName string `bson:"categoryname"`
}

//Create category
func (obj *Category) Create() error {

	catmdb := mongodb.GetMongoDriver()

	count, err := catmdb.Count(factory.CategoryCollection)
	if err != nil {
		log.Fatal(err)
	}

	obj.CategoryID = count + 1

	ierr := catmdb.Insert(factory.CategoryCollection, &obj)
	if ierr != nil {
		return ierr
	}

	return nil
}

//GetCategoryList all
func GetCategoryList() []List {

	catmdb := mongodb.GetMongoDriver()

	selectField := bson.M{
		"_id":          0,
		"categoryname": 1,
	}

	result, err := catmdb.Find(factory.CategoryCollection, nil, selectField, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	dbresult := make([]List, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		log.Fatal(merr)
	}
	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		log.Fatal(umerr)
	}

	return dbresult

}

//RemoveCategory from Documents
func RemoveCategory(categoryName string) {
	catmdb := mongodb.GetMongoDriver()

	whereQuery := bson.M{
		"categoryname": categoryName,
	}

	err := catmdb.Remove(factory.CategoryCollection, whereQuery)
	if err != nil {
		log.Fatal(err)
	}
}
