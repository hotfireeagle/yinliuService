package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

/**
** 使用sha256算法来生成hash
 */
func Sha256(str string) string {
	sumArr := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sumArr[:])
}

/**
*** 补位
 */
func pkcs7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	otherText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, otherText...)
}

/**
** 取出补位
 */
func pkcs7UnPadding(text []byte) []byte {
	length := len(text)
	del := int(text[length-1])
	return text[:(length - del)]
}

/**
** 获取到密钥的一些信息
 */
func getCipherBlockAndCipherSize(isEncode bool) (cipherSize int, cipherMode cipher.BlockMode) {
	secretByteSlice := []byte(Secret)
	cipherBlock, err := aes.NewCipher(secretByteSlice)
	if err != nil {
		return 0, nil
	}
	cipherSize = cipherBlock.BlockSize()
	if isEncode {
		cipherMode = cipher.NewCBCEncrypter(cipherBlock, secretByteSlice[:cipherSize])
	} else {
		cipherMode = cipher.NewCBCDecrypter(cipherBlock, secretByteSlice[:cipherSize])
	}
	return
}

/**
** 使用aes算法来进行加密
 */
func AesEncode(text []byte) []byte {
	cipherSize, blockMode := getCipherBlockAndCipherSize(true)
	text = pkcs7Padding(text, cipherSize)

	result := make([]byte, len(text))
	blockMode.CryptBlocks(result, text)
	return result
}

/**
** 使用aes算法进行解密
 */
func AesDecode(text []byte) []byte {
	_, blockMode := getCipherBlockAndCipherSize(false)
	resultSlice := make([]byte, len(text))
	blockMode.CryptBlocks(resultSlice, text)
	resultSlice = pkcs7UnPadding(resultSlice)
	return resultSlice
}

/**
** 把byte切片转换成字符串
 */
func ByteSlice2Str(byteSlice []byte) string {
	result := hex.EncodeToString(byteSlice)
	return result
}

/**
** 生成随机字符串
 */
func GenerateRandomString(l int) string {
	dict := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(dict)
	var result []byte
	randomFact := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[randomFact.Intn(len(bytes))])
	}
	str := string(result)
	return Sha256(str)
}
