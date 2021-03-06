package customer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//GetCustomerDetails to read customer info while signing in
func GetCustomerDetails() User {
	var user User
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.Name = strings.TrimRight(name, "\n")

	fmt.Println("Enter Address : ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.Address = strings.TrimRight(address, "\n")

	fmt.Println("Enter Email Address : ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.EmailAddress = strings.TrimRight(email, "\n")

	fmt.Println("Enter Contact Number : ")
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.ContactNumber = strings.TrimRight(number, "\n")

	fmt.Println("Enter Username : ")
	userName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.UserName = strings.TrimRight(userName, "\n")

	fmt.Println("Enter Password : ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading Customer details ")
		GetCustomerDetails()
	}
	user.Password = strings.TrimRight(password, "\n")

	return user
}

//GetLoginDetails for Login
func GetLoginDetails() LoginInfo {
	var user LoginInfo

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Username : ")
	userName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read username")
		GetLoginDetails()
	}
	user.UserName = strings.TrimRight(userName, "\n")

	fmt.Println("Enter Password : ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read password ")
		GetLoginDetails()
	}
	user.Password = strings.TrimRight(password, "\n")

	return user
}
