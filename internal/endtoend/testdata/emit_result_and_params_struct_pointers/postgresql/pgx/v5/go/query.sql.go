// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAll = `-- name: GetAll :many
SELECT a, b FROM foo
`

func (q *Queries) GetAll(ctx context.Context) ([]*Foo, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.A, &i.B); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllAByB = `-- name: GetAllAByB :many
SELECT a FROM foo WHERE b = $1
`

func (q *Queries) GetAllAByB(ctx context.Context, b pgtype.Int4) ([]pgtype.Int4, error) {
	rows, err := q.db.Query(ctx, getAllAByB, b)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Int4
	for rows.Next() {
		var a pgtype.Int4
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOne = `-- name: GetOne :one
SELECT a, b FROM foo WHERE a = $1 AND b = $2 LIMIT 1
`

type GetOneParams struct {
	A pgtype.Int4
	B pgtype.Int4
}

func (q *Queries) GetOne(ctx context.Context, arg *GetOneParams) (*Foo, error) {
	row := q.db.QueryRow(ctx, getOne, arg.A, arg.B)
	var i Foo
	err := row.Scan(&i.A, &i.B)
	return &i, err
}
