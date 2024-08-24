-- name: CreateUser :one
Insert into users (id, created_at,updated_at, name, api_keys)
values ($1,$2,$3,$4,
    encode(sha256(random()::text::bytea),'hex')
)
returning *;
-- name: GetUserByAPIKey :one
select * from users where api_keys = $1;
