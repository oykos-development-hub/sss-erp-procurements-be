ALTER TABLE articles ALTER COLUMN net_price TYPE decimal(10, 2) USING net_price::numeric(10,2);
ALTER TABLE contracts ALTER COLUMN net_value TYPE decimal(10, 2) USING net_value::numeric(10,2);
ALTER TABLE contracts ALTER COLUMN gross_value TYPE decimal(10, 2) USING gross_value::numeric(10,2);
ALTER TABLE contract_articles ALTER COLUMN net_value TYPE decimal(10, 2) USING net_value::numeric(10,2);
ALTER TABLE contract_articles ALTER COLUMN gross_value TYPE decimal(10, 2) USING gross_value::numeric(10,2);

