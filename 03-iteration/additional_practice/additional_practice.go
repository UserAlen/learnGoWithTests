package main

import "fmt"

var result_continue string

// continue skips the iteration
// iteration is ONE repetition of something
func RepeatContinue() string {
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		if n == 5 {
			result_continue += fmt.Sprintf("%d", n)
			fmt.Println(n)
			break
		}
		fmt.Println(n)
		result_continue += fmt.Sprintf("%d\n", n)
	}
	return result_continue
}

func main() {
	RepeatContinue()
}
