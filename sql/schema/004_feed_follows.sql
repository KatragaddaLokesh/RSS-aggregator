-- +goose Up

create Table feeds_follows(
    id UUID primary key,
    created_at Timestamp Not Null,
    updated_at Timestamp Not Null,
    user_id UUID not Null references users(id) on delete cascade,
    feed_id UUID not null references feeds(id) on delete cascade, 
    unique(user_id,feed_id)
);

-- +goose Down
Drop Table feeds_follows;

