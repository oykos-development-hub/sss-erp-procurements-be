CREATE TABLE IF NOT EXISTS contract_articles  (
    id serial PRIMARY KEY,
    article_id INTEGER NOT NULL,
    contract_id INTEGER NOT NULL,
    net_value INTEGER,
    gross_value INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (contract_id) REFERENCES contracts (id) ON UPDATE CASCADE ON DELETE CASCADE
);
