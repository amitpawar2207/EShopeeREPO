package customer

import (
	"EShopeeREPO/common/components/sqldb"
	"fmt"
	"log"

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
		fmt.Println("No rows affected")
		log.Fatal(err)
	}

	if rows.Next() {
		serr := rows.Scan(&id, &databasePassword, &role, &name)
		if serr != nil {
			fmt.Println("error while scanning password")
			log.Fatal(serr)
		}
	} else {
		fmt.Println("User is not registered.")
	}

	cerr := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Password))
	if cerr != nil {
		fmt.Println("error while comparing passwords")
	} else {
		fmt.Println("Welcome, to Shopee", name)
		Customer(id, db)
	}

}
