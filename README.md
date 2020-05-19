# EShopeeREPO
Following things are required to run this application - 
Golang, MySQL, MongoDB


Edit the mysql configurations from mysql_constants.go file. 
File path - common/factory/mysql_constants.go

Edit the mysql configurations from mongo_constants.go file. 
File path - common/factory/mongo_constants.go



Queries for MySQL to before running application -

create table users(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50),
	role VARCHAR(10),
	contact_number VARCHAR(10),
	email_address VARCHAR(20),
	address VARCHAR(100),
	username VARCHAR(100),
	password VARCHAR(100)
);

create table cart(
	cart_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	customer_id INT,
	product_id varchar(50),
	product_quantity INT,
	amount FLOAT,
	checkout bool,
	updatedat timestamp CURRENT_TIMESTAMP
	
);

create table orders(
	order_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	customer_id INT,
	final_amount FLOAT,
	discount FLOAT,
	checkout_date timestamp CURRENT_TIMESTAMP
);



To run this application use either ways- 
1. go run main.go
2. go buld shop 
   ./EShopeeREPO
