SELECT
    id,
    name,
    description,
    price,
    comment,
    created_at,
    updated_at
FROM
    products
WHERE
    id = $1
