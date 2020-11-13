// Generating Random Numbers in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64())

	fmt.Println((rand.Float64() * 6) + 10)
	fmt.Println((rand.Float64() * 9) + 9)
	fmt.Println()
	s1 := rand.NewSource(time.Now().UnixNano())
	p1 := rand.New(s1)

	fmt.Println(p1.Intn(100))
	fmt.Println(p1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(55)
	p2 := rand.New(s2)
	fmt.Println(p2.Intn(100))
	fmt.Println(p2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(5)
	p3 := rand.New(s3)
	fmt.Println(p3.Intn(100))
	fmt.Println(p3.Intn(100))
}
