package main

import "unicode"

func RemoveEven(a []int) (result []int) {
	result = []int{}
	for i := 0; i < len(a); i++ {
		if a[i] % 2 == 1 {
			result = append(result, a[i])
		}
	}
	return
}

func PowerGenerator(n int) (func() int) {
	base := n
	number := n
	return func() (result int) {
		result = number
		number = number * base
		return
	}
}

func DifferentWordsCount(s string) int {
	rstr := []rune(s)
	words := make(map[string]int)

	current := ""
	for i := 0; i < len(rstr); i++ {
		if unicode.IsLetter(rstr[i]) {
			current += string(unicode.ToLower(rstr[i]))
		} else {
			if len(current) != 0 {
				words[current] = 12345
				current = ""
			}
		}
	}
	if len(current) != 0 {
		words[current] = 12345
		current = ""
	}

	return len(words)
}

// func main(){
// 	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
// }