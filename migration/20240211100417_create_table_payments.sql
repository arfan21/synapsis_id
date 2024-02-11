-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS payments (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        transaction_id UUID NOT NULL,
        payment_method_id UUID NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_payments_transactions FOREIGN KEY (transaction_id) REFERENCES transactions (id),
        CONSTRAINT fk_payments_payment_methods FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;

-- +goose StatementEnd