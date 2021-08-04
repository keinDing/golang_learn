package main

import (
	"converter/temp"
	"fmt"
)

func init() {
	fmt.Println("Converter program initialized...")
}

func main() {
	for i := 0; i <= 50; i++ {
		f := temp.CToF(float64(i))
		fmt.Printf("%f Celsius = %f Fahrenheit\n", float64(i), f)
	}
}
