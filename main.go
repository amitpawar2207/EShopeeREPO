package main

import (
	"EShopeeREPO/users/admin"
	"EShopeeREPO/users/customer"
	"fmt"
)

func main() {
	HomePage()
}

//HomePage is the entry page for the Eshop
func HomePage() {
	fmt.Println("Select an Option :\n1. Admin Login\n2. Customer SignUp\n3. Customer Login")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		HomePage()
	}
	switch i {
	case 1:
		admin.AdminWork()
	case 2:
		customer.SignUp()
	case 3:
		customer.Login()

	}
}
