// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID        sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
