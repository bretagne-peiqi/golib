package main

// the interest lies in f4 and f2.
import (
	"fmt"
)

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f2() (t int) {
	t = 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f4() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func main() {

	x := f4()
	fmt.Printf("return value is %v \n", x)
}
