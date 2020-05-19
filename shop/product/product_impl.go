package product

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/factory"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//ProductList exported
type ProductList struct {
	Name string `bson:"name"`
}

//Create product
func (obj *Product) Create() {

	prodmdb := mongodb.GetMongoDriver()

	ierr := prodmdb.Insert(factory.ProductCollection, &obj)
	if ierr != nil {
		log.Fatal(ierr)
	}
	fmt.Println("New product", obj.Name, "created")
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(str, "\n")

}

//GetProductList all
func GetProductList() []ProductList {

	prodmdb := mongodb.GetMongoDriver()

	selectField := bson.M{
		"_id":  0,
		"name": 1,
	}

	result, err := prodmdb.Find(factory.ProductCollection, nil, selectField, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	dbresult := make([]ProductList, 0)
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

//UpdateProductData of product
func UpdateProductData(prod Product, oldName string) {
	prodmdb := mongodb.GetMongoDriver()
	fmt.Println("Select option to edit : ")
	fmt.Println("1. Product Name\n2. Brand\n3. Category\n4. Description\n5. Quantity\n6. Price")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		UpdateProductData(prod, oldName)
	}

	switch i {
	case 1:
		fmt.Println("Enter new name : ")
		prod.Name = readString()
	case 2:
		fmt.Println("Enter new brand name : ")
		prod.Brand = readString()
	case 3:
		fmt.Println("Enter new Category name : ")
		prod.Category = readString()
	case 4:
		fmt.Println("Enter new product description : ")
		prod.Description = readString()
	case 5:
		fmt.Println("Enter new quantity : ")
		var q int
		fmt.Scan(&q)
		prod.Quantity = q
	case 6:
		fmt.Println("Enter new price : ")
		var price float32
		_, perr := fmt.Scanf("%f", &price)
		if perr != nil {
			log.Fatal(perr)
		}
		prod.Price = price
	default:
		UpdateProductData(prod, oldName)
	}
	query := map[string]interface{}{
		"name": oldName,
	}

	value := bson.M{
		"$set": bson.M{
			"name":        prod.Name,
			"brand":       prod.Brand,
			"category":    prod.Category,
			"description": prod.Description,
			"quantity":    prod.Quantity,
			"price":       prod.Price,
		},
	}
	prodmdb.Update(factory.ProductCollection, query, value)
}

//RemoveProduct from Documents
func RemoveProduct(productName string) {

	prodmdb := mongodb.GetMongoDriver()

	whereQuery := bson.M{
		"name": productName,
	}

	err := prodmdb.Remove(factory.ProductCollection, whereQuery)
	if err != nil {
		log.Fatal(err)
	}
}

//GetProductsByCategory product list
func GetProductsByCategory(categoryName string) []ProductList {

	prodmdb := mongodb.GetMongoDriver()

	whereQuery := bson.M{
		"category": categoryName,
	}

	selectQeury := bson.M{
		"_id":  0,
		"name": 1,
	}

	result, err := prodmdb.Find(factory.ProductCollection, whereQuery, selectQeury, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	dbresult := make([]ProductList, 0)
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

//GetProductDetails return details about Product
func GetProductDetails(productName, productID string) Product {

	prodmdb := mongodb.GetMongoDriver()

	var whereQuery bson.M
	if productID == "" {
		whereQuery = bson.M{
			"name": productName,
		}
	} else {
		whereQuery = bson.M{
			"id": productID,
		}
	}

	selectQeury := bson.M{
		"id":          1,
		"name":        1,
		"brand":       1,
		"category":    1,
		"description": 1,
		"price":       1,
		"quantity":    1,
	}

	result, err := prodmdb.Find(factory.ProductCollection, whereQuery, selectQeury, 0, 1)
	if err != nil {
		log.Fatal(err)
	}

	dbresult := make([]Product, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		log.Fatal(merr)
	}

	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		log.Fatal(umerr)
	}

	return dbresult[0]
}
