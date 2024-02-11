-- +goose Up
-- +goose StatementBegin
-- products {
--   id uuid pk gen_random_uuid()
--   customer_id uuid fk
--   category_id uuid fk
--   name string
--   stok int 
--   created_at timestamp
--   updated_at timestamp
-- }
CREATE TABLE
    IF NOT EXISTS products (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        customer_id UUID,
        category_id UUID,
        name VARCHAR(255) NOT NULL,
        stok INT NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_products_customers FOREIGN KEY (customer_id) REFERENCES customers (id),
        CONSTRAINT fk_products_product_categories FOREIGN KEY (category_id) REFERENCES product_categories (id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;

-- +goose StatementEnd