package product

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rs/xid"
)

//CreateProduct func
func CreateProduct() {
	var product Product
	product.readProductData()
	product.Create()
}

func (product *Product) readProductData() {

	var price float32

	var quantity int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Product Name : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading product data ")
		product.readProductData()
	}
	product.Name = strings.TrimRight(name, "\n")

	fmt.Println("Enter Brand Name : ")
	brand, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading product data ")
		product.readProductData()
	}
	product.Brand = strings.TrimRight(brand, "\n")

	fmt.Println("Enter Category Name : ")
	category, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading product data ")
		product.readProductData()
	}
	product.Category = strings.TrimRight(category, "\n")

	fmt.Println("Enter Product Descrption : ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading product data ")
		product.readProductData()
	}
	product.Description = strings.TrimRight(desc, "\n")

	fmt.Println("Enter Product Price : ")
	_, perr := fmt.Scanf("%f", &price)
	if perr != nil {
		fmt.Println("Error while reading product data ")
		product.readProductData()
	}
	product.Price = price

	fmt.Println("Enter Product Quantity : ")
	fmt.Scan(&quantity)
	product.Quantity = quantity

	pid := xid.New()
	product.ID = pid.String()

}
