package customer

import (
	"EShopeeREPO/common/components/sqldb"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Login existing customers
func Login() {

	user := GetLoginDetails()

	var db sqldb.MysqlDriver
	sdbcf := sqldb.GetSQLDBConfig()

	ierr := db.Init(&sdbcf)
	if ierr != nil {
		fmt.Println(ierr)
	}

	var databasePassword, role, name string
	var id int
	rows, err := db.Query("SELECT id,password,role,name FROM users WHERE username=?", user.UserName)
	if err != nil {
		fmt.Println("Login Failed")
		Login()
	}

	if rows.Next() {
		serr := rows.Scan(&id, &databasePassword, &role, &name)
		if serr != nil {
			fmt.Println("error while scanning password")
			Login()
		}
	} else {
		fmt.Println("User is not registered.")
	}

	cerr := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Password))
	if cerr != nil {
		fmt.Println("Incorrect Password")
		Login()
	} else {
		fmt.Println("Welcome, to Shopee", name)
		Customer(id, db)
	}

}
