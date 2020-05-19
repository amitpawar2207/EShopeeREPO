package admin

import (
	"EShopeeREPO/shop/category"
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
	list = category.GetCategoryList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.CategoryName)
	}
	categoryWork()
}

func addNewCategory() {
	category.CreateCategory()
	categoryWork()
}

func removeCategory() {
	list := make([]category.List, 0)
	list = category.GetCategoryList()
	for index, item := range list {
		fmt.Println(strconv.Itoa(index+1) + ". " + item.CategoryName)
	}

	fmt.Println("Enter category number to remove category : ")
	var i int
	fmt.Scan(&i)
	category.RemoveCategory(list[i-1].CategoryName)
	categoryWork()
}
