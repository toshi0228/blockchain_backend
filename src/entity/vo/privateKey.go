package vo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
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
