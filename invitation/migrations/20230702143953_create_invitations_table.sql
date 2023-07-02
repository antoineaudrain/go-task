-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE invitations (
                             id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                             email VARCHAR(255) NOT NULL,
                             code VARCHAR(255) NOT NULL,
                             workspace_id UUID NOT NULL,
                             last_sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                             created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                             updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                             deleted_at TIMESTAMPTZ
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invitations;
-- +goose StatementEnd
