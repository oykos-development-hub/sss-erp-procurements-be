CREATE TABLE IF NOT EXISTS organization_unit_plan_limits (
    id serial PRIMARY KEY,
    organization_unit_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    limit_value INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (item_id) REFERENCES items (id) ON UPDATE CASCADE ON DELETE CASCADE
);
