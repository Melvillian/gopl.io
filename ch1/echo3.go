// Echo3 prints command-line arguments, written in a succinct manner
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

