package main

import (
	"fmt"
	"strconv"
	"os"
	"math/big"

	//--This is my common library
	"github.com/cgianelle/go_common"
)

const TAX_RATE = 0.055

type Item struct {
	quantity uint64
	price float64
}

type Receipt struct {
	subTotal, taxAmount, total float64
	items map[string]Item
	void bool
}

func NewReceipt() *Receipt {
	return &Receipt{items: make(map[string]Item)}
}

func (receipt *Receipt) voidSale() {
	receipt.void = true
}

func (receipt *Receipt) recordSales(item string, quantity uint64, price float64) {
	receipt.items[item] = Item{quantity, price}
}

func (receipt *Receipt) calculateTotals() {
	if !receipt.void {
		for _, item := range receipt.items {		
			receipt.subTotal += (float64(item.quantity) * item.price)
		}

		receipt.taxAmount = receipt.subTotal * TAX_RATE
		receipt.total = receipt.taxAmount + receipt.subTotal
	}
}

func (receipt *Receipt) printReceipt() {
	if !receipt.void {
		fmt.Println("================================================================")
		for name, item := range receipt.items {
			fmt.Printf("Item: %-16s%4d@$%3.2f %8.2f\n",name, item.quantity, item.price, float64(item.quantity) * item.price)
		}
		fmt.Println("----------------------------------------------------------------")
		fmt.Printf("%-16s: $%5.2f\n","Subtotal", receipt.subTotal)
		fmt.Printf("%-16s: $%5.2f\n","Tax", receipt.taxAmount)
		fmt.Printf("%-16s: $%5.2f\n","Total", receipt.total)
		fmt.Println("================================================================")
	} else {
		fmt.Println("Sale was voided")
	}
}

func main() {
	receipt := NewReceipt()
	my_scanner := common.New_EFP_Scanner()
	itemName := my_scanner.CreateScanner("Enter the name of the item: ")
	itemQuantity := my_scanner.CreateScanner("Enter the quantity of the item: ")
	itemPrice := my_scanner.CreateScanner("Enter the price of the item: ")

	for {
		item, ie := itemName()
		if ie != nil {
			fmt.Fprintln(os.Stderr, "item read error", ie)
			return
		}
		
		if len(item) < 1 {			
			break
		}

		if item == "void" {
			receipt.voidSale()
			break
		}
		
		quan, qe := itemQuantity()
		price, pe := itemPrice()

		if qe != nil || pe != nil {
			fmt.Println("Error reading input")
			fmt.Fprintln(os.Stderr, "quantity error", qe)
			fmt.Fprintln(os.Stderr, "price error", pe)
			return
		} else {
			q, errUnit := strconv.ParseUint(quan, 10, 64)
			if errUnit != nil {
				fmt.Printf("Unable to parse quantity, %s\n", quan)
				fmt.Fprintln(os.Stderr, "uint parse error ", errUnit)
				return
			}
			p, _, errFloat := big.ParseFloat(price, 10, 64, big.ToPositiveInf)
			if errFloat != nil {
				fmt.Printf("Unable to parse price, %s\n", price)
				fmt.Fprintln(os.Stderr, "uint parse error ", errFloat)
				return
			}
			p64, _ := p.Float64()
			receipt.recordSales(item, q, p64)
			fmt.Println()
		}
	}

	receipt.calculateTotals()
	receipt.printReceipt()
}
