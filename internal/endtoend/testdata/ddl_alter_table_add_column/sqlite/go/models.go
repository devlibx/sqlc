// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"database/sql"
)

type Venue struct {
	Name     sql.NullString
	Location sql.NullString
	Size     sql.NullInt64
}
