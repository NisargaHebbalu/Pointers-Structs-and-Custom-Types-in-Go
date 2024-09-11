package main

import "fmt"

type bill struct {
	name  string
	items map[string]int
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]int{"salad": 150, "coffee": 90},
		tip:   0,
	}

	return b
}

// receiver function
// format the bill
func (b bill) format() string {
	fs := "Bill Breakdown:\n"
	var total int = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v.....%v \n", k+":", v) //%-25v makes the variable have 25 character where the empty spaces will be filled with spaces
		total += v
	}

	fs += fmt.Sprintf("%-25v.....%v", "total:", total)

	return fs

}
