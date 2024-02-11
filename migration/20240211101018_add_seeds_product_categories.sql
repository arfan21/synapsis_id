-- +goose Up
-- +goose StatementBegin
INSERT INTO
    product_categories (name)
VALUES
    ('Electronics'),
    ('Fashion'),
    ('Home'),
    ('Toys'),
    ('Books');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM product_categories
WHERE
    name IN ('Electronics', 'Fashion', 'Home', 'Toys', 'Books');

-- +goose StatementEnd