CREATE OR REPLACE FUNCTION products.price_range(
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

ALTER FUNCTION products.price_range()
    OWNER TO postgres;

COMMENT ON FUNCTION products.price_range()
    IS 'Gets product price range as min and max values.';
