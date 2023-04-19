package model

import "time"

type Borrow struct {
	ID       uint
	BookID   uint
	UserID   uint
	Borrowed time.Time
	Returned *time.Time
}
