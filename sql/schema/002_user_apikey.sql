-- +goose Up
Alter Table users add column api_keys varchar(64) unique not Null default(
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
Alter Table users drop column api_keys;
