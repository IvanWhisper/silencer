package core

import (
	"fmt"
	"silencer/utils"
)

// 加密
func Encrypt(inChan <-chan []byte, key string) (<-chan []byte, error) {
	return crypt(inChan, key, utils.AesEncryptECB)
}

// 解密
func Decrypt(inChan <-chan []byte, key string) (<-chan []byte, error) {
	return crypt(inChan, key, utils.AesDecryptECB)
}

// 加解密
func crypt(inChan <-chan []byte, key string, f func([]byte, []byte) []byte) (<-chan []byte, error) {
	outCh := make(chan []byte)
	go func() {
		defer close(outCh)
		for p := range inChan {
			fmt.Println("加密/解密前长度：", len(p))
			en := f(p, []byte(key))
			fmt.Println("加密/解密后长度：", len(en))
			outCh <- en
		}
	}()
	return outCh, nil
}
