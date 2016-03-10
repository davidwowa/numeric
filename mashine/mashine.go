package mashine

import (
  "fmt"
	"math"
	"math/big"
)

type Mashine struct {
	Base big.Int
	Precision big.Int
	Exponent big.Int
	MaxExponent big.Int
	MinExponent big.Int
}

func (m Mashine) GetNumberOfNormNumbers () *big.Int {
  return big.NewInt(0)
}

func (m Mashine) GetOverflow() *big.Int {
	return new(big.Int).Exp(&m.Base, &m.MaxExponent, nil)
}

func (m Mashine) GetUnderflow() *big.Int {
	return new(big.Int).Exp(&m.Base, &m.MinExponent, nil)
}

func (m Mashine) GetPrecisionChopped() *big.Int {
  one := big.NewInt(1)
  ret := one.Sub(one, &m.Precision)
  return new(big.Int).Exp(&m.Base, ret, nil)
}

func (m Mashine) GetPrecisionRoundedToNearest() *big.Int {
	// return ((1/2)*m.Base)^(1-m.Precision)
	return big.NewInt(0)
}

func (m Mashine) ShowCommonInformations(){
	fmt.Println("Overflow : ", m.GetOverflow())
	fmt.Println("Underflow : ", m.GetUnderflow())
	fmt.Println("Precision choppped : ", m.GetPrecisionChopped())
	fmt.Println("Precision rounded to nearest : ", m.GetPrecisionRoundedToNearest())
	fmt.Println("Amount of numbers : ", m.GetNumberOfNormNumbers())
}

func (m Mashine) ShowCommonHWInformations(){
	fmt.Println("Overflow this hardware: ", math.MaxInt64)
}
