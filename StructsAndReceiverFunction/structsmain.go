package main

import (
	"fmt"
)

func main() {
	mybill := newBill("maryio's bill")

	fmt.Println(mybill)
	fmt.Println(mybill.format())
}
