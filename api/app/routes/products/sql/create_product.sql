INSERT INTO products.products (name, description, price, comment)
    VALUES ($1, $2, $3, $4)
RETURNING
    id
