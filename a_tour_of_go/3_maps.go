package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	//s = strings.ReplaceAll(s, ".", "")
	strings_value := strings.Fields(s)

	m := make(map[string]int)
	fmt.Println("%v", strings_value)
	for i := range strings_value {
		m[strings_value[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
