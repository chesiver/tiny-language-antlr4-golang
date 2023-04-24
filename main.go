//go:generate java -jar ./grammar/antlr-4.8-complete.jar -Xexact-output-dir -o parser -visitor -no-listener -Dlanguage=Go ./grammar/TinyLanguage.g4
package main

import (
	"os"
	"tiny-language-antlr4-llvm-ir/component"
	"tiny-language-antlr4-llvm-ir/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// contents, _ := os.ReadFile("./test.tl")
	contents, _ := os.ReadFile("./tl/while.tl")
	// fmt.Printf("!!!!: %v\n", string(contents))
	inputStream := antlr.NewInputStream(string(contents))
	lexer := parser.NewTinyLanguageLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTinyLanguageParser(tokens)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	tree := p.Prog()
	var visitor parser.TinyLanguageVisitor = component.NewEvalVisitor()
	visitor.Visit(tree)
}
