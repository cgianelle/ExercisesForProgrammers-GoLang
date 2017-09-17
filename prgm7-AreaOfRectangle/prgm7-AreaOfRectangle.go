package main

import (
	"bufio"
	"fmt"	
	"os"
	"math/big"	

	//--This is my common library
	"github.com/cgianelle/go_common"
)

var metersSquaredCoversion *big.Float = big.NewFloat(0.09290304)
var zeroFloat *big.Float = big.NewFloat(0)

func convertDimensionsToFloat(length, width string) (*big.Float, *big.Float, error) {
	x, _, errX := big.ParseFloat(length, 10, 64, big.ToPositiveInf)
	y, _, errY := big.ParseFloat(width, 10, 64, big.ToPositiveInf)

	if errX != nil {
		return nil, nil, errX
	} else if errY != nil {
		return nil, nil, errY
	} else {
		return x, y, nil
	}		
}

func calculateAreaInFeet(length, width *big.Float) float64 {
	x, _ := length.Float64()
        y, _ := width.Float64()

	return x * y
}

func calculateAreainMeters(length, width *big.Float) float64 {
	x, _ := length.Float64()
        y, _ := width.Float64()
        c, _ := metersSquaredCoversion.Float64()

	return x * y * c
}

func output(length, width string, areaFeet, areaMeters float64) {
	fmt.Printf("You entered dimensions %s feet by %s feet\n", length, width)
	fmt.Println("The area is")
	fmt.Printf("%.3f square feet\n", areaFeet)
	fmt.Printf("%.3f square meters\n", areaMeters)
}

func main() {
	var (
		lenStr, widthStr string
		lenErr, widthErr error
	)
	
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	roomLength := common.ScanLine("What is the length of your room in feet? ", scanner)
        roomWidth := common.ScanLine("What is the width of your room in feet ", scanner)

	lenStr, lenErr = roomLength()
	widthStr, widthErr = roomWidth()

	if lenErr != nil || widthErr != nil {
		fmt.Fprintln(os.Stderr, "length: reading standard input:", lenErr)
		fmt.Fprintln(os.Stderr, "width: reading standard input:", widthErr)
	} else if len(lenStr) < 1 || len(widthStr) < 1 {
		fmt.Println("Good Bye")
	} else {
		length, width, convertErr := convertDimensionsToFloat(lenStr, widthStr)
		
		if convertErr != nil {
			fmt.Fprintln(os.Stderr, "converting str to float: ", convertErr)
		} else if length.Cmp(zeroFloat) == 0 || width.Cmp(zeroFloat) == 0 {
			fmt.Println("Zero feet is not a valid dimension")
		} else {
			areaFeet := calculateAreaInFeet(length, width)
			areaMeters := calculateAreainMeters(length, width)
			output(lenStr, widthStr, areaFeet, areaMeters)
		}
	}	

}
