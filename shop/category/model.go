package category

import "time"

//Category is the Category of the products
type Category struct {
	CategoryID   int `json:"categoryId" bson:"categoryId"`
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
