package vo

import "time"

type UnixTimeNow int64

func NewUnixTimeNow() UnixTimeNow {
	return UnixTimeNow(time.Now().UnixNano())
}

func (utn UnixTimeNow) Value() int64 {
	return int64(utn)
}

func (utn UnixTimeNow) Equals(utn2 UnixTimeNow) bool {
	return utn == utn2
}
