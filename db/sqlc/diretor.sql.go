// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: post.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createDiretor = `-- name: CreateDiretor :one
INSERT INTO diretores (
    nome,
    sexo
) VALUES (
    $1, $2
)
RETURNING id, nome, sexo, data_criacao, data_atualizacao, data_remocao
`

type CreateDiretorParams struct {
	Nome string `json:"nome"`
	Sexo Sexgen `json:"sexo"`
}

func (q *Queries) CreateDiretor(ctx context.Context, arg CreateDiretorParams) (Diretor, error) {
	row := q.queryRow(ctx, q.createDiretorStmt, createDiretor, arg.Nome, arg.Sexo)
	var i Diretor
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sexo,
		&i.DataCriacao,
		&i.DataAtualizacao,
		&i.DataRemocao,
	)
	return i, err
}

const deleteDiretor = `-- name: DeleteDiretor :exec
DELETE FROM diretores
WHERE id = $1
`

func (q *Queries) DeleteDiretor(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteDiretorStmt, deleteDiretor, id)
	return err
}

const getDiretorById = `-- name: GetDiretorById :one
SELECT id, nome, sexo, data_criacao, data_atualizacao, data_remocao FROM diretores
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDiretorById(ctx context.Context, id uuid.UUID) (Diretor, error) {
	row := q.queryRow(ctx, q.getDiretorByIdStmt, getDiretorById, id)
	var i Diretor
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sexo,
		&i.DataCriacao,
		&i.DataAtualizacao,
		&i.DataRemocao,
	)
	return i, err
}

const listDiretores = `-- name: ListDiretores :many
SELECT id, nome, sexo, data_criacao, data_atualizacao, data_remocao FROM diretores
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListDiretoresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDiretores(ctx context.Context, arg ListDiretoresParams) ([]Diretor, error) {
	rows, err := q.query(ctx, q.listDiretoresStmt, listDiretores, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Diretor
	for rows.Next() {
		var i Diretor
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Sexo,
			&i.DataCriacao,
			&i.DataAtualizacao,
			&i.DataRemocao,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDiretor = `-- name: UpdateDiretor :one
UPDATE diretores
set
nome = coalesce($1, nome),
sexo = coalesce($2, sexo)
WHERE id = $3
RETURNING id, nome, sexo, data_criacao, data_atualizacao, data_remocao
`

type UpdateDiretorParams struct {
	Nome sql.NullString `json:"nome"`
	Sexo NullSexgen     `json:"sexo"`
	ID   uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateDiretor(ctx context.Context, arg UpdateDiretorParams) (Diretor, error) {
	row := q.queryRow(ctx, q.updateDiretorStmt, updateDiretor, arg.Nome, arg.Sexo, arg.ID)
	var i Diretor
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sexo,
		&i.DataCriacao,
		&i.DataAtualizacao,
		&i.DataRemocao,
	)
	return i, err
}
