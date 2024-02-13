-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS payments (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        transaction_id UUID NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_payments_transactions FOREIGN KEY (transaction_id) REFERENCES transactions (id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;

-- +goose StatementEnd