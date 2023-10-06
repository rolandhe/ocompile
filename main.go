package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/rolandhe/ocompile/parser"
	"github.com/rolandhe/ocompile/transfer"
)

func main() {
	input, _ := antlr.NewFileStream("simple.service")
	lexer := parser.NewServiceLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewServiceParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Document()

	def := &transfer.Definition{}

	antlr.ParseTreeWalkerDefault.Walk(transfer.NewWalkListener(def), tree)

	if def.Err != nil {
		fmt.Println(def.Err)
	}

	fmt.Println(1)
}
