package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"git.dian.so/leto/util/byte2str"
)

type AesType uint8

const (
	Aes128 AesType = iota
	Aes192
	Aes256
)

// AesEncrypt aes加密
func AesEncrypt(data, key []byte, t AesType) (res []byte, err error) {
	key = aesKeyDeal(key, t)
	res, err = aesEncry(data, key)
	return
}

// AesDecrypt aes解密
func AesDecrypt(data, key []byte, t AesType) (res []byte, err error) {
	key = aesKeyDeal(key, t)
	if res, err = aesDecry(data, key); err != nil {
		return nil, err
	}
	return byte2str.ByteDelZero(res), err
}

func aesDecry(crypted, key []byte) (b []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			b = nil
			err = errors.New("parse error")
			return
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return pkcs5Subbing(origData, blockSize), nil
}

func aesKeyDeal(key []byte, t AesType) (keyReal []byte) {
	keyLen := 0
	if t == 0 {
		keyLen = 16
	} else if t == 1 {
		keyLen = 24
	} else {
		keyLen = 32
	}
	keyReal = make([]byte, keyLen)
	copy(keyReal, key)
	return
}

func aesEncry(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Subbing(ciphertext []byte, blockSize int) []byte {
	l := len(ciphertext)
	if int(ciphertext[l-1]) > blockSize {
		return nil
	}
	return ciphertext[:l-int(ciphertext[l-1])]
}
