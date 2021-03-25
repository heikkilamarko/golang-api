INSERT INTO products.products (name, description, price, comment, created_at)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id
