CREATE TABLE IF NOT EXISTS contract_article_overages (
    id serial PRIMARY KEY,
    article_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles (id) ON UPDATE CASCADE ON DELETE CASCADE
);
