package main

import (
	"fmt"

	factory "github.com/noctispine/go-design-patterns/creational/factory/drums"
)

func main() {
	handDrum, _ := factory.DrumFactory(factory.DrumTypes.HAND, 5)
	marchingDrum, _ := factory.DrumFactory(factory.DrumTypes.MARCHING, 10)
	printDrumFieldValues(handDrum)
	printDrumFieldValues(marchingDrum)

}

func printDrumFieldValues(drum factory.IDrum) {
	fmt.Printf("name: %s\n", drum.GetName())
	fmt.Printf("age: %d\n", drum.GetAge())
}
