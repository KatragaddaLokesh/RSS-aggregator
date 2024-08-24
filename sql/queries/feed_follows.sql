-- name: CreateFeedFollow :one
Insert into feeds_follows (id, created_at, updated_at, user_id, feed_id)
values ($1,$2,$3,$4,$5)
returning *;

-- name: GetFeedFollows :many
select * from feeds_follows where user_id=$1;


-- name: DeleteFeedFollows :exec
delete from feeds_follows where id = $1 and user_id=$2;