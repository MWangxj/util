package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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
	data, key = aesDataDeal(data, key, t)
	res, err = aesEncry(data, key)
	return
}

// AesDecrypt aes解密
func AesDecrypt(data, key []byte, t AesType) (res []byte, err error) {
	data, key = aesDataDeal(data, key, t)
	if res, err = aesDecry(data, key); err != nil {
		return nil, err
	}
	return byte2str.ByteDelZero(res), err
}

func aesDecry(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return origData, nil
}

func aesDataDeal(orignData, key []byte, t AesType) (orignDataReal, keyReal []byte) {
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
	orignLen := len(orignData)
	if orignLen%keyLen != 0 {
		orignLen = (orignLen/keyLen + 1) * keyLen
	}
	orignDataReal = make([]byte, orignLen)
	copy(orignDataReal, orignData)
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
