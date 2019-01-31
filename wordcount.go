package words

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println("Gink")
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	s := "If it looks like a duck, swims like a duck,and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)

}
