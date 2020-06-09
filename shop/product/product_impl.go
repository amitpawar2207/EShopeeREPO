package product

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/factory"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//ProductList exported
type ProductList struct {
	Name string `bson:"name"`
}

//Create product
func (obj *Product) Create() error {

	prodmdb, merr := mongodb.GetMongoDriver()
	if merr != nil {
		return merr
	}

	ierr := prodmdb.Insert(factory.ProductCollection, &obj)
	if ierr != nil {
		return fmt.Errorf("Error while inserting product data in product collection ", ierr)
	}
	fmt.Println("New product", obj.Name, "created")
	return nil
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error while reading string data")
	}
	return strings.TrimRight(str, "\n"), nil

}

//GetProductList all
func GetProductList() ([]ProductList, error) {

	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return nil, fmt.Errorf("Error while initializing mongo connection ", mgerr)
	}

	selectField := bson.M{
		"_id":  0,
		"name": 1,
	}

	result, err := prodmdb.Find(factory.ProductCollection, nil, selectField, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("Error while fetching product data from product collection ", err)
	}

	dbresult := make([]ProductList, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		return nil, fmt.Errorf("Error while marshaling product data ", merr)
	}
	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		return nil, fmt.Errorf("Error while unmarshaling product data ", merr)
	}

	return dbresult, nil
}

//UpdateProductData of product
func UpdateProductData(prod Product, oldName string) error {
	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return mgerr
	}
	fmt.Println("Select option to edit : ")
	fmt.Println("1. Product Name\n2. Brand\n3. Category\n4. Description\n5. Quantity\n6. Price")
	var i int
	var err error
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		UpdateProductData(prod, oldName)
	}

	switch i {
	case 1:
		fmt.Println("Enter new name : ")
		prod.Name, err = readString()
		if err != nil {
			fmt.Println(err)
			UpdateProductData(prod, oldName)
		}
	case 2:
		fmt.Println("Enter new brand name : ")
		prod.Brand, err = readString()
		if err != nil {
			fmt.Println(err)
			UpdateProductData(prod, oldName)
		}
	case 3:
		fmt.Println("Enter new Category name : ")
		prod.Category, err = readString()
		if err != nil {
			fmt.Println(err)
			UpdateProductData(prod, oldName)
		}
	case 4:
		fmt.Println("Enter new product description : ")
		prod.Description, err = readString()
		if err != nil {
			fmt.Println(err)
			UpdateProductData(prod, oldName)
		}
	case 5:
		fmt.Println("Enter new quantity : ")
		var q int
		_, serr := fmt.Scan(&q)
		if serr != nil {
			fmt.Println(fmt.Errorf("Error while scanning quantity "))
			UpdateProductData(prod, oldName)
		}
		prod.Quantity = q
	case 6:
		fmt.Println("Enter new price : ")
		var price float32
		_, perr := fmt.Scanf("%f", &price)
		if perr != nil {
			fmt.Println(fmt.Errorf("Error while scanning price "))
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
	uerr := prodmdb.Update(factory.ProductCollection, query, value)
	if uerr != nil {
		return fmt.Errorf("Error while updating product data ", uerr)
	}
	return nil
}

//RemoveProduct from Documents
func RemoveProduct(productName string) error {

	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return mgerr
	}
	whereQuery := bson.M{
		"name": productName,
	}

	err := prodmdb.Remove(factory.ProductCollection, whereQuery)
	if err != nil {
		fmt.Errorf("Error while removing product from product collection ", err)
	}

	return nil
}

//GetProductsByCategory product list
func GetProductsByCategory(categoryName string) ([]ProductList, error) {

	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return nil, mgerr
	}
	whereQuery := bson.M{
		"category": categoryName,
	}

	selectQeury := bson.M{
		"_id":  0,
		"name": 1,
	}

	result, err := prodmdb.Find(factory.ProductCollection, whereQuery, selectQeury, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("Error while fetching product data from product collection ", err)
	}

	dbresult := make([]ProductList, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		return nil, fmt.Errorf("Error while marshaling product data ", merr)
	}
	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		return nil, fmt.Errorf("Error while marshaling product data ", umerr)
	}
	return dbresult, nil
}

//GetProductDetails return details about Product
func GetProductDetails(productName, productID string) (Product, error) {

	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return Product{}, fmt.Errorf("Error while initializing mongo connection ", mgerr)
	}
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
		return Product{}, fmt.Errorf("Error while fetching product data from product collection ", err)
	}

	dbresult := make([]Product, 0)
	byteData, merr := json.Marshal(result)
	if merr != nil {
		return Product{}, fmt.Errorf("Error while marshaling product data ", merr)
	}

	umerr := json.Unmarshal(byteData, &dbresult)
	if umerr != nil {
		return Product{}, fmt.Errorf("Error while unmarshaling product data ", umerr)
	}

	return dbresult[0], nil
}
