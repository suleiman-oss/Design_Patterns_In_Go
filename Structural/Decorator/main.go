package main

import (
	"fmt"

	icecream "github.com/suleiman-oss/Design_Patterns_in_Go/Structural/Decorator/IceCream"
)

func main() {
	i := icecream.NewCream(icecream.NewCone())
	fmt.Printf("The cost is %v\n", i.GetCost())
	fmt.Printf("The description is %v", i.GetDescription())
}
