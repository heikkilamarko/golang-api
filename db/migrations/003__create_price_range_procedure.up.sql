CREATE OR REPLACE PROCEDURE price_range_proc(
    INOUT min_price numeric,
    INOUT max_price numeric)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    select min(price), max(price)
    into min_price, max_price
    from products;
END;
$BODY$;

COMMENT ON PROCEDURE price_range_proc(numeric, numeric)
    IS 'Gets product price range as min and max values.';
