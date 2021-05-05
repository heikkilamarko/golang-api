UPDATE
    products
SET
    name = $1,
    description = $2,
    price = $3,
    comment = $4,
    updated_at = $5
WHERE
    id = $6
