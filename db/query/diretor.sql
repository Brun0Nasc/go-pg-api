-- name: CreateDiretor :one
INSERT INTO diretores (
    nome,
    sexo
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetDiretorById :one
SELECT * FROM diretores
WHERE id = $1 LIMIT 1;

-- name: ListDiretores :many
SELECT * FROM diretores
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateDiretor :one
UPDATE diretores
set
nome = coalesce(sqlc.narg('nome'), nome),
sexo = coalesce(sqlc.narg('sexo'), sexo)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteDiretor :exec
DELETE FROM diretores
WHERE id = $1;