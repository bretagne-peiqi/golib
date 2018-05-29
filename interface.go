/*an empty interface{} can hold everything, anything; an empty struct{} hold nothing
interface and reflect is highly combined
*/

package main

import (
	"fmt"
	"log"
	"reflect"
)

type Type1 struct {
	Name  string
	count int
}

type Type2 struct {
	Name string
}

func Process(t interface{}) {
	fmt.Println("Type of t: ", reflect.TypeOf(t))

	v, ok := t.(Type2) // How ??
	//v := t.(Type2) // Panic
	log.Fatalf("failed to switch type %v and value is %v \n", ok, v)

	v.Name = "type2"
	fmt.Println("Type of v: ", reflect.TypeOf(v), ":", v, "\n")
}

func main() {
	t := Type1{Name: "type1", count: 32}

	Process(t)
	fmt.Println(t)
}
