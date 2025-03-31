create table product (
	id SERIAL primary key,
	product_name varchar(50) not null,
	price numeric(10, 2) not null
);

select * from product p

insert into product(product_name, price) values('Sushi', 100);

