CREATE TABLE product (
	ID SERIAL PRIMARY KEY,
	product_name varchar(50) NOT NULL,
	price numeric(10, 2) NOT NULL
);

SELECT * FROM product p

INSERT INTO product(product_name, price) VALUES('Sushi', 100);

