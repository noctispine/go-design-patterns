package main

import (
	"fmt"

	builder "github.com/noctispine/go-design-patterns/creational/builder/drums"
)

func main() {
	dBuilder := builder.NewDrumBuilder()
	dBuilder.Name("yoyo")
	dBuilder.DrumType(builder.DrumTypes.HAND)
	drum := dBuilder.Build()
	printDrumDetails(drum)
}

func printDrumDetails(d builder.IDrum) {
	fmt.Printf("Name %s\n", d.GetName())
	fmt.Printf("Type: %s\n", d.GetType())
	fmt.Printf("Material: %s\n", d.GetMaterial())
}
