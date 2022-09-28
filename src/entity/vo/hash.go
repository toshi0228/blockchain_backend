package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Hash [32]byte
type HexHash string

// NewHash 文字列をハッシュして16進数  byte(string) => byte32
func NewHash(byte []byte) Hash {
	return sha256.Sum256(byte)
}

// Value 32byteのハッシュ値を返す
func (h Hash) Value() [32]byte {
	return h
}

// ValueToHex ハッシュしたデータを16進数に変換
func (h Hash) ValueToHex() string {
	return fmt.Sprintf("%x", h)
}

// NewHexHashToByte32  hex => byte32
func NewHexHashToByte32(hexString string) HexHash {
	return HexHash(hexString)
}

// Value 16進数のハッシュ値から32byteのハッシュ値
func (hh HexHash) Value() [32]byte {
	hexHash, _ := hex.DecodeString(string(hh))

	//ハッシュの値は32byte配列なので定義
	byte32 := [32]byte{}

	for index := range hexHash {
		byte32[index] = hexHash[index]
	}
	return byte32
}
