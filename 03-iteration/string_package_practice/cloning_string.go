package main

import (
	"fmt"
	"strings"
)

func CloningString(character string, num int) string {
	var result string = character
	for i := 0; i >= num; i++ {
		result += strings.Clone(character)
	}
	return result
}
func main() {
	fmt.Println(CloningString("man", 2))
}
