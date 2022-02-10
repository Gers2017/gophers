package examples

import (
	"fmt"
	"math"
	"strings"
)

func VecCLI()  {
	fmt.Println(strings.Repeat("-*-", 12))
	
	fmt.Println("Please enter the x value")
	x := 0.0
	fmt.Scanln(&x)
	
	fmt.Println("Please enter the y value")
	y := 0.0
	fmt.Scanln(&y)
	
	magnitude := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	angle := 180 * math.Atan2(y, x) / math.Pi
	
	fmt.Printf("The magnitude [%v], angle [%v°] \n", magnitude, angle)
	fmt.Println(strings.Repeat("-*-", 12))
}