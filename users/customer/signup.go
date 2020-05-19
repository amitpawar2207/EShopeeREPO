package customer

import (
	"EShopeeREPO/common/components/sqldb"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//SignUp to register new customers
func SignUp() {
	var user User

	var db sqldb.MysqlDriver
	sdbcf := sqldb.GetSQLDBConfig()

	var err error

	err = db.Init(&sdbcf)
	if err != nil {
		fmt.Println(err)
	}

	rows, uerr := db.Query("SELECT username FROM users WHERE username=?", user.UserName)
	var hashedPassword []byte
	var usrName string
	if rows.Next() {
		rows.Scan(&usrName)
	}
	if usrName == user.UserName {
		fmt.Println("This user is already registered.")
		SignUp()
	}
	if uerr == nil {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("unable to incrypt password")
			log.Fatal(err)
		}
	} else {
		log.Fatal(uerr)
	}

	//create new user
	query := "INSERT INTO `users` (`name`, `contact_number`, `address`, `username`, `password`, `role`, `email_address`) VALUES (?,?,?,?,?,?,?);"

	sqlresult, eerr := db.Execute(query, user.Name, user.ContactNumber, user.Address, user.UserName, string(hashedPassword), user.Role, user.EmailAddress)
	if eerr != nil {
		log.Fatal(eerr)
	}
	num, rerr := sqlresult.RowsAffected()
	if num < 0 || rerr != nil {
		fmt.Println("No rows affected")
		log.Fatal(rerr)
	}

	var id int
	rows, serr := db.Query("SELECT id FROM users WHERE username=?", user.UserName)
	if serr != nil {
		log.Fatal(serr)
	}

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Welcomem to the EShopee !!!\n")

	Customer(id, db)

}
