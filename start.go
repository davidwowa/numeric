package main

import (
    "fmt"
	"../NumericAlgorithms/input"
)

func main() {
  currentMashine := input.NewMashine()
	fmt.Println("Mashine", currentMashine)
	currentMashine.ShowCommonInformations()
	currentMashine.ShowCommonHWInformations()
}
