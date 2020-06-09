package admin

import (
	"EShopeeREPO/shop/category"
	"EShopeeREPO/shop/product"
	"fmt"
	"strconv"
)

func categoryWork() {
	fmt.Println("Select an Option :\n1. View Categories\n2. Create Category\n3. Remove Category\n00. Back")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		categoryWork()
	}

	switch i {
	case 1:
		viewCategoryList()
	case 2:
		addNewCategory()
	case 3:
		removeCategory()
	case 00:
		AdminWork()
	default:
		AdminWork()
	}
}

func viewCategoryList() {
	list := make([]category.List, 0)
	var err error
	list, err = category.GetCategoryList()
	if err != nil {
		fmt.Println(err)
		categoryWork()
	}
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.CategoryName)
	}
	fmt.Println("Enter Category number to view products from the Category or press 00 to back:")
	var i int
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Println("Please Enter valid Input...")
		category.GetCategoryList()
	}
	if i == 00 {
		categoryWork()
	}

	prodList := make([]product.ProductList, 0)
	var perr error
	prodList, perr = product.GetProductsByCategory(list[i-1].CategoryName)
	if perr != nil {
		fmt.Println(perr)
	}
	for index, item := range prodList {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.Name)
	}
	categoryWork()
}

func addNewCategory() {
	category.CreateCategory()
	categoryWork()
}

func removeCategory() {
	list := make([]category.List, 0)
	var err error
	list, err = category.GetCategoryList()
	if err != nil {
		fmt.Println(err)
		categoryWork()
	}
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.CategoryName)
	}

	fmt.Println("Enter category number to remove category : ")
	var i int
	fmt.Scan(&i)
	category.RemoveCategory(list[i-1].CategoryName)
	categoryWork()
}
