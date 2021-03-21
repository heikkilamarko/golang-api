SELECT
    id,
    name,
    description,
    price,
    comment
FROM
    products.products
WHERE
    id = $1
