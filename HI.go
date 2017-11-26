package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi Mom")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)

		text = strings.TrimRight(text, "\n")

		if text == "Break" {
			fmt.Println("We broke")
			break
		}
	}
}

func sum(x, y int) (sum int) {
	sum = x + y

	return
}
