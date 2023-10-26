CREATE TABLE IF NOT EXISTS contracts  (
    id serial PRIMARY KEY,
    public_procurement_id INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL,
    serial_number TEXT NOT NULL,
    date_of_signing DATE NOT NULL,
    date_of_expiry DATE,
    net_value integer,
    gross_value integer,
    vat_vale integer,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    file_id INTEGER,
    FOREIGN KEY (public_procurement_id) REFERENCES items (id) ON UPDATE CASCADE ON DELETE CASCADE
);
