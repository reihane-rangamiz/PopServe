package main

import "fmt"



func main() {
	conf, err := AppConfig()
	fmt.Print(conf)
	fmt.Print(err)
}