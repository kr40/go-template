package main

// This is a comment.
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")

	a := "Kartik"
	fmt.Println(a)

	b, c := 1, 2
	fmt.Println(b, c)

	e := "String"
	fmt.Println(e)

	const n = 360

	fmt.Printf("Type of n is %T\n", n)

	fmt.Println(math.Sin(n))

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
}
