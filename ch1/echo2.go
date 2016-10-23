// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", ""
	for _, arg := range os.Args[1:] { // _ is for the unused index value return by range
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
