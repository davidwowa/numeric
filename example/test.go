package main

import (
    "fmt"
    "math/big"
)

func main() {
    two := big.NewInt(2)
    hundred := big.NewInt(5)
    fmt.Printf("2 ** 100    is %d\n", ExpByPowOfTwo(two, hundred))
}

func ExpByPowOfTwo(base, power *big.Int) *big.Int {
    result := big.NewInt(1)
    zero := big.NewInt(0)
    for power != zero {
        if modBy2(power) != zero {
            multiply(result, base)
        }
        power = divideBy2(power)
        base = multiply(base, base)
    }
    return result
}

func modBy2(x *big.Int) *big.Int {
    return big.NewInt(0).Mod(x, big.NewInt(2))
}

func divideBy2(x *big.Int) *big.Int {
    return big.NewInt(0).Div(x, big.NewInt(2))
}

func multiply(x, y *big.Int) *big.Int {
    return big.NewInt(0).Mul(x, y)
}
