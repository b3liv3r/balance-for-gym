CREATE TABLE IF NOT EXISTS users_wallets (
    id INT UNIQUE,
    balance DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS transactions_history (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT,
    amount DOUBLE PRECISION,
    transaction_type VARCHAR(255),
    description VARCHAR(255),
    transaction_date TIMESTAMP
);