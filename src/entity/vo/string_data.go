package vo

import (
	"crypto/sha256"
	"fmt"
)

type StringData string

func NewStringData(stringData string) StringData {
	return StringData(stringData)
}

func (s StringData) Value() string {
	return string(s)
}

func (s StringData) Equals(s2 StringData) bool {
	return s == s2
}

// Hash ハッシュ化
func (s StringData) Hash() [32]byte {
	return sha256.Sum256([]byte(s))
}

func (s StringData) HashToHex() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
