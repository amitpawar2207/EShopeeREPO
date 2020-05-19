package category

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//CreateCategory func
func CreateCategory() {

	var cat Category
	cat.CategoryName = readCategoryName()
	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()
	cat.Create()
}

func readCategoryName() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter category Name : ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(text, "\n")
}
