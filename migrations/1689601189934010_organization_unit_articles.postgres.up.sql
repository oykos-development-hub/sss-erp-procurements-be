CREATE TABLE IF NOT EXISTS organization_unit_articles (
    id serial PRIMARY KEY,
    article_id INTEGER NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    status text,
    is_rejected boolean,
        used_articles INTEGER DEFAULT 0,
    rejected_description text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles (id) ON UPDATE CASCADE ON DELETE CASCADE
);
