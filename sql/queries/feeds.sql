-- name: CreateFeed :one
Insert into feeds (id, created_at,updated_at, name, url,user_id)
values ($1,$2,$3,$4,$5,$6)
returning *;

-- name: GetFeed :many
select * from feeds;

-- name: GetNextFeedToFetch :many
select * from feeds
order by last_fetched_at Asc Nulls first 
Limit $1;


-- name: MarkFeedAsFetched :one
update feeds set last_fetched_at = now(),
updated_at = now()
where id = $1
returning *;
