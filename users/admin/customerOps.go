package admin

import (
	"EShopeeREPO/common/components/sqldb"
	"EShopeeREPO/users/customer"
	"fmt"
	"log"
)

func viewCustomerDetails(user customer.User) {
	fmt.Println("Customer details :")
	fmt.Println("Name : ", user.Name)
	fmt.Println("Email ID : ", user.EmailAddress)
	fmt.Println("Contact Number : ", user.ContactNumber)
}

func customerWork(db sqldb.MysqlDriver) {
	userList := getCustomerList()
	for index, user := range userList {
		fmt.Println(index+1, ".", user.Name)
	}
	fmt.Println("Press 00 to Back Or Enter number to view details :")
	var i int
	fmt.Scan(&i)
	switch i {
	case 00:
		AdminWork()
	case i:
		viewCustomerDetails(userList[i-1])
		fmt.Println("1. View customer cart\n2. View customer bills\n00. Back")
		var j int
		fmt.Scan(&j)
		switch j {
		case 00:
			AdminWork()
		case 1:
			customer.ViewCart(userList[i-1].ID, true, db)
			customerWork(db)
		case 2:
			ViewCutomerBills(userList[i-1].ID)
		}
	}

}

func getCustomerList() []customer.User {
	fmt.Println("Customer List:")
	var mdb sqldb.MysqlDriver
	sdbcf := sqldb.GetSQLDBConfig()

	err := mdb.Init(&sdbcf)
	if err != nil {
		fmt.Println(err)
	}
	userList := make([]customer.User, 0)
	query := "SELECT id, name, contact_number, address, email_address FROM users WHERE role=?;"

	rows, serr := mdb.Query(query, "customer")
	if serr != nil {
		log.Fatal(serr)
	}

	for rows.Next() {
		var name, contactNumber, address, email string
		var usr customer.User
		var id int
		err := rows.Scan(&id, &name, &contactNumber, &address, &email)
		if err != nil {
			fmt.Println("Error while scanning name")
		}
		usr.ID = id
		usr.Name = name
		usr.EmailAddress = email
		usr.Address = address
		usr.ContactNumber = contactNumber

		userList = append(userList, usr)
	}

	return userList
}

//ViewCutomerBills shows details of bills
func ViewCutomerBills(custID int) {

	var mdb sqldb.MysqlDriver
	sdbcf := sqldb.GetSQLDBConfig()

	err := mdb.Init(&sdbcf)
	if err != nil {
		fmt.Println(err)
	}

	oQuery := "SELECT checkout_date FROM orders WHERE customer_id=? order by checkout_date;"
	orows, serr := mdb.Query(oQuery, custID)
	if serr != nil {
		log.Fatal(serr)
	}

	dates := make([]string, 0)
	for orows.Next() {
		var date string
		orows.Scan(&date)
		dates = append(dates, date)
	}

	fmt.Println("dates - ")
	fmt.Println(dates)
	for index, date := range dates {
		fmt.Println(index+1, ". billed date -", date)
	}
	fmt.Println("enter the bill number to view bill details : ")
	var k int
	fmt.Scan(&k)

	showbill(dates[k-1], mdb)

}

func showbill(date string, mdb sqldb.MysqlDriver) {
	data := make([]customer.Cart, 0)
	cQuery := "SELECT cart_id, product_id, product_quantity, amount FROM cart WHERE  checkout = ? AND updatedat = ? ;"
	crows, cerr := mdb.Query(cQuery, true, date)
	if cerr != nil {
		log.Fatal(cerr)
	}

	var cID, prodQuant int
	var pID string
	var amount float32
	for crows.Next() {
		var cData customer.Cart
		crows.Scan(&cID, &pID, &prodQuant, &amount)
		cData.CartID = cID
		cData.ProductID = pID
		cData.ProductQuantity = prodQuant
		cData.Amount = amount
		data = append(data, cData)
	}

	fmt.Println("data - ")
	fmt.Println(data)
	products := customer.GetProductsInfo(data)
	customer.ShowBill(data, products)
}
