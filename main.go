package main

import (
	"fmt"
	// math package
	"math"
)

// parameter/argument - "n" stands for "name", and the argument is a string
func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("goodbye %v \n", n)
}

// potentially we could pass one of the functions as an argument in here
func cycleNames(n []string, f func(string)) {
	// _ = empty index; v = value
	for _, v := range n {
		// function(value)
		f(v)
	}
}

// r = radius
// if we are returning a value, we have to say what type of value we are returning,
// that's why we specify float64 before { } brackets
func circleArea(r float64) float64 {
	// formula for Pi
	return math.Pi * r * r
}

func main() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
	// // fmt.Println("Hello World!")

	// // String - string are double-quotes ONLY
	// var nameOne string = "mario"

	// // Go can automatically infer the type by analysing the variable "luigi" and output it as a string
	// var nameTwo = "luigi"

	// // empty string value
	// var nameThree string

	// fmt.Println("first run: "+nameOne, nameTwo, nameThree)

	// nameThree = "bowser"

	// fmt.Println("second run: "+nameOne, nameTwo, nameThree)

	// // no var keyword, the used the colon ":" means var, it is used for var declaration
	// // you cannot use the colon outside of a function
	// namefour := "yoshi"

	// fmt.Println("third run: " + namefour)

	// var ageOne int = 10
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits and memory limits - int16, int8, etc
	// // var numOne int8 = 25
	// // var numTwo int8 = -128

	// // unit can't have negative number - it must be positive.
	// // var numThree uint8 = 25

	// // float HAS TO HAVE bit size - float 32 or float 64
	// // most of the time, we use float64
	// // var scoreOne float32 = 25.98
	// // var scoreTwo float64 = 883242341321231.5
	// // default inference is float64 in Go.
	// // scoreThree := 1.5

	// age := 32
	// name := "Beaumont"

	// fmt.Println("my age is", age, "and my name is", name)

	// //Printf (formated strings) - note: the order matters!
	// // %_ = format specifier

	// // value
	// fmt.Printf("my age is %v and my name is %v \n", age, name)

	// // quotes
	// fmt.Printf("my age is %q and my name is %q \n", age, name)

	// // type of the variable
	// fmt.Printf("age is of type %T \n", age)

	// // no rounding
	// fmt.Printf("your score is %f points! \n", 255.55)

	// // rounding to 1 decimal
	// fmt.Printf("your score is %0.1f points! \n", 255.55)

	// // Sprintf - save the string for use later
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	// fmt.Println("the saved string is:", str)

	// var name [3]int = [3]int{1, 2, 3}

	// names := [4]string{"a", "b", "c", "d"}

	// fmt.Println(name, names)

	// var greetings = []string{"hi", "hello", "yo"}
	// 	// fmt.Println(greetings, len(greetings))

	// 	// slices[2] = "hola"
	// 	// fmt.Println(slices)

	// 	// rangeOne := slices[:1]

	// 	// fmt.Println(rangeOne)

	// 	names := "hello there friends"
	// 	fmt.Println(strings.Contains(names, "hello"))

	// 	fmt.Println(strings.ReplaceAll(names, "hello", "hi"))

	// 	index := sort.StringSlice(greetings)

	// 	fmt.Println(index)

	// x := 0

	//while-loop
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for y := 0; y < len(greetings); y++ {
	// 	fmt.Println(greetings[y])
	// }

	// for-in
	// for index, value := range greetings {
	// 	fmt.Printf("the position at index %v is %v \n", index, value)
	// }

	// for _, value := range greetings {
	// 	fmt.Printf("the value is %v \n", value)
	// 	// note: altering the value here does not alter the actual value of the original string
	// 	value = "new string"
	// }
	// // see console, the value is unchanged
	// fmt.Println(greetings)

	// age := 45

	// fmt.Println(age <= 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	// // position and value
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("breaking at pos", index)
	// 		// break the current continuation and continue with the loop
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		// "break" breaks out of the loop completely.
	// 		break
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n", index, value)
	// }
}
