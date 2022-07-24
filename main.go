package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/buhurt/wordcount"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	wordLine, _ := reader.ReadString('\n')
	fmt.Println(getCount(wordLine))
}
