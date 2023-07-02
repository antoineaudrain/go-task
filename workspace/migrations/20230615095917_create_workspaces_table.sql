-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE workspaces (
                            id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                            name VARCHAR(255) NOT NULL,
                            owner_id UUID NOT NULL,
                            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workspaces;
-- +goose StatementEnd
