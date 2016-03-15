package main

import (
    "fmt"
    "math"
	"../numeric/input"
)

func main() {
  currentMachine := input.NewMachine()
	fmt.Println("Machine", currentMachine)
	currentMachine.ShowCommonInformations()
	showCommonHWInformations()
}

func showCommonHWInformations(){
	fmt.Println("Overflow this hardware: ", math.MaxInt64)
}
