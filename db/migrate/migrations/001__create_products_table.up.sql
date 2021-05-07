DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id serial,
    name text NOT NULL,
    description text,
    price numeric(10, 2) NOT NULL DEFAULT 0.00,
    comment text,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

