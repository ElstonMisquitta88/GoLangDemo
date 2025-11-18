package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	s1 := "Elston "
	s1 = "Welcome to " + s1
	fmt.Println(s1)

	say("Hello Go!")
	conditionalExample(10)
	demoarray()

	r := rect{width: 3, height: 4}
	measure(r)

	x := 10
	p := &x // p is a pointer to x

	fmt.Println("x =", x)   // 10
	fmt.Println("p =", p)   // memory address
	fmt.Println("*p =", *p) // 10 (value at that memory address)

	x, y := vals(3, 7)
	fmt.Println("X Value :", x)
	fmt.Println("Y Value :", y)
}

func vals(a int, b int) (int, int) {
	return a + 1, b + 1
}

func demoarray() {
	//Arrays
	arr := [3]int{10, 20, 30}
	fmt.Println("Array elements are:", arr[0], arr[1], arr[2])

	//Slice
	arr1 := []int{40, 50, 60}
	fmt.Println("Array elements are:", arr1[0], arr1[1], arr1[2])
}

func conditionalExample(valnum int) {
	{
		if valnum > 5 {
			fmt.Println("Value is greater than 5")
		} else {
			fmt.Println("Value is less than or equal to 5")
		}
	}

}

func say(message string) {
	fmt.Println("You said: ", message)
}
