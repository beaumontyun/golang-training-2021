package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// String - string are double-quotes ONLY
	var nameOne string = "mario"

	// Go can automatically infer the type by analysing the variable "luigi" and output it as a string
	var nameTwo = "luigi"

	// empty string value
	var nameThree string

	fmt.Println("first run: "+nameOne, nameTwo, nameThree)

	nameThree = "bowser"

	fmt.Println("second run: "+nameOne, nameTwo, nameThree)

	// no var keyword, the used the colon ":" means var, it is used for var declaration
	// you cannot use the colon outside of a function
	namefour := "yoshi"

	fmt.Println("third run: " + namefour)

	var ageOne int = 10
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory limits - int16, int8, etc
	// var numOne int8 = 25
	// var numTwo int8 = -128

	// unit can't have negative number - it must be positive.
	// var numThree uint8 = 25

	// float HAS TO HAVE bit size - float 32 or float 64
	// most of the time, we use float64
	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 883242341321231.5
	// default inference is float64 in Go.
	// scoreThree := 1.5

	age := 32
	name := "Beaumont"

	fmt.Println("my age is", age, "and my name is", name)

	//Printf (formated strings) - note: the order matters!
	// %_ = format specifier

	// value
	fmt.Printf("my age is %v and my name is %v \n", age, name)

	// quotes
	fmt.Printf("my age is %q and my name is %q \n", age, name)

	// type of the variable
	fmt.Printf("age is of type %T \n", age)

	// no rounding
	fmt.Printf("your score is %f points! \n", 255.55)

	// rounding to 1 decimal
	fmt.Printf("your score is %0.1f points! \n", 255.55)

	// Sprintf - save the string for use later
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	fmt.Println("the saved string is:", str)
}
