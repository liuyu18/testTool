package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

var dictName = pflag.StringP("dictName", "d", "Dict.txt", "Input dict file name")
var shellName = pflag.StringP("shellName", "s", "shell.php", "Input shell file name")

func main() {
	pflag.Parse()
	for i := 0; i < 100; i++ {
		file := *shellName + strconv.Itoa(i)
		fmt.Println(file)
		hashMd5 := md5.New()
		hashMd5.Write([]byte(file))
		hashMd5String := hex.EncodeToString(hashMd5.Sum(nil))
		fmt.Println(hashMd5String)
		f, err := os.OpenFile(*dictName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.WriteString(hashMd5String + ".php" + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}
