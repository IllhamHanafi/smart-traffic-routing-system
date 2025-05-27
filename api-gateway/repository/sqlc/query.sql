-- name: GetUserByEmail :one
SELECT id, name, role, email, password FROM "user"
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (
    id, name, role, email, password, created_at, created_by, updated_at, updated_by
) VALUES (
    gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id;