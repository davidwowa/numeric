package main

import (
    "fmt"
	"../numeric/input"
)

func main() {
  currentMashine := input.NewMashine()
	fmt.Println("Mashine", currentMashine)
	currentMashine.ShowCommonInformations()
	currentMashine.ShowCommonHWInformations()
}
