package main

import "fmt"

func plus(a, b int) int {
	return a + b
}
func plusv(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 4, 5
}
func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	a, b := 2, 4
	fmt.Println(a, b)
	name := "Welcome to go Programming language"
	fmt.Println(name)
	var i = 0
	for i <= 50 {
		if i%2 == 0 {

			fmt.Println(i)
		}
		i = i + 1
	}
	//Arrays in the golang
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 5
	a[4] = 4

	fmt.Println("The value in the array is ", a)
	fmt.Println("The size of array is", len(a))
	b := [5]int{5, 6, 7, 8, 9}
	fmt.Println("idx:", b)
	v := []int{6, 7, 8, 9, 0}
	v2 := []int{12, 14, 14, 17}
	v = append(v, v2...)
	fmt.Println(v)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 6
	fmt.Println("value of the map is", m)
	res := plus(2, 3)
	fmt.Printf("The value of the summation %v\n", res)
	res2 := plusv(5, 6, 7)
	fmt.Printf("The value of the second summation is %v\n", res2)
	a, b := vals()
	_, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	total := 0
	i := 0
	for i <= 50 {
		total += i
		i += 1
	}
	fmt.Println(total)
	sums(1, 2)
	sums(2, 3, 4)

}
