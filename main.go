// Package main is the entry point
package main

import (
	"fmt"

	"github.com/PixelBrewer/normalize/internal/config"
)

func main() {
	fmt.Printf("Welcome to Normalize! \n")
	fmt.Printf("Just enter a command in the following format:")
	fmt.Printf("normalize -os-flag file-type")
	config.FindConfigToml()
}
