package main

import (
	"fmt"
	"gin-mall/pkg/helper/aesutil"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	aes := aesutil.NewPkcs7Cbc(key)
	origin := "123456789"
	encrypted, err := aes.Base64Encrypted(origin)
	fmt.Println("加密：", encrypted)
	//encrypted = "tVIs1T89WnKMXFMUdO/BUA=="
	decrypted, err := aes.Base64Decrypted(encrypted)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ccccccc", err, decrypted)

}
