-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspace_users (
                                 workspace_id UUID NOT NULL,
                                 user_id UUID NOT NULL,
                                 PRIMARY KEY (workspace_id, user_id),
                                 FOREIGN KEY (workspace_id) REFERENCES workspaces (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workspace_users;
-- +goose StatementEnd
