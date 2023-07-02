-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE members (
                            id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                            workspace_id UUID NOT NULL,
                            user_id UUID NOT NULL,
                            role VARCHAR(255) NOT NULL,
                            invitation_id VARCHAR(255) NOT NULL,
                            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS members;
-- +goose StatementEnd
