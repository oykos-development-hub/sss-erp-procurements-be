CREATE TABLE IF NOT EXISTS items (
    id serial PRIMARY KEY,
    budget_indent_id INTEGER NOT NULL,
    plan_id INTEGER NOT NULL,
    is_open_procurement BOOLEAN NOT NULL,
    title TEXT NOT NULL,
    article_type TEXT NOT NULL, 
    status TEXT,
    serial_number TEXT,
    date_of_publishing DATE,
    date_of_awarding DATE,
    file_id INTEGER,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    FOREIGN KEY (plan_id) REFERENCES plans (id) ON UPDATE CASCADE ON DELETE CASCADE
);
