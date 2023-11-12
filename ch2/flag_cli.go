package main

import (
	"flag"
	"fmt"
)

// var name = flag.String("name", "World", "A name to say hello to.")
var name string
var spanish bool
var french bool

// & is a pointer to a variable... * will de-reference the pointer
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
	flag.BoolVar(&french, "french", false, "Use French language")
	flag.BoolVar(&french, "f", false, "Use French language")
	flag.StringVar(&name, "name", "World", "A name to say hello to")
	flag.StringVar(&name, "n", "World", "A name to say hello to")
}

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", name)
	} else {
		if french == true {
			fmt.Printf("Salut %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
	}
}
