//permutations examples

package main

import (
	"fmt"
)

func main() {
	results := permutations("abc")
	fmt.Println(results)
}

// We assume they are unique
func permutations(w string) []string {
	if w == "" {
		return []string{" "}
	}
	perms := []string{}
	for i, r := range w {
		remain := w[:i] + w[i+1:]
		//fmt.Println(remain)
		for _, result := range permutations(remain) {
			perms = append(perms, fmt.Sprintf("%c", r)+result)
		}
		//perms = append(perms, fmt.Sprintf("%c\n", r))
	}
	return perms
}
