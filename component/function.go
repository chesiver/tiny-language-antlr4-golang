package component

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Function struct {
	parentScope *Scope
	params      []antlr.TerminalNode
	block       antlr.ParseTree
}

func (f *Function) invoke(args []*TLValue, functions map[string]*Function) *TLValue {
	if len(args) != len(f.params) {
		fmt.Printf("Illegal Function call\n")
		return nil
	}
	scopeNext := Scope{
		parent:     f.parentScope,
		variables:  make(map[string]*TLValue, 0),
		isFunction: true,
	}
	for i := 0; i < len(f.params); i++ {
		scopeNext.assignParam(f.params[i].GetText(), args[i])
	}
	return nil
}
