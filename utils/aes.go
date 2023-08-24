package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"go-web-example/errorcode"
)

// AesEncrypt AES加密,CBC 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
func AesEncrypt(src, key string) (string, error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)

	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, keyByte)
	blockMode.CryptBlocks(crypted, encryptBytes)

	return fmt.Sprintf("%X", crypted), nil
}

// AesDecrypt AES解密
func AesDecrypt(src, key string) (string, error) {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}

	return string(crypted), nil
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errorcode.New(errorcode.ErrParams, "加密字符串错误", nil)
	}

	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
