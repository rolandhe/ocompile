package main

import (
	"fmt"
	"github.com/rolandhe/ocompile/transfer"
	"github.com/rolandhe/ocompile/transfer/render/client"
)

func main() {
	def, err := transfer.ParseIdl("simple.service")
	fmt.Println(def, err)
	err = client.RenderSvc("/Users/hexiufeng/github/ocompile/target", def)
	fmt.Println(err)
	fmt.Println(1)
}
