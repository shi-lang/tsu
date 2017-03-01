package main

import "fmt"

func main() {
	ast := Parse("setSlot (x, 4 + (5) ; nil) ; vec(:a, :b, :c) # some comment")
	for _, obj := range ast {
		fmt.Printf("%s\n", obj)
	}
}
