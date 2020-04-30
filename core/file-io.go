package core

import (
	"fmt"
	"io"
	"os"
)

// 缓冲区
const BLOCK_SIZE = 1024 * 1024

// 读取文件
func Read(inFile string, size int64) (<-chan []byte, error) {
	outCh := make(chan []byte)
	go func() {
		in, err := os.Open(inFile)
		if err != nil {
			return
		}
		defer fmt.Println("close end")
		defer in.Close()
		defer close(outCh)
		defer fmt.Println("close")
		for {
			buf := make([]byte, size)
			rn, err := in.Read(buf)
			if err != nil && err != io.EOF {
				break
			}
			if err == io.EOF {
				break
			}
			outCh <- buf[:rn]
			fmt.Printf("read %d", rn)
		}
	}()
	return outCh, nil
}

// 写入文件
func Write(inChan <-chan []byte, outFile string) error {
	out, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer out.Close()
	var index int64
	for p := range inChan {
		_, err = out.WriteAt(p, index)
		index += int64(len(p))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("已经解密写入文件：", index)
	return nil
}

func WriteString(data, outfile string) error {
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()
	l, err := out.WriteString(data)
	if err != nil {
		return err
	}
	fmt.Println(l, "bytes written successfully")
	return nil
}
