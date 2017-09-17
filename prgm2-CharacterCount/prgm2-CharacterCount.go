package main

import (
	"bufio"
	"fmt"
	"os"
)

func output(inputStr string) {
    var inputLength int = len(inputStr)
    if inputLength >= 1 {
        fmt.Printf("Did you know, '%s', has %d characters?\n", inputStr, inputLength)
    } else {
       fmt.Println("What was that, I couldn't hear you...You must type something")
    }
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("What is your input string? ")
	for scanner.Scan() {
		output(scanner.Text())
                fmt.Print("What is your input string? ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}




