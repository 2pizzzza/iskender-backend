-- name: GetUserByUsername :one
SELECT * FROM Users WHERE username = $1;
