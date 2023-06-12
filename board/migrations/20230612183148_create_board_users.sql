-- +goose Up
-- +goose StatementBegin
CREATE TABLE board_users (
                                 id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                                 board_id UUID NOT NULL,
                                 user_id UUID NOT NULL,
                                 status VARCHAR(255) NOT NULL DEFAULT 'Pending',
                                 created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                 updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                 FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS board_users;
-- +goose StatementEnd
