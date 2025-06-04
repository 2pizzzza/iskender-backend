-- name: GetUserByUsernameAndPassword :one
SELECT * FROM Users WHERE email = $1 AND password = $2;