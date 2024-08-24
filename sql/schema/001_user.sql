-- +goose Up

create Table users(
    id UUID primary key,
    created_at Timestamp Not Null,
    updated_at Timestamp Not Null,
    name Text not Null
);

-- +goose Down
Drop Table users;