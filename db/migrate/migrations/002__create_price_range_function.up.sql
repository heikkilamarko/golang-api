CREATE OR REPLACE FUNCTION price_range_func(
    OUT min_price NUMERIC,
    OUT max_price NUMERIC)
    RETURNS record
    LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    select min(price), max(price)
    into min_price, max_price
    from products;
END;
$BODY$;

COMMENT ON FUNCTION price_range_func()
    IS 'Gets product price range as min and max values.';
