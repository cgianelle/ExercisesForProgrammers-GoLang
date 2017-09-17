package main

import (
	"fmt"
	"strconv"
	"errors"
	"os"

	//--This is my common library
	"github.com/cgianelle/go_common"
)

const slicesPerPie = 8

type PizzaParty struct {
	numPeople, numPizza, slicesPer, leftOver uint64
}

func input(people, pizza func() (string, error)) (*PizzaParty, error) {
	peopleCnt, peopleErr := people()
	pizzaCnt, pizzaErr := pizza()

	if peopleErr != nil {
		return nil, peopleErr
	} else if pizzaErr != nil {
		return nil, pizzaErr
	} else {
		peCnt, peErr := strconv.ParseUint(peopleCnt, 10, 64)
		piCnt, piErr := strconv.ParseUint(pizzaCnt, 10, 64)
		if peErr != nil {
			return nil, peErr
		} else if piErr != nil {
			return nil, piErr
		} else if peCnt < 1 || piCnt < 1 {
			return nil, errors.New("invalid people or pizza count")
		} else {
			pizzaParty := new(PizzaParty)
			pizzaParty.numPeople = peCnt
			pizzaParty.numPizza = piCnt
			return pizzaParty, nil
		}
	}
}

func calcSlicesPer(pp *PizzaParty) {
	totalSliceCount := pp.numPizza * slicesPerPie
	pp.slicesPer = totalSliceCount / pp.numPeople
	pp.leftOver = totalSliceCount % pp.numPeople
}

func output(pp *PizzaParty) {
	fmt.Printf("%d people with %d pizza(s)\n", pp.numPeople, pp.numPizza)
	fmt.Printf("Each person gets %d slices of pizza\n", pp.slicesPer)
	fmt.Printf("There are %d leftover slices\n", pp.leftOver)	
}

func main() {
	var my_scanner *common.EFP_Scanner = common.New_EFP_Scanner()
	people := my_scanner.CreateScanner("How many people? ")
	pizza := my_scanner.CreateScanner("How many pizzas do you have? ")

	pizzaParty, err := input(people, pizza)
	if err != nil {
		fmt.Fprintln(os.Stderr, "input error", err)
		return
	}
	calcSlicesPer(pizzaParty)
	output(pizzaParty)
}
