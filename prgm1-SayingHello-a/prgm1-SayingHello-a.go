package main

import (
	"bufio"
	"fmt"
	"os"
)

func concat(name string) (string) {
       return fmt.Sprintf("Hello, %s, nice to meet you!", name)
}

func output(message string) {
	fmt.Println(message)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("What is your name? ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
        	output(concat(scanner.Text()))
	}
}
