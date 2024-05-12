CREATE TABLE IF NOT EXISTS users_wallets (
    id BIGINT UNIQUE,
    balance DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS transactions_history (
    transaction_id SERIAL PRIMARY KEY,
    user_id BIGINT,
    amount DOUBLE PRECISION,
    transaction_type VARCHAR(255),
    description VARCHAR(255),
    transaction_date TIMESTAMP
);