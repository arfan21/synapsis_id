-- +goose Up
-- +goose StatementBegin
CREATE TYPE transaction_status AS ENUM (
    'WAITING_PAYMENT',
    'PROCESSING',
    'COMPLETED',
    'FAILED'
);

CREATE TABLE
    IF NOT EXISTS transactions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        customer_id UUID,
        status transaction_status NOT NULL DEFAULT 'WAITING_PAYMENT',
        payment_method_id UUID NOT NULL,
        total_amount DECIMAL NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_transactions_customers FOREIGN KEY (customer_id) REFERENCES customers (id),
        CONSTRAINT fk_payments_payment_methods FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS transaction_status;

DROP TABLE IF EXISTS transactions;

-- +goose StatementEnd