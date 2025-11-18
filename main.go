package main

import "fmt"

func main() {
	//fmt.Println("Hello, Go!")
	//var s1 string
	//s1 := "Learn Go!"
	//fmt.Println(s1)

	say("Hello Go!")
	conditionalExample(10)
	demoarray()
}

func demoarray() {
	//Arrays
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("Array elements are:", arr[0], arr[1], arr[2])
}

func conditionalExample(val_01 int) {
	{
		if val_01 > 5 {
			fmt.Println("Value is greater than 5")
		} else {
			fmt.Println("Value is less than or equal to 5")
		}
	}

}

func say(message string) {
	fmt.Println("You said: ", message)
}
