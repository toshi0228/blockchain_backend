package entity

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int // 公開鍵　x座標
	S *big.Int // 送信者の秘密鍵から導き出した値
}

func NewSignature(r *big.Int, s *big.Int) *Signature {
	return &Signature{r, s}
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
