package test

import (
	"silencer/utils"
	"testing"
)

func TestGenerator(t *testing.T) {
	//rsa 密钥文件产生
	utils.Generator(1024)
}
