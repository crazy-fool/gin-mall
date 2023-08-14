package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type Pkcs7Cbc struct {
	Key []byte
}

func NewPkcs7Cbc(key []byte) UtilAes {
	return &Pkcs7Cbc{
		Key: key,
	}
}

/*
*

	PHP对应算法：

	$en = openssl_encrypt($params,$this->cipherAlgo, $this->secret,OPENSSL_RAW_DATA,substr($this->secret,0,16));

	return base64_encode($en);
*/
func (a *Pkcs7Cbc) Base64Encrypted(str string) (string, error) {
	encryptedStr, err := a.Encrypt([]byte(str))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedStr), err
}

/*
*
  - - 解密
    php方法：
    $str = base64_decode($encrypt_data);
    $ret = openssl_decrypt($str,$this->cipherAlgo, $this->secret,OPENSSL_RAW_DATA,substr($this->secret,0,16));
*/
func (a *Pkcs7Cbc) Base64Decrypted(str string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	decrypted, err := a.Decrypt(encrypted)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (a *Pkcs7Cbc) Encrypt(origin []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origin = a.pkcs7Padding(origin, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, a.Key[:blockSize])
	crypted := make([]byte, len(origin))
	//执行加密
	blocMode.CryptBlocks(crypted, origin)
	return crypted, nil
}

func (a *Pkcs7Cbc) Decrypt(encryptedData []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, a.Key[:blockSize])
	origData := make([]byte, len(encryptedData))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, encryptedData)

	//去除填充字符串
	origData, err = a.pkcs7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

func (a *Pkcs7Cbc) pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// 填充的反向操作，删除填充字符串
func (a *Pkcs7Cbc) pkcs7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unPadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unPadding)], nil
	}
}
