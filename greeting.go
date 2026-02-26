// Greeting
// This is a repo for Golang development demo.
// This module prints a greet sentence, with color.
//
// Greet function usage: `greeting.Greet("Your name")
package greeting

import "github.com/k0kubun/pp/v3"

// Greet is a function that print greeting message with color
func Greet(name string) error {
	pp.Printf("Greeting %s\n", name)
	return nil
}
