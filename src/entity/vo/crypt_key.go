package vo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/dustinxie/ecc"
	"math/big"
)

type CryptKey struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewCryptKey() *CryptKey {
	// 秘密鍵からECDSAで公開鍵を生成
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &CryptKey{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}
}

func (c CryptKey) PrivateKeyValue() string {
	return fmt.Sprintf("%x", c.privateKey.D.Bytes())

}

func (c CryptKey) PublicKeyValue() string {
	return fmt.Sprintf("%x%x", c.publicKey.X.Bytes(), c.publicKey.Y.Bytes())
}

// String2bigIntTuple 文字列の公開鍵を取り出す
func String2bigIntTuple(s string) (big.Int, big.Int) {
	byteX, _ := hex.DecodeString(s[:64])
	byteY, _ := hex.DecodeString(s[64:])

	var bix big.Int
	var biy big.Int

	_ = bix.SetBytes(byteX)
	_ = biy.SetBytes(byteY)

	return bix, biy
}

// PublicKeyFromString 文字列から公開鍵を取得
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2bigIntTuple(s)
	return &ecdsa.PublicKey{Curve: ecc.P256k1(), X: &x, Y: &y}
}

// PrivateKeyFromString 文字列から秘密鍵を取得
func PrivateKeyFromString(privateKeyHex string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, _ := hex.DecodeString(privateKeyHex)
	var bi big.Int
	_ = bi.SetBytes(b)

	return &ecdsa.PrivateKey{PublicKey: *publicKey, D: &bi}
}
