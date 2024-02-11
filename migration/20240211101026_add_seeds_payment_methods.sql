-- +goose Up
-- +goose StatementBegin
INSERT INTO
    payment_methods (name)
VALUES
    ('OVO');

INSERT INTO
    payment_methods (name)
VALUES
    ('Gopay');

INSERT INTO
    payment_methods (name)
VALUES
    ('Dana');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM payment_methods
WHERE
    name = 'OVO';

DELETE FROM payment_methods
WHERE
    name = 'Gopay';

DELETE FROM payment_methods
WHERE
    name = 'Dana';

-- +goose StatementEnd