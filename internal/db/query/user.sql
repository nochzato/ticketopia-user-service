-- name: CreateUser :one
INSERT INTO users (
	username,
	full_name,
	password,
	email
) VALUES (
	$1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
	password = COALESCE(sqlc.narg(password), password),
	full_name = COALESCE(sqlc.narg(full_name), full_name),
	email = COALESCE(sqlc.narg(email), email)
WHERE
    id = sqlc.arg(id)
RETURNING *;
