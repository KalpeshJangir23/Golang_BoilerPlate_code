package main

//-- Every Go program is made up of packages.
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("hello world")
	fmt.Println(math.Sqrt(2))
	fmt.Println(rand.Intn(12))
}

