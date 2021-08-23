package main

import "fmt"

// for the prints

//pointers
func swap(m1, m8 *int) {
	var temp int
	temp = *m8
	*m8 = *m1
	*m1 = temp
}

func main() {
	var m1 int // declare a variable
	m1 = 2
	fmt.Println((m1))

	//declare multiples variables
	var (
		m3 = 2
		m4 = 3
	)

	m6 := 5

	fmt.Println(m3 + m4 + m6)

	swap(&m3, &m4)

	fmt.Println(m3, m4)
}
