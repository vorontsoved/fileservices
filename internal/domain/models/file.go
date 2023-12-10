package models

import (
	"time"
)

type File struct {
	ID         int
	DateCreate time.Time
	DateModify time.Time
}
