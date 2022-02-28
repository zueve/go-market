CREATE TABLE IF NOT EXISTS customer (
    id SERIAL,
    login VARCHAR(64) NOT NULL,
    password_hash VARCHAR(64) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_customer_login ON customer(login);
