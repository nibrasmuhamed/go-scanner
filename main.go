/*
Copyright © 2022 Muhammed Nibras nibrasmn027@gmail.com

*/
package main

import (
	"fmt"

	cmd "github.com/nibrasmuhamed/go-scanner/cmd"
)

func main() {
	fmt.Println("\u001b[34m", welcome, "\u001b[0m")
	cmd.Execute()
}
