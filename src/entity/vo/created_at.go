package vo

import (
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (ca CreatedAt) Value() time.Time {
	return time.Time(ca)
}

func (ca CreatedAt) Equals(ca2 CreatedAt) bool {
	return ca == ca2
}
