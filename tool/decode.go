package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"yinliuService/constant"
)

// 使用sha256算法来进行hash
func Sha256(str string) string {
	sumArr := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sumArr[:])
}

/**
** 使用pbkdf2算法对密码进行加密
 */
func PBKDF2Encode(password string) string {
	passwordBytes := pbkdf2.Key([]byte(password), constant.PasswordSalt, 4096, 32, sha256.New)
	return hex.EncodeToString(passwordBytes) // 最终生成的密码将占用64 byte
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
	cipherBlock, err := aes.NewCipher(constant.TokenSalt)
	if err != nil {
		ErrLog(err)
		return 0, nil
	}
	cipherSize = cipherBlock.BlockSize()
	if isEncode {
		cipherMode = cipher.NewCBCEncrypter(cipherBlock, constant.TokenSalt[:cipherSize])
	} else {
		cipherMode = cipher.NewCBCDecrypter(cipherBlock, constant.TokenSalt[:cipherSize])
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

func ByteSlice2Str(byteSlice []byte) string {
	result := hex.EncodeToString(byteSlice)
	return result
}