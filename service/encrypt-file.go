package service

import (
	"fmt"
	"io/ioutil"
	"silencer/core"
	"silencer/utils"
)

func EncryptFile(source, target string) {
	key := utils.GetRandomString(32)
	core.WriteString(key, "key.txt")
	// 读文件
	crd, _ := core.Read(source, core.BLOCK_SIZE)
	// 加密
	ench, _ := core.Encrypt(crd, key)
	// 写文件
	core.Write(ench, target)
}

func DecryptFile(source, target, keyfile string) {
	if len(keyfile) == 0 {
		keyfile = "key.txt"
	}
	data, err := ioutil.ReadFile(keyfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	key := string(data)
	// 读文件
	d, _ := core.Read(source, core.BLOCK_SIZE+16)
	// 解密
	dech, _ := core.Decrypt(d, key)
	core.Write(dech, target)
}
