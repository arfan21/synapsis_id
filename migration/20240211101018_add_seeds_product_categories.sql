-- +goose Up
-- +goose StatementBegin
INSERT INTO
    product_categories (name)
VALUES
    ("Electronics");

INSERT INTO
    product_categories (name)
VALUES
    ("Fashion");

INSERT INTO
    product_categories (name)
VALUES
    ("Home");

INSERT INTO
    product_categories (name)
VALUES
    ("Toys");

INSERT INTO
    product_categories (name)
VALUES
    ("Books");

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM product_categories
WHERE
    name = "Electronics";

DELETE FROM product_categories
WHERE
    name = "Fashion";

DELETE FROM product_categories
WHERE
    name = "Home";

DELETE FROM product_categories
WHERE
    name = "Toys";

DELETE FROM product_categories
WHERE
    name = "Books";

-- +goose StatementEnd