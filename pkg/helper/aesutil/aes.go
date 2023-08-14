package aesutil

type UtilAes interface {
	Encrypt(origin []byte) ([]byte, error)
	Decrypt(encryptedData []byte) ([]byte, error)
	Base64Encrypted(str string) (string, error)
	Base64Decrypted(str string) (string, error)
}
