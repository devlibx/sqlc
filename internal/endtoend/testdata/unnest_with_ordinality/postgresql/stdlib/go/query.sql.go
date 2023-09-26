// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const getValues = `-- name: GetValues :many
SELECT id, index::bigint, value::text
FROM array_values AS x, unnest(values) WITH ORDINALITY AS y (value, index)
`

type GetValuesRow struct {
	ID    int64
	Index int64
	Value string
}

func (q *Queries) GetValues(ctx context.Context) ([]GetValuesRow, error) {
	rows, err := q.db.QueryContext(ctx, getValues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetValuesRow
	for rows.Next() {
		var i GetValuesRow
		if err := rows.Scan(&i.ID, &i.Index, &i.Value); err != nil {
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
