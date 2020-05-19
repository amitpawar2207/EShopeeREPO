package admin

import (
	"EShopeeREPO/shop/product"
	"fmt"
	"strconv"
)

func productWork() {
	fmt.Println("Select an Option :\n1. View Products\n2. Create Product\n3. Remove Product\n4. Update Product Information\n00. Back")
	var i int
	fmt.Scan(&i)
	switch i {
	case 1:
		viewProductList()
	case 2:
		addNewProduct()
	case 3:
		removeProduct()
	case 4:
		updateProduct()
	case 00:
		AdminWork()
	default:
		AdminWork()
	}
}

func addNewProduct() {
	product.CreateProduct()
	productWork()
}

func removeProduct() {
	list := make([]product.ProductList, 0)
	list = product.GetProductList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}

	fmt.Println("Enter Product number to remove product")
	var i int
	fmt.Scan(&i)
	product.RemoveProduct(list[i-1].Name)
	fmt.Println(list[i-1].Name, " removed.")
	productWork()
}

func viewProductList() {
	list := make([]product.ProductList, 0)
	list = product.GetProductList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index) + ". " + item.Name)
	}
	productWork()

}

func updateProduct() {
	list := make([]product.ProductList, 0)
	list = product.GetProductList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}

	fmt.Println("Enter Product number to update product info : ")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		updateProduct()
	}
	prod := product.GetProductDetails(list[i-1].Name, "")
	product.UpdateProductData(prod, prod.Name)
	productWork()
}
