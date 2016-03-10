package main

import (
	"fmt"
	"unsafe"
	"strconv"
	"math/big"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
  wg.Add(2)

	go func (){
		defer wg.Done()
		for x := 0;  x < 100000; x++ {
			fmt.Println("core 1")
		}
	}()

	go func (){
		defer wg.Done()
		for x := 0; x < 100000; x++ {
			fmt.Println("      core 2")
		}
	}()
	wg.Wait()
}

func test() *big.Int {
	return big.NewInt(3)
}

func add2(){
	fmt.Println("example 1")
	a := 1.
	b := .8
	c := .2
	d := a - b - c
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("Ergebnis: ", d)
}

func add(){
	fmt.Println("example 1")
	a := -1e308
	b := -1e308
	c := -1e308
	d := a + b + c
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("Ergebnis: ", d)
}

func getMaxFloatNumber(){
	fmt.Println("max number : ", 1e308)
}

func testCodeFromVita64(){
	x := float64(1.)
	y := float64(x / float64(3.))
	z := float64(y * float64(3.))
	difference := float64(z - y)
	fmt.Printf("float64(1.0):\t %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&x))), 2))
	fmt.Printf("float64( y ):\t %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&y))), 2))
	fmt.Printf("float64( z ):\t %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&z))), 2))
	fmt.Printf("float64( d ):\t %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&difference))), 2))
}

func testCodeFromVita32(){
	x := float32(1.)
	y := float32(x / float32(3.))
	z := float32(y * float32(3.))
	difference := float64(z - y)
	fmt.Printf("float32(1.0):\t %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&x))), 2))
	fmt.Printf("float32( y ):\t %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&y))), 2))
	fmt.Printf("float32( z ):\t %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&z))), 2))
	fmt.Printf("float32( d ):\t %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&difference))), 2))
}
