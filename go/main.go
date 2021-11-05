package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.CommandLine.Parse(os.Args[1:])
	fmt.Println("Hello CLI")
	fmt.Println(flag.Args())
}
