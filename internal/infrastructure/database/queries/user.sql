-- name: CreateUser :exec
INSERT INTO "users" ("email", "password")
VALUES ($1, $2);

-- name: GetUserByID :one
SELECT id, email, verified, registered_at FROM users
WHERE id =  $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, verified, registered_at, password FROM users
WHERE email =  $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET
  email = COALESCE(sqlc.narg(email), email),
  password = COALESCE(sqlc.narg(password), password),
  verified = COALESCE(sqlc.narg(verified), verified)
WHERE id = sql.arg(id);

