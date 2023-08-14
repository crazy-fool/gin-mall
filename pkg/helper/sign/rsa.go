package sign

import (
	"crypto/rsa"
	"fmt"
	"sort"
	"strings"
)

type UtilSign interface {
	GenerateSign(str string) (string, error)
	VerifySign(origin string, sign string) bool
	SortParams(params map[string]any) string
}

type rsaSign struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func (r rsaSign) GenerateSign(str string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r rsaSign) VerifySign(origin string, sign string) bool {
	//TODO implement me
	panic("implement me")
}

func (r rsaSign) SortParams(params map[string]any) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sortedParams := make([]string, 0, len(params))
	for _, key := range keys {
		value := params[key]
		if key == "sign" || value == nil || value == "" {
			continue
		}

		switch v := value.(type) {
		case int, uint, int16, int32, int64:
			kv := fmt.Sprintf("%s=%d", key, v)
			sortedParams = append(sortedParams, kv)
		case float64, float32:
			kv := fmt.Sprintf("%s=%f", key, v)
			sortedParams = append(sortedParams, kv)
		default:
			kv := fmt.Sprintf("%s=%s", key, v.(string))
			sortedParams = append(sortedParams, kv)
		}
	}

	return strings.Join(sortedParams, "&")
}

func NewRsaSign(pubKey string, priKey string) UtilSign {
	sign := new(rsaSign)
	if pubKey != "" {
		sign.publicKey = nil
	}
	if pubKey != "" {
		sign.publicKey = nil
	}
	return sign
}
