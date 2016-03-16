package machine

import (
  "fmt"
  "math"
)

type Machine struct {
  Base float64
	PrecisionMantisse float64
  PrecisionExponent float64
  MaxExponent float64
	MinExponent float64
}

var (
  minusOne float64 = -1.
)

// x_max = ±(1-b^(-r))*(b*^(b^s -1))
func (m Machine) GetBiggestNumber() float64 {
  // b^(-r) r-> numbers for mantisse
  var minusr float64 = m.PrecisionMantisse * minusOne
  var br float64 = math.Pow(m.Base, minusr)
  // (1-b^(-r))
  var oneMinusBr float64 = 1 - br
  // b^s - 1
  var sMinusOne float64 = m.PrecisionExponent - 1
  // b^(s-1)
  var bSminusOne float64 = math.Pow(m.Base, sMinusOne)
  // b^(b^(s-1))
  var bbs float64 = math.Pow(m.Base, bSminusOne)
  // ---
  var result float64 = oneMinusBr * bbs

  fmt.Println("biggest number ->±", result)
  return result
}

// x_min = ±b^(-b^(s))
func (m Machine) GetSmallestNumber() float64 {
  var minusB float64 = m.Base * minusOne
  var bPows float64 = math.Pow(minusB, m.PrecisionExponent)
  var result float64 = math.Pow(m.Base, bPows)
  fmt.Println("smallest number ->±", result)
  return result
}

// epsilon = (base/2)*base^-precision
func (m Machine) GetMachineEpsilon() float64 {
  var base2 float64 = m.Base / 2.0
  var negPrecision float64 = m.PrecisionMantisse * minusOne
  var powBasePrecision float64 = math.Pow(m.Base, negPrecision)
  var result float64 = base2 * powBasePrecision
  fmt.Println("machine epsilon ->", result)
  return result
}

func (m Machine) ShowCommonInformations(){
  fmt.Println("Base(b) : ", m.PrecisionMantisse)
	fmt.Println("Precision of mantise(r) : ", m.PrecisionMantisse)
	fmt.Println("Precision of exponent(s) : ", m.PrecisionExponent)
	fmt.Println("Max exponent : ", m.MaxExponent)
	fmt.Println("Min exponent : ", m.MinExponent)
}
