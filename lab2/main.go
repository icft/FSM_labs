package main

import (
	"fmt"
	"lab2/AST"
)

func main() {
	var s = "b|ca+(())(65:kodw)(f)b{1,2}a\\2(565:moemveo)"
	str := AST.AddConcatenations(s)
	tokens := AST.CreateTokens(str)
	fmt.Println(str)
	fmt.Println(tokens)
	fmt.Println(AST.CreateRPN(tokens))
}
