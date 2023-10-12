package transfer

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/rolandhe/ocompile/parser"
	"log"
)

func ParseIdl(fileName string) (*Definition, error) {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		return nil, err
	}

	lexer := parser.NewServiceLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewServiceParser(stream)
	def := &Definition{}
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(newAcceptErrorListener(def))

	tree := p.Document()

	if def.err != nil {
		log.Printf("%v", def.err)
		return nil, err
	}

	antlr.ParseTreeWalkerDefault.Walk(newWalkListener(def), tree)
	if def.err != nil {
		log.Printf("%v", def.err)
		return nil, def.err
	}
	return def, nil
}
