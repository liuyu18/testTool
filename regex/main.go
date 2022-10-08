package toolregex

import (
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"os"
)

func openLocalSourceFile() ([]string, error) {
	f, err := os.OpenFile("regex/source.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file error :", err)
		return nil, err
	}
	defer f.Close()
	_, err = f.Stat()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var sourcesString []string
	for scanner.Scan() {
		sourcesString = append(sourcesString, scanner.Text())
	}
	return sourcesString, nil
}

func ToolRegex() {
	strs, err := openLocalSourceFile()
	if err != nil {
		fmt.Println(err)
	}
	re := regexp2.MustCompile(`(?<=\[)[a-zA-Z]+(?=])`, 0)

	for _, v := range strs {
		if m, _ := re.FindStringMatch(v); m != nil {
			writeToTargetFile(m.String())
		}
	}
}

func writeToTargetFile(s string) {
	f, err := os.OpenFile("result.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}
