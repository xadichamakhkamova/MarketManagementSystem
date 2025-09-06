CREATE TABLE IF NOT EXISTS debts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    jshshir TEXT NOT NULL,
    address TEXT NOT NULL,
    bag_id TEXT NOT NULL,
    price FLOAT NOT NULL,
    price_paid FLOAT NOT NULL,
    acquaintance TEXT DEFAULT NULL,
    collateral TEXT DEFAULT NULL,
    deadline TIMESTAMP NOT NULL,
    status INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- 1 berilmadi
-- 2 berildi