-- name: GetUserByEmail :one
SELECT * FROM Users WHERE email = $1;


-- name: CreateUser :one
INSERT INTO Users (username, email, password)
VALUES ($1, $2, $3)
RETURNING id, username, email, password, created_at;