// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const listFoo = `-- name: ListFoo :many
SELECT bar, bam, baz FROM foo
`

func (q *Queries) ListFoo(ctx context.Context) ([]Foo, error) {
	rows, err := q.db.QueryContext(ctx, listFoo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.Bar, &i.Bam, &i.Baz); err != nil {
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
