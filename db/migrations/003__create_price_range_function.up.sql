CREATE OR REPLACE FUNCTION products.price_range_func(
    OUT min_price numeric,
    OUT max_price numeric)
    RETURNS record
    LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    select min(price), max(price)
    into min_price, max_price
    from products.products;
END;
$BODY$;

COMMENT ON FUNCTION products.price_range_func()
    IS 'Gets product price range as min and max values.';
