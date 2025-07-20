CREATE TABLE users (
id SERIAL PRIMARY KEY,
username varchar(200) not null,
names VARCHAR(200) not null,
email VARCHAR(100) UNIQUE,
password TEXT,
role VARCHAR(50),
amount numeric(10,2) DEFAULT 0 NOT NULL,
code_referal varchar(50) not null,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE categories (
id SERIAL PRIMARY KEY,
name VARCHAR(100),
description TEXT
);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
name_product VARCHAR(100),
description TEXT,
price NUMERIC(15,2),
stock INT,
code_scan varchar(100) not null,
category_id INT REFERENCES categories(id),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cart (
id SERIAL PRIMARY KEY,
user_id INT UNIQUE REFERENCES users(id),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cart_items (
id SERIAL PRIMARY KEY,
cart_id INT REFERENCES cart(id),
product_id INT REFERENCES products(id),
quantity INT
);

CREATE TABLE orders (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id),
order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
status VARCHAR(50),
total_amount NUMERIC(10,2),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
id SERIAL PRIMARY KEY,
order_id INT REFERENCES orders(id),
product_id INT REFERENCES products(id),
quantity INT,
price NUMERIC(10,2)
);

CREATE TABLE payments (
id SERIAL PRIMARY KEY,
order_id INT UNIQUE REFERENCES orders(id),
payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
payment_method VARCHAR(50),
amount NUMERIC(10,2),
status VARCHAR(50)
);

