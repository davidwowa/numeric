package input

import (
	"bufio"
  "fmt"
  "os"
	"regexp"
	"strconv"
	"../../numeric/machine"
)

func NewMachine() machine.Machine {

	currentMachine := machine.Machine{}

	reader := bufio.NewReader(os.Stdin)
  fmt.Print("Create machine\nx=±m*b^(±e)\nm=mantise\nb=basis\n±e=max and min exponent\n")

	fmt.Print("Enter the base:\n")
	baseString, _ := reader.ReadString('\n')

	currentMachine.Base = strToInt(baseString)

	fmt.Print("Enter the precision of mantise:\n")
	precisionMantiseString, _ := reader.ReadString('\n')

	currentMachine.PrecisionMantisse = strToInt(precisionMantiseString)

	fmt.Print("Enter the precision of exponent:\n")
	precisionExponentString, _ := reader.ReadString('\n')

	currentMachine.PrecisionExponent = strToInt(precisionExponentString)

	fmt.Print("Enter the MaxExponent:\n")
	maxExponentString, _ := reader.ReadString('\n')

	currentMachine.MaxExponent = strToInt(maxExponentString)

	fmt.Print("Enter the MinExponent:\n")
	minExponentString, _ := reader.ReadString('\n')

	currentMachine.MinExponent = strToInt(minExponentString)

	return currentMachine
}

func strToInt(str string) float64 {
	str = checkString(str)
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

func checkString(str string) string {
	var leadingInt = regexp.MustCompile(`[0-9]+\.[0-9]+`)

	str = leadingInt.FindString(str)
    if str == "" {
        fmt.Println("is not a number", str)
		return ""
    }
	return str
}
