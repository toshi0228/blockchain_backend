package vo

import "encoding/hex"

type Hash [32]byte
type HexHash string

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

//
//func (s StringData) Equals(s2 StringData) bool {
//	return s == s2
//}
//
//// Hash ハッシュ化
//func (s StringData) Hash() [32]byte {
//	return sha256.Sum256([]byte(s))
//}
//
//func (s StringData) HashToHex() string {
//	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
//}
