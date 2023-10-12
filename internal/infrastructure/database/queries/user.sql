-- name: CreateUser :one
INSERT INTO "users" ("id", "email", "password", "verified", "registered_at")
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, registered_at;

-- name: GetUserByID :one
SELECT id, email, verified, registered_at FROM users
WHERE id =  $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET
  email = COALESCE(sqlc.narg(email), email),
  password = COALESCE(sqlc.narg(password), password),
  verified = COALESCE(sqlc.narg(verified), verified)
WHERE id = sql.arg(id);

