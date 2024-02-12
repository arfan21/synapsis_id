-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS carts (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        customer_id UUID,
        product_id UUID,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_carts_customers FOREIGN KEY (customer_id) REFERENCES customers (id),
        CONSTRAINT fk_carts_products FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
        CONSTRAINT unique_cart UNIQUE (customer_id, product_id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;

-- +goose StatementEnd