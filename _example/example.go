package main

import (
	"fmt"

	"github.com/mickep76/color"
)

func main() {
	fmt.Printf("%s%s%s%s\n", color.Red, color.Bold, "Hello World!", color.Reset)

	for i := 1; i < 255; i++ {
		fmt.Printf("%s#", color.Fg256(uint8(i)))
	}
	fmt.Printf("%s\n", color.Reset)

	for i := 1; i < 255; i++ {
		fmt.Printf("%s#", color.Bg256(uint8(i)))
	}
	fmt.Printf("%s\n", color.Reset)
}
