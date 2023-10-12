package main

import (
	"fmt"
	"github.com/rolandhe/ocompile/transfer"
)

func main() {
	def, err := transfer.ParseIdl("simple.service")
	fmt.Println(def, err)
	fmt.Println(1)
}
