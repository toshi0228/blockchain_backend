package vo

import "github.com/google/uuid"

type ID uint32

func NewID() ID {
	return ID(uuid.New().ID())
}

func (i ID) Value() uint32 {
	return uint32(i)
}

func (i ID) Equals(i2 ID) bool {
	return i == i2
}
