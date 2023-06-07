-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspace_users (
                                 id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                                 workspace_id UUID NOT NULL,
                                 user_id UUID NOT NULL,
                                 status VARCHAR(255) NOT NULL DEFAULT 'Pending',
                                 created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                 updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                 FOREIGN KEY (workspace_id) REFERENCES workspaces (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workspace_users;
-- +goose StatementEnd
