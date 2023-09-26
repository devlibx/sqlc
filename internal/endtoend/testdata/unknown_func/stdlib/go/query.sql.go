// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const listFoos = `-- name: ListFoos :one
SELECT id FROM foo WHERE id = frobnicate($1)
`

func (q *Queries) ListFoos(ctx context.Context, frobnicate interface{}) (string, error) {
	row := q.db.QueryRowContext(ctx, listFoos, frobnicate)
	var id string
	err := row.Scan(&id)
	return id, err
}
