-- +goose Up
-- +goose StatementBegin
INSERT INTO
    payment_methods (name)
VALUES
    ('OVO'),
    ('Gopay'),
    ('Dana');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM payment_methods
WHERE
    name IN ('OVO', 'Gopay', 'Dana');

-- +goose StatementEnd