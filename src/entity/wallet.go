package entity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	BlockChainAddress string
}

func NewWallet() *Wallet {
	w := &Wallet{}

	// 秘密鍵からECDSAで公開鍵を生成
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	//公開鍵をハッシュ関数SHA-256に通しハッシュ値を得る
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	//そのハッシュ値をさらにハッシュ関数RIPEMD-160に通しハッシュ値を得る
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	//ハッシュ値の先頭にプレフィックスとして00を加える
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	//ハッシュ関数SHA-256に通す
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	//もう一度ハッシュ関数SHA-256に通す
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	//4バイトのチェックサムを一番後ろに加える
	chsum := digest6[:4]
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])

	//Base58のフォーマットでエンコーディングする
	address := base58.Encode(dc8)
	w.BlockChainAddress = address

	return w
}

func (w *Wallet) BlockchainAddress() string {
	return w.BlockChainAddress
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
