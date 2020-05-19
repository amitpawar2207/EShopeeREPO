package customer

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/components/sqldb"
	"EShopeeREPO/common/factory"
	"EShopeeREPO/shop/product"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func getProductList(custID int, db sqldb.MysqlDriver) {
	list := make([]product.ProductList, 0)
	list = product.GetProductList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}
	fmt.Println("Enter product number to view product details :")
	var i int
	fmt.Scan(&i)
	showProductDetails(list[i-1].Name, custID, db)

}

func showProductDetails(productName string, custID int, db sqldb.MysqlDriver) {
	product := product.GetProductDetails(productName, "")
	fmt.Println("Product Details :")
	fmt.Println("Product Name : ", product.Name)
	fmt.Println("Product Brand : ", product.Brand)
	fmt.Println("Product Category : ", product.Category)
	fmt.Println("Product Description : ", product.Description)
	fmt.Println("Product Price : ", product.Price)
	fmt.Println("Product Quantity : ", product.Quantity)
	if product.Quantity == 0 {
		fmt.Println("This Product is currently unavailable.")
	} else {
		fmt.Println("Enter 1 to add this product to cart.")
		var i int
		fmt.Scan(&i)
		var quantity int
		if i == 1 {
			quantity = acceptQuanity(product.Quantity)
			addProductToCart(product, quantity, custID, db)
		}
	}
	Customer(custID, db)
}

func acceptQuanity(availableQuantity int) int {
	fmt.Println("Enter the quanity of products to add : ")
	var quantity int
	fmt.Scan(&quantity)
	if quantity > availableQuantity {
		fmt.Println("Please enter required quantity less than available quanity.")
		acceptQuanity(availableQuantity)
	}
	return quantity
}

func addProductToCart(product product.Product, quanity, custID int, db sqldb.MysqlDriver) {

	//check for existing products to avoid duplicate entries

	var cID, pQuant int
	row, aerr := db.Query("SELECT cart_id, product_quantity from cart where customer_id=? AND product_id=? AND checkout=?", custID, product.ID, false)
	if aerr != nil {
		log.Fatal(aerr)
	}
	if row.Next() {
		serr := row.Scan(&cID, &pQuant)
		if serr != nil {
			log.Fatal(serr)
		}
		newPQuant := quanity + pQuant
		newAmt := float32(newPQuant) * product.Price
		query := "UPDATE cart SET product_quantity = ?, amount = ?, updatedat = ? WHERE cart_id = ? AND product_id = ?"
		res, eerr := db.Execute(query, newPQuant, newAmt, time.Now(), cID, product.ID)
		if eerr != nil {
			log.Fatal(eerr)
		}
		num, rerr := res.RowsAffected()
		if num < 0 || rerr != nil {
			fmt.Println("No rows affected")
			log.Fatal(rerr)
		}
	} else {
		pQuant = quanity
		amount := product.Price * float32(quanity)
		query := "INSERT INTO `cart` (`customer_id`, `product_id`, `product_quantity`, `amount`, `checkout`, `updatedat`) VALUES (?,?,?,?,?,?);"
		res, eerr := db.Execute(query, custID, product.ID, quanity, amount, false, time.Now())
		if eerr != nil {
			log.Fatal(eerr)
		}
		num, rerr := res.RowsAffected()
		if num < 0 || rerr != nil {
			fmt.Println("No rows affected")
			log.Fatal(rerr)
		}
	}

	prodmdb := mongodb.GetMongoDriver()

	//update product quantity
	mquery := map[string]interface{}{
		"id": product.ID,
	}
	updatedQuantity := product.Quantity - pQuant
	value := bson.M{
		"$set": bson.M{
			"quantity": updatedQuantity,
		},
	}
	uerr := prodmdb.Update(factory.ProductCollection, mquery, value)
	if uerr != nil {
		log.Fatal(uerr)
	}
	fmt.Println("Product added to the cart.")
}
