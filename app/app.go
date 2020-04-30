package app

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"silencer/service"
)

func Start() {
	app := cli.NewApp()
	app.Name = "silencer with go"
	app.Usage = "encryption tool operate by go"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "operation, o",
			Value: "en",
			Usage: "encryption operation",
		},
		cli.StringFlag{
			Name:  "source, s",
			Value: "",
			Usage: "target file",
		},
		cli.StringFlag{
			Name:  "target, t",
			Value: "",
			Usage: "target file",
		},
		cli.StringFlag{
			Name:  "key, k",
			Value: "",
			Usage: "key file",
		},
	}
	app.Action = func(c *cli.Context) error {
		operation := c.String("operation")
		source := c.String("source")
		target := c.String("target")
		key := c.String("key")
		sIsEx, err := exists(source)
		if err != nil {
			return err
		}
		if sIsEx {
			panic("源文件不存在")
		}
		if operation == "en" {
			service.EncryptFile(source, target)
			fmt.Println("Result ", 1)
		} else if operation == "de" {
			service.DecryptFile(source, target, key)
			fmt.Println("Result ", 1)
		} else {
			//help()
			//cli.HelpFlag.Apply()
		}
		return nil
	}
	app.Run(os.Args)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
