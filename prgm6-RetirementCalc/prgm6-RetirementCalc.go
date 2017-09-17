package main

import (
	"bufio"
	"fmt"	
	"os"
	"strconv"
	"time"

	//--This is my common library
	"github.com/cgianelle/go_common"
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


func main() {
	var (
		curAge, retireAge string
		curErr, retireErr error
	)
	var currentYear int = time.Now().Year()
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	currentAge := common.ScanLine("What is your current age? ", scanner)
        retirementAge := common.ScanLine("What age would you like to retire? ", scanner)

	curAge, curErr = currentAge()
	retireAge, retireErr = retirementAge()

	if curErr != nil || retireErr != nil {
		fmt.Fprintln(os.Stderr, "currentAge: reading standard input:", curErr)
		fmt.Fprintln(os.Stderr, "retirementAge: reading standard input:", retireErr)
	} else if len(curAge) < 1 || len(retireAge) < 1 {
		fmt.Println("Good Bye")
	} else {		
		ageNow, ageThen, err := convertStrToInt(curAge, retireAge)
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting str to int: ", err)
		} else if ageNow >= ageThen {
			fmt.Println("You qualify for retirement!")
		} else {
			yearsToRetirement := ageThen - ageNow
			retirementYear := currentYear + yearsToRetirement

			fmt.Printf("You have %d years until retirement\n", yearsToRetirement)
			fmt.Printf("It's %d now, so you can retire in %d\n", currentYear, retirementYear)
		}
	}	
}
