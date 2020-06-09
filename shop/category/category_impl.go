package category

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/factory"
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//List of categories
type List struct {
	CategoryName string `bson:"categoryname"`
}

//Create category
func (obj *Category) Create() error {

	catmdb, merr := mongodb.GetMongoDriver()
	if merr != nil {
		return merr
	}

	ierr := catmdb.Insert(factory.CategoryCollection, &obj)
	if ierr != nil {
		return ierr
	}

	return nil
}

//GetCategoryList all
func GetCategoryList() ([]List, error) {

	catmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return nil, mgerr
	}

	selectField := bson.M{
		"_id":          0,
		"categoryname": 1,
	}

	result, err := catmdb.Find(factory.CategoryCollection, nil, selectField, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("Error while finding categories in category collection")
	}

	dbresult := make([]List, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		return nil, fmt.Errorf("Error while marshaling data ", merr)
	}
	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		return nil, fmt.Errorf("Error while unmarshaling data ", umerr)
	}

	return dbresult, nil

}

//RemoveCategory from Documents
func RemoveCategory(categoryName string) error {
	catmdb, merr := mongodb.GetMongoDriver()
	if merr != nil {
		return merr
	}

	whereQuery := bson.M{
		"categoryname": categoryName,
	}

	err := catmdb.Remove(factory.CategoryCollection, whereQuery)
	if err != nil {
		return fmt.Errorf("Error while removing data ", err)
	}

	return nil
}
