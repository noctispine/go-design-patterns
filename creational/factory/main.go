package main

import (
	"fmt"

	"github.com/noctispine/go-design-patterns/creational/factory/drums"
)

func main() {
	handDrum, _ := drums.DrumFactory(drums.DrumTypes.HAND, 5)
	marchingDrum, _ := drums.DrumFactory(drums.DrumTypes.MARCHING, 10)
	printDrumFieldValues(handDrum)
	printDrumFieldValues(marchingDrum)

}

func printDrumFieldValues(drum drums.IDrum) {
	fmt.Printf("name: %s\n", drum.GetName())
	fmt.Printf("age: %d\n", drum.GetAge())
}
