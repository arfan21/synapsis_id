-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS transactions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        customer_id UUID,
        status VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_transactions_customers FOREIGN KEY (customer_id) REFERENCES customers (id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;

-- +goose StatementEnd