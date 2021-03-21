SELECT
    id,
    name,
    description,
    price,
    comment
FROM
    products.products
ORDER BY
    id
LIMIT $1 OFFSET $2
