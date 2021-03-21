UPDATE
    products.products
SET
    name = $1,
    description = $2,
    price = $3,
    comment = $4
WHERE
    id = $5
