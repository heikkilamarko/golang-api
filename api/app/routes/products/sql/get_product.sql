SELECT
    id,
    name,
    description,
    price,
    comment,
    created_at
FROM
    products.products
WHERE
    id = $1
