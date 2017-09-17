package main

import (
	"bufio"
	"fmt"	
	"os"
	"strconv"

	//--This is my common library
	"github.com/cgianelle/go_common"
)

type Operation int

const (
	ADD Operation = iota
	SUB
	MULT
	DIV
)

func convertStrToInt(str1, str2 string) (int, int, error) {
	x, errX := strconv.Atoi(str1)
	y, errY := strconv.Atoi(str2)

	if errX != nil {
		return -1, -1, errX
	} else if errY != nil {
		return -1, -1, errY
	} else {
		return x, y, nil
	}	
}

func add(n1, n2 int) int {
	return n1 + n2
}

func sub(n1, n2 int) int {
	return n1 - n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

func div(n1, n2 int) int {
	return n1 / n2
}

func calculate(n1, n2 int) [4]int {
	return [...]int{ADD: add(n1, n2), SUB: sub(n1, n2), MULT: mult(n1, n2), DIV: div(n1, n2)}	
}

func output(n1, n2 int, arr [4]int) {
	fmt.Printf("%d + %d = %d\n", n1, n2, arr[ADD])
	fmt.Printf("%d - %d = %d\n", n1, n2, arr[SUB])
	fmt.Printf("%d * %d = %d\n", n1, n2, arr[MULT])
	fmt.Printf("%d / %d = %d\n", n1, n2, arr[DIV])
}

func main() {
	var (
		n1,n2 string
		n1e, n2e error
	)

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	num1 := common.ScanLine("What is the first number? ", scanner)
        num2 := common.ScanLine("What is the second number? ", scanner)

	for {
		n1, n1e = num1()
		n2, n2e = num2()

		if n1e != nil || n2e != nil {
			fmt.Fprintln(os.Stderr, "num1: reading standard input:", n1e)
			fmt.Fprintln(os.Stderr, "num2: reading standard input:", n2e)
			break			
		} else if len(n1) < 1 || len(n2) < 1 {
			break
		} else {		
			num1, num2, err := convertStrToInt(n1, n2)
			if err != nil {
				fmt.Fprintln(os.Stderr, "converting str to int: ", err)
				break
			} else {				
				output(num1, num2, calculate(num1, num2))
			}
		}
	}	
}
