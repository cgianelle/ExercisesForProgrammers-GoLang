package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanLine(message string, scanner *bufio.Scanner) func() (string, error) {
	return func() (string, error) {
		fmt.Print(message)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
                        fmt.Println(err)
			return "", err
		} else {
        		return scanner.Text(), nil
		}
	}
}

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	quote := scanLine("What is the quote? ", scanner)
        author := scanLine("Who is the author? ", scanner)
        
	var q, eq = quote()
        var a, aq = author()

	if aq != nil || eq != nil {
		fmt.Fprintln(os.Stderr, "quote: reading standard input:", eq)
		fmt.Fprintln(os.Stderr, "author: reading standard input:", aq)
	} else {		
		fmt.Println(fmt.Sprintf("%s, says \"%s\"", a, q))		
	}
}




