package admin

import (
	"EShopeeREPO/common/components/sqldb"
	"fmt"
	"log"
)

//AdminWork funcs
func AdminWork() {
	var db sqldb.MysqlDriver
	sdbcf := sqldb.GetSQLDBConfig()
	err := db.Init(&sdbcf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Select an Option :\n1. Category\n2. Product\n3. View Customers")
	var i int
	_, serr := fmt.Scan(&i)
	if serr != nil {
		fmt.Println("Please Enter valid Input...")
		AdminWork()
	}
	switch i {
	case 1:
		categoryWork()

	case 2:
		productWork()

	case 3:
		customerWork(db)
	default:
		fmt.Println("Wrong Input")
		AdminWork()

	}
	cerr := db.Close()
	if cerr != nil {
		log.Fatal(cerr)
	}
}
