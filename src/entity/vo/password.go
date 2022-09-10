package vo

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

type Password string

func NewPassword(plaintext string) Password {
	return Password(plaintext)
}

func (p Password) Value() string {
	return string(p)
}

func (p Password) Equals(p2 Password) bool {
	return p == p2
}

// Encrypt 暗号化
func (p Password) Encrypt() Password {
	cryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(p)))
	return Password(cryptText)
}

// IsValidPassword パスワードのバリデーション
func (p Password) IsValidPassword() bool {

	if len(p) <= 12 {
		return false
	}

	isLowerReg := regexp.MustCompile("[a-z]")
	isLower := isLowerReg.MatchString(p.Value())

	isUpperReg := regexp.MustCompile("[A-Z]")
	isUpper := isUpperReg.MatchString(p.Value())

	isNumberReg := regexp.MustCompile("[0-9]")
	isNumber := isNumberReg.MatchString(p.Value())

	isSymbolReg := regexp.MustCompile("[!$%&'*+./:;<=>?@[\\]^_`|~(\\-),{#}]")
	isSymbol := isSymbolReg.MatchString(p.Value())

	return isLower && isUpper && isNumber && isSymbol
}
