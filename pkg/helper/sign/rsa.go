package sign

import "crypto/rsa"

type UtilSign interface {
	GenerateSign(str string) (string, error)
	VerifySign(origin string, sign string) bool
	SortParams(params map[string]any) string
}

type rsaSign struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewRsaSign(pubKey string, priKey string) {

}
