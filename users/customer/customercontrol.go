package customer

import (
	"EShopeeREPO/common/components/sqldb"
	"EShopeeREPO/shop/category"
	"EShopeeREPO/shop/product"
	"fmt"
	"strconv"
	"time"
)

//Customer functionality
func Customer(custID int, db sqldb.MysqlDriver) error {
	fmt.Println("Select an Option :\n1. View Categories\n2. View Products\n3. View Cart")
	var i int
	_, serr := fmt.Scan(&i)
	if serr != nil {
		fmt.Println("Please Enter valid Input...")
		Customer(custID, db)
	}
	switch i {
	case 1:
		err := getCategoryList(custID, db)
		if err != nil {
			return err
		}
		Customer(custID, db)

	case 2:
		err := getProductList(custID, db)
		if err != nil {
			return err
		}
		Customer(custID, db)

	case 3:
		err := ViewCart(custID, false, db)
		if err != nil {
			return err
		}
		Customer(custID, db)
	default:
		Customer(custID, db)
	}
	cerr := db.Close()
	if cerr != nil {
		fmt.Errorf("Error while closing mysql connection ", cerr)
	}
	return nil
}

func getCategoryList(custID int, db sqldb.MysqlDriver) error {
	catList := make([]category.List, 0)
	var gcerr error
	catList, gcerr = category.GetCategoryList()
	if gcerr != nil {
		return gcerr
	}
	for index, item := range catList {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.CategoryName)
	}

	fmt.Println("Enter Category number to view products from the Category or press 00 to back:")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		getCategoryList(custID, db)
	}
	if i == 00 {
		Customer(custID, db)
	}

	prodList := make([]product.ProductList, 0)
	var perr error
	prodList, perr = product.GetProductsByCategory(catList[i-1].CategoryName)
	if perr != nil {
		return perr
	}
	for index, item := range prodList {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}
	return nil
}

//ViewCart shows cart details
func ViewCart(custID int, isAdmin bool, db sqldb.MysqlDriver) error {

	//prodmdb := mongodb.GetMongoDriver()
	cartRecords, gcerr := getCartRecords(custID, false, db)
	if gcerr != nil {
		return gcerr
	}
	products, gperr := GetProductsInfo(cartRecords)
	if gperr != nil {
		return gperr
	}
	totalAmount, finalAmount := ShowBill(cartRecords, products)
	if isAdmin {

	} else {
		fmt.Println("Enter product number to remove product from the cart or Press 11111 to checkout")
		var i int
		fmt.Scan(&i)
		switch {
		case i <= len(cartRecords):
			removeProductFromCart(cartRecords[i-1], db)
		case i == 11111:
			var order Order
			order.CustomerID = custID
			order.FinalAmount = finalAmount
			order.Discount = totalAmount - finalAmount

			order.CheckoutDate = time.Now().Local()
			placeOrder(order, db)
			updateCartRecords(cartRecords, order.CheckoutDate, db)
			fmt.Println("Order successfully placed.")
		default:
			fmt.Println("wrong Input")
			Customer(custID, db)
		}
		Customer(custID, db)
	}
	return nil
}

//ShowBill to display bill
func ShowBill(cartRecords []Cart, products map[string]product.Product) (float32, float32) {
	var totalAmount, finalAmount float32 = 0, 0
	fmt.Println("Cart Details : ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Index\tName\tBrand\tPrice\tQuantity\tAmount")
	fmt.Println("--------------------------------------------------------------------")
	for index, item := range cartRecords {
		if v, found := products[item.ProductID]; found {
			fmt.Println(index+1, v.Name, v.Brand, v.Price, v.Quantity, item.Amount)
			totalAmount = totalAmount + item.Amount
		}
	}
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Total amount - Rs. ", totalAmount)
	if totalAmount > 10000 {
		fmt.Println("Discount - Rs. ", 500)
		finalAmount = totalAmount - 500
	} else {
		fmt.Println("Discount - Rs. ", 0)
		finalAmount = totalAmount
	}
	fmt.Println("Final Amount - Rs. ", finalAmount)
	fmt.Println("--------------------------------------------------------------------")
	return totalAmount, finalAmount
}

func placeOrder(order Order, db sqldb.MysqlDriver) error {

	query := "INSERT INTO `orders` (`customer_id`, `final_amount`, `discount`,  `checkout_date`) VALUES (?,?,?,?);"
	rows, qerr := db.Execute(query, order.CustomerID, order.FinalAmount, order.Discount, order.CheckoutDate)
	if qerr != nil {
		return fmt.Errorf("Error while placing order ")
	}
	num, rerr := rows.RowsAffected()
	if num < 0 || rerr != nil {
		return fmt.Errorf("Error while placing order ")
	}
	return nil
}

func getCartRecords(custID int, checkoutFlag bool, db sqldb.MysqlDriver) ([]Cart, error) {

	cartRecords := make([]Cart, 0)

	var pID string
	var cID, pQuantity int
	var amt float32
	rows, serr := db.Query("SELECT cart_id, product_id, product_quantity, amount FROM cart WHERE customer_id=? AND checkout=?", custID, checkoutFlag)
	if serr != nil {
		return nil, fmt.Errorf("Error while fecting cart details")
	}
	for rows.Next() {
		err := rows.Scan(&cID, &pID, &pQuantity, &amt)
		if err != nil {
			return nil, fmt.Errorf("Error while fecting cart details")
		}
		cartRecords = append(cartRecords, Cart{CartID: cID, CustomerID: custID, ProductID: pID, ProductQuantity: pQuantity, Amount: amt})
	}
	return cartRecords, nil
}

//GetProductsInfo to create productinfo map
func GetProductsInfo(records []Cart) (map[string]product.Product, error) {

	products := make(map[string]product.Product)
	for _, item := range records {
		prod, err := product.GetProductDetails("", item.ProductID)
		if err != nil {
			return nil, err
		}
		products[prod.ID] = prod
	}
	return products, nil
}

func removeProductFromCart(cartRecord Cart, db sqldb.MysqlDriver) error {

	query := "DELETE FROM cart WHERE cart_id = ?;"
	rows, qerr := db.Execute(query, cartRecord.CartID)
	if qerr != nil {
		return fmt.Errorf("Error while removing item from cart")
	}

	num, rerr := rows.RowsAffected()
	if num <= 0 || rerr != nil {
		return fmt.Errorf("Error while removing item from cart")
	}

	err := ViewCart(cartRecord.CustomerID, false, db)
	if err != nil {
		return err
	}
	return nil
}

func updateCartRecords(cartRecords []Cart, checkoutDate time.Time, db sqldb.MysqlDriver) error {

	str := "("
	for i, id := range cartRecords {
		if i < (len(cartRecords) - 1) {
			str = str + strconv.Itoa(id.CartID) + ","
		} else {
			str = str + strconv.Itoa(id.CartID) + ")"
		}
	}
	query := "UPDATE cart SET checkout = ?, updatedat = ? WHERE `cart_id` IN " + str + ";"
	rows, qerr := db.Execute(query, true, checkoutDate)
	if qerr != nil {
		return fmt.Errorf("Error while updating item in cart")
	}

	num, rerr := rows.RowsAffected()
	if num <= 0 || rerr != nil {
		return fmt.Errorf("Error while updating item in cart")
	}

	return nil
}
