CREATE TABLE dashboard (
    product_id VARCHAR(24) PRIMARY KEY, 
    units_sold SERIAL,
    color VARCHAR(24) NOT NULL,
    cost_price BIGINT DEFAULT 0 NOT NULL,
    selling_price  BIGINT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

