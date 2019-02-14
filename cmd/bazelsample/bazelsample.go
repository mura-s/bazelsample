package main

import (
	"fmt"
	"os"

	"github.com/mura-s/bazelsample"
)

func main() {
	str, err := bazelsample.Sample(true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed")
		os.Exit(1)
	}
	fmt.Println(str)
}
