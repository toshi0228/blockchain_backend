package entity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"golang.org/x/crypto/ripemd160"
	"time"
)

type Wallet struct {
	id                vo.ID             `json:"id"`
	userID            vo.ID             `json:"userId"`
	blockchainAddress string            `json:"blockchainAddress"`
	privateKey        *ecdsa.PrivateKey `json:"privateKey"`
	publicKey         *ecdsa.PublicKey  `json:"publicKey"`
	createdAt         vo.CreatedAt      `json:"createdAt"`
	updatedAt         vo.UpdatedAt      `json:"updatedAt"`
}

func (w *Wallet) Id() vo.ID {
	return w.id
}

func (w *Wallet) UserID() vo.ID {
	return w.userID
}

func (w *Wallet) CreatedAt() vo.CreatedAt {
	return w.createdAt
}

func (w *Wallet) UpdatedAt() vo.UpdatedAt {
	return w.updatedAt
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
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

func NewWallet(userID uint32) (*Wallet, error) {
	w := &Wallet{}

	w.id = vo.NewID()
	w.userID = vo.ID(userID)
	w.createdAt = vo.NewCreatedAt()
	w.updatedAt = vo.NewUpdatedAt()

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
	w.blockchainAddress = address

	return w, nil
}

//==================================================
//リレーション用のWaller
//==================================================

type GetWallet struct {
	id                vo.ID        `json:"id"`
	userID            vo.ID        `json:"userId"`
	blockchainAddress string       `json:"blockchainAddress"`
	createdAt         vo.CreatedAt `json:"createdAt"`
	updatedAt         vo.UpdatedAt `json:"updatedAt"`
}

func (g GetWallet) Id() vo.ID {
	return g.id
}

func (g GetWallet) UserID() vo.ID {
	return g.userID
}

func (g GetWallet) BlockchainAddress() string {
	return g.blockchainAddress
}

func (g GetWallet) CreatedAt() vo.CreatedAt {
	return g.createdAt
}

func (g GetWallet) UpdatedAt() vo.UpdatedAt {
	return g.updatedAt
}

func NewGetWallet(id uint32, userID uint32, blockchainAddress string, createdAt time.Time, updatedAt time.Time) (*GetWallet, error) {
	return &GetWallet{
		id:                vo.ID(id),
		userID:            vo.ID(userID),
		blockchainAddress: blockchainAddress,
		createdAt:         vo.CreatedAt(createdAt),
		updatedAt:         vo.UpdatedAt(updatedAt),
	}, nil
}
