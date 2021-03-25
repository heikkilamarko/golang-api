CREATE OR REPLACE PROCEDURE products.price_range_proc(
    INOUT min_price numeric,
    INOUT max_price numeric)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    select min(price), max(price)
    into min_price, max_price
    from products.products;
END;
$BODY$;

COMMENT ON PROCEDURE products.price_range_proc(numeric, numeric)
    IS 'Gets product price range as min and max values.';
