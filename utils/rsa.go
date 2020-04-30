package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Generator ...
// RSA 生成 公钥 私钥
func Generator(bits int) error {
	privateKey, err := makePrivateKey(bits)
	if nil != err {
		return err
	}
	return makePublicKey(privateKey)
}

// 生成私钥
func makePrivateKey(bits int) (*rsa.PrivateKey, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	blocker := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// try create file
	file, err := os.Create("private.pem")
	if err != nil {
		return nil, err
	}

	err = pem.Encode(file, &blocker)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// 生成配对的公钥
func makePublicKey(privateKey *rsa.PrivateKey) error {
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	blocker := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, &blocker)
	if err != nil {
		return err
	}
	return nil
}

//RSA加密
// plainText 要加密的数据
// path 公钥匙文件地址
func RSA_Encrypt(plainText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA解密
// cipherText 需要解密的byte数据
// path 私钥文件路径
func RSA_Decrypt(cipherText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}
