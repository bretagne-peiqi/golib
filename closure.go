package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {
	// returns an anonymous function
	return func(message string) {
		fmt.Println(message)
	}
}

func outer(name string) {
	// variable in outer function
	text := "Modified " + name

	// foo is a inner function and has access to text variable, is a closure
	// closures have access to variables even after exiting this block
	foo := func() {
		fmt.Println(text)
	}

	// calling the closure
	foo()
}

func main() {
	// named function
	printMessage("Hello function!")

	// anonymous function declared and called
	k := func(message string) bool {
		fmt.Println(message)
		return false
	}

	if flag := k("testing ..."); !flag {
		fmt.Printf("flag is flase \n")
	}

	// gets anonymous function and calls it
	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller!")

}
