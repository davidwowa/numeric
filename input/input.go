package input

import (
	"bufio"
  "fmt"
  "os"
	"regexp"
	"math/big"
	"../../numeric/machine"
)

func NewMachine() machine.Machine {

	currentMachine := machine.Machine{}

	reader := bufio.NewReader(os.Stdin)
  fmt.Print("Create machine\nx=±m*b^(±e)\nm=mantise\nb=basis\n±e=max and min exponent\n")

	fmt.Print("Enter the precision of mantise:\n")
	precisionMantiseString, _ := reader.ReadString('\n')

	currentMachine.PrecisionMantisse = *strToInt(precisionMantiseString)

	fmt.Print("Enter the precision of exponent:\n")
	precisionExponentString, _ := reader.ReadString('\n')

	currentMachine.PrecisionExponent = *strToInt(precisionExponentString)

	fmt.Print("Enter the MaxExponent:\n")
	maxExponentString, _ := reader.ReadString('\n')

	currentMachine.MaxExponent = *strToInt(maxExponentString)

	fmt.Print("Enter the MinExponent:\n")
	minExponentString, _ := reader.ReadString('\n')

	currentMachine.MinExponent= *strToInt(minExponentString)

	return currentMachine
}

func strToInt(str string) *big.Int {
	str = checkString(str)

	i := new(big.Int)
    i.SetString(str, 10)

	return i
}

func checkString(str string) string {
	var leadingInt = regexp.MustCompile(`^[-+]?\d+`)

	str = leadingInt.FindString(str)
    if str == "" {
        fmt.Println("is not a number", str)
		return ""
    }
	return str
}
