package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%x - %b \n", 42, 42)
	fmt.Printf("%#x - %b \n", 42, 42)

	//print decimal, hex, binary, oct
	for i := 0; i < 200; i++ {
		fmt.Printf("%d - %#x - %b - %x \n",
			i, i, i, i)
	}
	//godoc.org/fmt - lists all the formatting options.
	//deciaml, binary, hex, - seperated by tabs
	for j := 0; j < 200; j++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", j, j, j, j)
	}
}
