package greet
// this (package golangmodules) is a reusable library package and not an executable program (unlike main).

import "strings"

// exported functions start with Capital letter
// other packages can call it -> p1, p2
func Hello (name string) string {
	
	clean := normalizeName(name)
	return "Hello, " + clean 

}

func normalizeName (name string) string {
	n := strings.TrimSpace(name)
	if n == "" {
		return "Guest"
	}
	return strings.ToUpper(n)
}