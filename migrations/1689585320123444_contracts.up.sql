CREATE TABLE IF NOT EXISTS contracts  (
    id serial PRIMARY KEY,
    public_procurement_id INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL,
    serial_number TEXT NOT NULL,
    date_of_signing DATE NOT NULL,
    date_of_expiry DATE,
    net_value TEXT NOT NULL,
    gross_value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    file_id INTEGER,
    FOREIGN KEY (public_procurement_id) REFERENCES items (id) ON UPDATE CASCADE ON DELETE CASCADE
);
