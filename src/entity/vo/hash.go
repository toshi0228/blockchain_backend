package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Hash [32]byte
type HexHash string

// NewHashToHex 文字列をハッシュして16進数  byte(string) => byte32
func NewHashToHex(byte []byte) Hash {
	return sha256.Sum256(byte)
}

// Value 32byteのハッシュ値から16進数をhash値を取得
func (h Hash) Value() string {
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
