package component

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"tiny-language-antlr4-llvm-ir/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// var returnValue *ReturnValue = &ReturnValue{}

type EvalVisitor struct {
	parser.BaseTinyLanguageVisitor
	scope     *Scope
	functions map[string]*Function
}

func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{
		scope:     NewScope(),
		functions: make(map[string]*Function),
	}
}

func (e *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("visit input type: %v\n", reflect.TypeOf(tree))
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("Error: - %v\n", t.GetText())
		return nil
	default:
		return tree.Accept(e)
	}
}

func (e *EvalVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	fmt.Printf("Calculating Program: %s\n", ctx.GetText())
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		// fmt.Printf("> %s\n", n.(antlr.ParseTree).GetText())
		e.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (e *EvalVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	fmt.Printf("Enter VisitBlock\n")
	scope := Scope{e.scope, make(map[string]*TLValue), false}
	// if fctx := ctx.FunctionDecl(0); fctx != nil {
	// 	e.Visit(fctx)
	// }
	for _, sctx := range(ctx.AllStatement()) {
		e.Visit(sctx)
	}
	// if ectx := ctx.Expression(); ectx != nil {
	// 	returnValue.value = e.Visit(ectx).(*TLValue)
	// 	e.scope = scope.parent
	// 	return returnValue, true
	// }
	e.scope = scope.parent
	return nil
}

func (e *EvalVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Printf("Enter VisitStatement\n")
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Printf("Enter VisitAssignment\n")
	val := e.Visit(ctx.Expression()).(*TLValue)
	name := ctx.Identifier().GetText()
	e.scope.Assign(name, val)
	return nil
}

// func (e *EvalVisitor) VisitIdentifierFunctionCall(ctx *parser.IdentifierFunctionCallContext) *TLValue {
// 	var params []parser.IExpressionContext
// 	if exprList := ctx.ExprList(); exprList != nil {
// 		params = exprList.(*parser.ExprListContext).AllExpression()
// 	}
// 	id := ctx.Identifier().GetText() + strconv.Itoa(len(params))
// 	if function := e.functions[id]; function != nil {
// 		var args []*TLValue
// 		for _, param := range params {
// 			args = append(args, e.Visit(param).(*TLValue))
// 		}
// 		return function.invoke(args, e.functions)
// 	}
// 	return nil
// }

func (e *EvalVisitor) VisitPrintlnFunctionCall(ctx *parser.PrintlnFunctionCallContext) interface{} {
	fmt.Printf("Enter VisitPrintlnFunctionCall\n")
	if expr := ctx.Expression(); expr != nil {
		fmt.Printf("%v\n", e.Visit(expr).(*TLValue).value)
	} else {
		fmt.Println()
	}
	return nil
}

func (e *EvalVisitor) VisitExpressionExpression(ctx *parser.ExpressionExpressionContext) interface{} {
	val := e.Visit(ctx.Expression())
	return val.(*TLValue)
}

func (e *EvalVisitor) add(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() + right.asDouble()}
}

func (e *EvalVisitor) sub(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() - right.asDouble()}
}

func (e *EvalVisitor) mult(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() * right.asDouble()}
}

func (e *EvalVisitor) div(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() / right.asDouble()}
}

func (e *EvalVisitor) VisitAddExpression(ctx *parser.AddExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.TinyLanguageLexerAdd:
		return e.add(left, right)
	case parser.TinyLanguageLexerSubtract:
		return e.sub(left, right)
	default:
		return nil
	}
}

func (e *EvalVisitor) VisitMultExpression(ctx *parser.MultExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.TinyLanguageLexerMult:
		return e.mult(left, right)
	case parser.TinyLanguageLexerMod:
		return e.div(left, right)
	default:
		return nil
	}
}

func (e *EvalVisitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	return &TLValue{math.Pow(left.asDouble(), right.asDouble())}
}

func (e *EvalVisitor) VisitNumberExpression(ctx *parser.NumberExpressionContext) interface{} {
	num, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return &TLValue{num}
}

func (e *EvalVisitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	id := ctx.Identifier().GetText()
	val := e.scope.resolve(id)
	return val
}

func (e *EvalVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	fmt.Printf("Enter VisitFunctionCallExpression\n")
	val := e.Visit(ctx.FunctionCall()).(*TLValue)
	return val
}
