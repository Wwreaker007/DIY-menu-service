-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id    SERIAL  PRIMARY KEY,
    user_id     TEXT    NOT NULL,
    order_data       JSONB    NOT NULL,
    created_on  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_on  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_on  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
