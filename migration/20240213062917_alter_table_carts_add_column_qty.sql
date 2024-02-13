-- +goose Up
-- +goose StatementBegin
ALTER TABLE carts
ADD COLUMN qty INT NOT NULL DEFAULT 1;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE carts
DROP COLUMN qty;

-- +goose StatementEnd