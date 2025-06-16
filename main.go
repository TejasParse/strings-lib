package main

import (
	"fmt"

	reverse "github.com/TejasParse/strings-lib/reverse"
)

func main() {

	strVar := "This is Tejas!"

	var reversedStr string = reverse.Reverse(strVar)

	x := 10
	y := 20

	fmt.Println("Result of Addition: ", reverse.Addition(x, y))

	fmt.Println("Output: ", reversedStr)

}
