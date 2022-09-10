package vo

import (
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (ua UpdatedAt) Value() time.Time {
	return time.Time(ua)
}

func (ua UpdatedAt) Equals(ua2 UpdatedAt) bool {
	return ua == ua2
}
