-- +goose Up

create Table feeds(
    id UUID primary key,
    created_at Timestamp Not Null,
    updated_at Timestamp Not Null,
    name Text not Null,
    url Text unique not Null,
    user_id UUID Not Null references users(id) on delete cascade
    
);

-- +goose Down
Drop Table feeds;
