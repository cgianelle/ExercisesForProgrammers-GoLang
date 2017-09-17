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
	var (
		a,q string
		eq, ea error
	)
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	quote := scanLine("What is the quote? ", scanner)
        author := scanLine("Who is the author? ", scanner)
        
	for {
		q, eq = quote()
		a, ea = author()

		if ea != nil || eq != nil {
			fmt.Fprintln(os.Stderr, "quote: reading standard input:", eq)
			fmt.Fprintln(os.Stderr, "author: reading standard input:", ea)
			break			
		} else if len(q) < 1 || len(a) < 1 {
			break
		} else {		
			fmt.Println(fmt.Sprintf("%s, says \"%s\"\n", a, q))			
		}
	}
}




