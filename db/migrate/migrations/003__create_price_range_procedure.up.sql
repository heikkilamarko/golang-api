CREATE OR REPLACE PROCEDURE price_range_proc(
    INOUT min_price NUMERIC,
    INOUT max_price NUMERIC)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    select min(price), max(price)
    into min_price, max_price
    from products;
END;
$BODY$;

COMMENT ON PROCEDURE price_range_proc(NUMERIC, NUMERIC)
    IS 'Gets product price range as min and max values.';
