package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	for f := 1; f < 100000; f++ {
		echoFast()
	}
	fmt.Println(time.Since(start).Seconds())
}

// takes on average 1.85 seconds after at least 5 subsequent runs
func echoSlow() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep += " "
	}
	fmt.Println(s)
}

// takes on average 1.80 seconds after at least 5 subsequent runs
func echoFast() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}