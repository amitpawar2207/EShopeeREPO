package product

//Product for details of any Product
type Product struct {
	ID          string  `bson:"id"`
	Name        string  `bson:"name"`
	Brand       string  `bson:"brand"`
	Category    string  `bson:"category"`
	Description string  `bson:"description"`
	Price       float32 `bson:"price"`
	Quantity    int     `bson:"quantity"`
}
