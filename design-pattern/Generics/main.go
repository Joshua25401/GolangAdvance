package main

import (
	"fmt"
	"reflect"
)

func printIntOrFloat[V int | float64](num V) {
	fmt.Println("Num = ", num, " type = ", reflect.TypeOf(num))
}

func main() {
	printIntOrFloat(10)
	printIntOrFloat(10.3)
}
