package input

import (
	"bufio"
  "fmt"
  "os"
//	"strconv"
	"regexp"
//	"math"
	"math/big"
	"../../numeric/mashine"
)

func NewMashine() mashine.Mashine {

	currentMashine := mashine.Mashine{}

	reader := bufio.NewReader(os.Stdin)
  fmt.Print("Create you own Numbersystem\nEnter the Base:\n")
  baseString, _ := reader.ReadString('\n')

	currentMashine.Base = *strToInt(baseString)

	fmt.Print("Enter the Precision:\n")
	precisionString, _ := reader.ReadString('\n')

	currentMashine.Precision = *strToInt(precisionString)

	fmt.Print("Enter the Exponent:\n")
	exponentString, _ := reader.ReadString('\n')

	currentMashine.Exponent = *strToInt(exponentString)

	fmt.Print("Enter the MaxExponent:\n")
	maxExponentString, _ := reader.ReadString('\n')

	currentMashine.MaxExponent = *strToInt(maxExponentString)

	fmt.Print("Enter the MinExponent:\n")
	minExponentString, _ := reader.ReadString('\n')

	currentMashine.MinExponent= *strToInt(minExponentString)

	return currentMashine
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
