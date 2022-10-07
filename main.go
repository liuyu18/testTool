package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello world")
	for i := 0; i < 256; i++ {
		variant := fmt.Sprintf("%02x\n", i)
		generalWrite(strings.Replace(variant, "\n", "", -1))
	}
}

func generalWrite(param string) {
	f, err := os.OpenFile("text.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error :", err)
		return
	}
	defer f.Close()
	fmt.Println(param)
	_, err = f.WriteString("\\x" + param)
	if err != nil {
		fmt.Println(err)
		return
	}
}
