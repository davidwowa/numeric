package mashine

import (
  "fmt"
	"math/big"
)

type Machine struct {
	PrecisionMantisse big.Int
  PrecisionExponent big.Int
  MaxExponent big.Int
	MinExponent big.Int
}

// biggest number to show
// x_max = (1-b^(-r))*(b*^(b^s -1))
func (m Machine) GetBiggestNumber() *big.Int {
  one := new(big.Int).SetInt64(1)
  minusOne := new(big.Int).SetInt64(-1)
  minusr := one.Mul(&m.PrecisionMantisse, minusOne)
  br := new(big.Int).Exp(&m.PrecisionMantisse, minusr, new(big.Int).SetInt64(0))
  fmt.Println("br ->", br)
  return minusr
}
// TODO WDZ
func (m Machine) GetNumberOfNormNumbers () *big.Int {
  return big.NewInt(0)
}

func (m Machine) GetOverflow() *big.Int {
	return new(big.Int).Exp(&m.PrecisionMantisse, &m.MaxExponent, nil)
}

func (m Machine) GetUnderflow() *big.Int {
	return new(big.Int).Exp(&m.PrecisionMantisse, &m.MinExponent, nil)
}

func (m Machine) GetPrecisionChopped() *big.Int {
  one := big.NewInt(1)
  ret := one.Sub(one, &m.PrecisionMantisse)
  return new(big.Int).Exp(&m.PrecisionMantisse, ret, nil)
}

func (m Machine) GetPrecisionRoundedToNearest() *big.Int {
	// return ((1/2)*m.Base)^(1-m.Precision)
	return big.NewInt(0)
}

func (m Machine) ShowCommonInformations(){
	fmt.Println("Precision of mantise(r) : ", m.PrecisionMantisse)
	fmt.Println("Precision of exponent(s) : ", m.PrecisionExponent)
	fmt.Println("Max exponent : ", m.MaxExponent)
	fmt.Println("Min exponent : ", m.MinExponent)
}
