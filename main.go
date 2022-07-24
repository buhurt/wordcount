package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	//fmt.Print("Enter text: ")
	wordLine := reader.Text()
	line := strings.TrimSuffix(string(wordLine), "\n")
	if line == "" {
		fmt.Println(0)
	} else {
		fmt.Println(GetCount(line))
	}
}
