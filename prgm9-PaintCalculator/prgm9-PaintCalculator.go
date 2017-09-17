package main

import (
	"fmt"
	"errors"
	"os"
	"math/big"
	"math"

	//--This is my common library
	"github.com/cgianelle/go_common"
)

const PAINT_COVERAGE = 350.0

var zeroFloat *big.Float = big.NewFloat(0)

type PaintNeeded struct {
	length, width, area float64
	gallonsNeeded uint64
}

func input(length, width func() (string, error)) (*PaintNeeded, error) {
	ceilLen, lenErr := length()
	ceilWidth, widErr := width()

	if lenErr != nil {
		return nil, lenErr
	} else if widErr != nil {
		return nil, widErr
	} else {
		cl, _, clErr := big.ParseFloat(ceilLen, 10, 64, big.ToPositiveInf)
		cw, _, cwErr := big.ParseFloat(ceilWidth, 10, 64, big.ToPositiveInf)
		if clErr != nil {
			return nil, clErr
		} else if cwErr != nil {
			return nil, cwErr
		} else {
			if cl.Cmp(zeroFloat) <= 0 || cw.Cmp(zeroFloat) <= 0 {
				return nil, errors.New("invalid width or length, must be greater than 0")
			} else {
				pn := new(PaintNeeded)
				pn.length, _ = cl.Float64()
				pn.width, _ = cw.Float64()
				return pn, nil
			}
		}
	}
}

func calculatePaintNeeded(pn *PaintNeeded) {
	pn.area = pn.length * pn.width	
	gn := math.Ceil(pn.area / PAINT_COVERAGE)
	gallons := big.NewFloat(gn)
	pn.gallonsNeeded, _ = gallons.Uint64()
}

func output(pn *PaintNeeded) {
	fmt.Printf("You will need to purchase %d gallons\n", pn.gallonsNeeded)
	fmt.Printf("of paint to cover %.3f square feet\n", pn.area)	
}

func main() {
	var my_scanner *common.EFP_Scanner = common.New_EFP_Scanner()
	length := my_scanner.CreateScanner("Width of the ceiling in feet? ")
	width := my_scanner.CreateScanner("Length of the ceiling in feet? ")
	
	pn, err := input(length, width)
	if err != nil {
		fmt.Fprintln(os.Stderr, "input error", err)
		return		
	}
	calculatePaintNeeded(pn)
	output(pn)
}
