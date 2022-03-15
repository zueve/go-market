CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    login VARCHAR(64) NOT NULL,
    password_hash VARCHAR(64) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_customer_login ON customer(login);

CREATE TABLE IF NOT EXISTS billing (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    amount INT NOT NULL CHECK (amount >= 0),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);

CREATE TABLE IF NOT EXISTS billing_order (
    id SERIAL PRIMARY KEY,
    billing_id INT NOT NULL,
    invoice VARCHAR(16) NOT NULL UNIQUE,
    direction VARCHAR(16) NOT NULL,
    amount INT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (billing_id) REFERENCES billing (id)
);


CREATE TABLE IF NOT EXISTS accrual (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    invoice BIGINT NOT NULL UNIQUE,
    amount INT NOT NULL,
    "status" VARCHAR(16) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);
