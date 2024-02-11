-- +goose Up
-- +goose StatementBegin
-- product_categories {
--   id uuid pk gen_random_uuid()
--   name string
--   created_at timestamp
--   updated_at timestamp
-- }
CREATE TABLE
    IF NOT EXISTS product_categories (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now ()
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_categories;

-- +goose StatementEnd