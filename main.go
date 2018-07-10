package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := strings.Fields(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if matches(text, args) {
			fmt.Println(text)
		}
	}
}
func matches(s string, words []string) bool {
	for _, w := range words {
		if w[0] == '!' && strings.Contains(s, w[1:]) {
			return false
		} else if w[0] != '!' && !strings.Contains(s, w) {
			return false
		}
	}
	return true
}
