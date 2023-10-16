CREATE TABLE IF NOT EXISTS plans (
    id serial PRIMARY KEY,
    year TEXT NOT NULL,
    title TEXT NOT NULL,
    active BOOLEAN NOT NULL,
    serial_number TEXT,
    date_of_publishing DATE,
    date_of_closing DATE,
    pre_budget_id INTEGER,
    is_pre_budget BOOLEAN NOT NULL,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (pre_budget_id) REFERENCES plans (id) ON UPDATE CASCADE ON DELETE CASCADE
);
