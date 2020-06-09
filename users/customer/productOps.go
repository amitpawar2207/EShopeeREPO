package customer

import (
	"EShopeeREPO/common/components/mongodb"
	"EShopeeREPO/common/components/sqldb"
	"EShopeeREPO/common/factory"
	"EShopeeREPO/shop/product"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func getProductList(custID int, db sqldb.MysqlDriver) error {
	list := make([]product.ProductList, 0)
	list, err := product.GetProductList()
	if err != nil {
		return err
	}
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}
	fmt.Println("Enter product number to view product details :")
	var i int
	fmt.Scan(&i)
	showProductDetails(list[i-1].Name, custID, db)
	return nil
}

func showProductDetails(productName string, custID int, db sqldb.MysqlDriver) error {
	product, gperr := product.GetProductDetails(productName, "")
	if gperr != nil {
		return gperr
	}
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
		fmt.Println("Enter 1 to add this product to cart or press 00 to Back")
		var i int
		var quantity int
		_, serr := fmt.Scan(&i)
		if serr != nil {
			showProductDetails(productName, custID, db)
		}
		switch i {
		case 1:
			quantity = acceptQuanity(product.Quantity)
			addProductToCart(product, quantity, custID, db)
		case 00:
			getProductList(custID, db)
		default:
			getProductList(custID, db)
		}
	}
	Customer(custID, db)
	return nil
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

func addProductToCart(product product.Product, quanity, custID int, db sqldb.MysqlDriver) error {

	//check for existing products to avoid duplicate entries

	var cID, pQuant int
	row, aerr := db.Query("SELECT cart_id, product_quantity from cart where customer_id=? AND product_id=? AND checkout=?", custID, product.ID, false)
	if aerr != nil {
		return fmt.Errorf("Error while adding product to cart ")
	}
	if row.Next() {
		serr := row.Scan(&cID, &pQuant)
		if serr != nil {
			return fmt.Errorf("Error while adding product to cart ")
		}
		newPQuant := quanity + pQuant
		newAmt := float32(newPQuant) * product.Price
		query := "UPDATE cart SET product_quantity = ?, amount = ?, updatedat = ? WHERE cart_id = ? AND product_id = ?"
		res, eerr := db.Execute(query, newPQuant, newAmt, time.Now(), cID, product.ID)
		if eerr != nil {
			return fmt.Errorf("Error while adding product to cart ")
		}
		num, rerr := res.RowsAffected()
		if num < 0 || rerr != nil {
			return fmt.Errorf("Error while adding product to cart ")
		}
	} else {
		pQuant = quanity
		amount := product.Price * float32(quanity)
		query := "INSERT INTO `cart` (`customer_id`, `product_id`, `product_quantity`, `amount`, `checkout`, `updatedat`) VALUES (?,?,?,?,?,?);"
		res, eerr := db.Execute(query, custID, product.ID, quanity, amount, false, time.Now())
		if eerr != nil {
			return fmt.Errorf("Error while adding product to cart ")
		}
		num, rerr := res.RowsAffected()
		if num < 0 || rerr != nil {
			return fmt.Errorf("Error while adding product to cart ")
		}
	}

	prodmdb, mgerr := mongodb.GetMongoDriver()
	if mgerr != nil {
		return mgerr
	}

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
		return fmt.Errorf("Error while adding product to cart ")
	}
	fmt.Println("Product added to the cart.")
	return nil
}
