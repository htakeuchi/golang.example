package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		boolFlag   bool
		intFlag    int
		stringFlag string
	)

	const (
		usage           = "favorite language"
		defaultLanguage = "golang"
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file ..]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&stringFlag, "language", defaultLanguage, usage)
	flag.StringVar(&stringFlag, "l", defaultLanguage, usage+" (shorthand)")
	flag.BoolVar(&boolFlag, "bool", false, "usage of bool flag")
	flag.IntVar(&intFlag, "int", 0, "usage of int flag")

	flag.Parse()

	fmt.Println("language:", stringFlag)
	fmt.Println("bool flag: ", boolFlag)
	fmt.Println("int flag:", intFlag)

	fmt.Println(flag.Args())
}
