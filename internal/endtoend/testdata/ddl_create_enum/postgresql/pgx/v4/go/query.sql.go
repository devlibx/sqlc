// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const listFoo = `-- name: ListFoo :many
SELECT val FROM foo
`

func (q *Queries) ListFoo(ctx context.Context) ([]Foobar, error) {
	rows, err := q.db.Query(ctx, listFoo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foobar
	for rows.Next() {
		var val Foobar
		if err := rows.Scan(&val); err != nil {
			return nil, err
		}
		items = append(items, val)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
