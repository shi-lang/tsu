package main

import "fmt"

func main() {
	ast := Parse("var[x 5] [a b c] # some comment")
	for _, obj := range ast {
		fmt.Printf("%s\n", obj)
	}
}
