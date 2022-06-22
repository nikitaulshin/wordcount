package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src := os.Args[1]
	words := strings.Fields(src)
	fmt.Println(len(words))
}
