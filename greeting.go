package greeting

import "github.com/k0kubun/pp/v3"

func Greet(name string) error {
	pp.Printf("Greeting %s\n", name)
	return nil
}
