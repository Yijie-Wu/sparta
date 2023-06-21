package main

import (
	"fmt"
	"sparta/cmd"
)

var banner = `
banner
`

func init() {
	fmt.Println(banner)
}

// @version 1.0.0
// @title Open Source Software Manager
// @description An Open Source Software Apply and Manager System
func main() {
	defer cmd.StopApplication()
	cmd.StartApplication()
}
