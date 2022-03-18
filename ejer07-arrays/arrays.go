package main

import "fmt"

// var table [10]int
// var matrix [5][5] int
func main()  {
	// table[1] = 5
	// table[5] = 2

	// table := [5]int {11,21,31,41,51}
	// for i := 0; i < len(table); i++ {
	// 	fmt.Printf("Valor %d : %d \n", i+1, table[i])
	// }

	// Slices 
	matrix := []float64 {11.2, 2.44, 1.23}
	fmt.Println(matrix)
	variant()
	variant2()
	variant3()
}

func variant()  {
	elements := [5]int {9,8,7,6,4}
	porcion := elements[2:4]
	fmt.Println(porcion)

}

func variant2()  {
	elements := make([]int, 5, 20) //Type, length and capacity on memory
	// porcion := elements[2:4]
	fmt.Printf("Largo: %d, Capacidad: %d \n",len(elements), cap(elements))
}

func variant3()  {
	nums := make([]int, 0, 0)
	for i := 0; i < 102400; i++ {
		nums=append(nums, i)
	}
	fmt.Printf("Largo: %d, Capacidad: %d \n",len(nums), cap(nums))
}