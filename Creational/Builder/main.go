package main

import (
	"fmt"

	pizza "github.com/suleiman-oss/Design_Patterns_in_Go/Creational/Builder/Pizza"
)

func main() {
	/*Without Builder
	p:=newPizza("Small","Thin","Pepperoni")
	*/
	pb := pizza.NewPizzaBuilder()
	p := pb.SetSize("Small").SetCrust("Thin").AddToppings("Pepperoni").Build()
	fmt.Println(p)
}
