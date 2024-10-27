---name: CreateSession :one
insert into sessions ("name", "refresh_token", "expires_at")
values ($1, $2, $3, $4)