package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := strings.Fields(os.Args[1])
	fmt.Println(len(words))
}
