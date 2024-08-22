create table products (
	id SERIAL primary key,
	name varchar(50) not null,
	price numeric(10, 2) not null
);

select * from products;

drop table product;

insert into products (name, price) values ('Sushi', 100);
