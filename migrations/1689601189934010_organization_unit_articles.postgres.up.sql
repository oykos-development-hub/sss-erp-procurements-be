CREATE TABLE IF NOT EXISTS organization_unit_articles (
    id serial PRIMARY KEY,
    article_id INTEGER NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    status text,
    is_rejected boolean,
    rejected_description text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
