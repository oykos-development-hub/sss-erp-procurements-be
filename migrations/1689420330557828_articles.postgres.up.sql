CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    budget_indent_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    net_price integer NOT NULL,
    vat_percentage TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (item_id) REFERENCES items (id) ON UPDATE CASCADE ON DELETE CASCADE
);
