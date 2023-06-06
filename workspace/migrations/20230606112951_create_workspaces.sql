-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspaces (
                            id UUID PRIMARY KEY NOT NULL UNIQUE,
                            name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workspaces;
-- +goose StatementEnd
