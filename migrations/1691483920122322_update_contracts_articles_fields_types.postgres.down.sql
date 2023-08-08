ALTER TABLE articles ALTER COLUMN net_price TYPE text USING net_price::text;
ALTER TABLE contracts ALTER COLUMN net_value TYPE text USING net_value::text;
ALTER TABLE contracts ALTER COLUMN gross_value TYPE text USING gross_value::text;
ALTER TABLE contract_articles ALTER COLUMN net_value TYPE text USING net_value::text;
ALTER TABLE contract_articles ALTER COLUMN gross_value TYPE text USING gross_value::text;
