-- +goose Up
-- +goose StatementBegin
-- payment_methods {
--   id uuid pk gen_random_uuid()
--   name string
--   created_at timestamp
--   updated_at timestamp
-- }
CREATE TABLE
    IF NOT EXISTS payment_methods (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now ()
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_methods;

-- +goose StatementEnd