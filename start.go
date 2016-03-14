package main

import (
    "fmt"
    "math"
	"../numeric/input"
)

func main() {
  currentMachine := input.NewMachine()
	fmt.Println("Machine", currentMashine)
	currentMachine.ShowCommonInformations()
	showCommonHWInformations()
}

func showCommonHWInformations(){
	fmt.Println("Overflow this hardware: ", math.MaxInt64)
}
