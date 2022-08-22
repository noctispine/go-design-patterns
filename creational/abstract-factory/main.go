package main

import (
	"fmt"

	"github.com/noctispine/go-design-patterns/creational/abstract-factory/sports"
)

func main() {
	adidasFactory, _ := sports.GetSportsFactory("adidas")
	nikeFactory, _ := sports.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	fmt.Println(nikeShoe, nikeShirt, adidasShoe, adidasShirt)
}
