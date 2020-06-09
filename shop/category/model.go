package category

import "time"

//Category is the Category of the products
type Category struct {
	CategoryName string    `bson:"categoryname"`
	CreatedAt    time.Time `bson:"createdat"`
	UpdatedAt    time.Time `bson:"updatedat"`
}
