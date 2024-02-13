-- +goose Up
-- +goose StatementBegin
ALTER TABLE transaction_details
ADD COLUMN qty INT NOT NULL DEFAULT 0;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE transaction_details
DROP COLUMN qty;

-- +goose StatementEnd