package configstore

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 41
	fmt.Println(intNum)

	// var dog Dog = Dog{name: "Rex", age: 3}

	dog := &Dog{name: "Rex", age: 3}

	utf8.RuneCountInString(dog.name)
}

type Dog struct {
	name string
	age  int
}
