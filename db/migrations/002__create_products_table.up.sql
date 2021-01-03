DROP TABLE IF EXISTS products.products;

CREATE TABLE products.products
(
    id SERIAL,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    comment TEXT,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);
