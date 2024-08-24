-- +goose Up

create Table posts(
    id UUID primary key,
    created_at Timestamp Not Null,
    updated_at Timestamp Not Null,
    title Text not Null,
    description text,
    published_at Timestamp not null,
    url text not null unique,
    feed_id UUID not null references feeds(id) on delete cascade 
);

-- +goose Down
Drop Table posts;
