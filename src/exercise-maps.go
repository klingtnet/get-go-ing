package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, word := range strings.Fields(s) {
		// the map values are initialized with zero so you can omit the next conditional, but with complex values it could be necessary to initialize them
		if _, ok := wc[word]; ok == false {
			wc[word] = 0
		}
		wc[word] = wc[word] + 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
