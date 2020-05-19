package customer

import "time"

//Cart table
type Cart struct {
	CartID          int       `json:"cart_id"`
	CustomerID      int       `json:"customer_id"`
	ProductID       string    `json:"product_id"`
	ProductQuantity int       `json:"product_quantity"`
	Amount          float32   `json:"amount"`
	Checkout        bool      `json:"checkout"`
	UpdatedAt       time.Time `json:"updatedat"`
}

//Order table
type Order struct {
	OrderID      int       `json:"order_id"`
	CustomerID   int       `json:"customer_id"`
	FinalAmount  float32   `json:"final_amount"`
	Discount     float32   `json:"discount"`
	CheckoutDate time.Time `json:"checkout_date"`
}

//LoginInfo to login
type LoginInfo struct {
	UserName string
	Password string
}

//User table
type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EmailAddress  string `json:"email_address"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
}
