package sys_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 加密
func Encrypt(raw_bytes []byte) string {

	// 定义加密密钥和明文 16字节
	key := []byte("147258369abcdef0")
	//raw_bytes := []byte(raw)
	// 创建一个AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 创建一个密码分组模式
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// 加密明文
	ciphertext := make([]byte, len(raw_bytes))
	stream.XORKeyStream(ciphertext, raw_bytes)
	// 输出加密结果
	//fmt.Printf("%x\n", ciphertext)
	return string(ciphertext)
}

// 解密
func Decrypt(enc string) string {
	// 定义加密密钥和明文 16字节
	key := []byte("147258369abcdef0")
	ciphertext := []byte(enc)
	// 创建一个AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 创建一个密码分组模式
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// 解密密文
	decrypted := make([]byte, len(ciphertext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(decrypted, ciphertext)

	// 输出解密结果
	fmt.Printf("%s\n", decrypted)
	return string(decrypted)
}

/*
type AESCipher struct {
}
func prepare()  {

}
*/
