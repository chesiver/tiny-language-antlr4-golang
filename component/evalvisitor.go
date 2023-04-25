package component

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"tiny-language-antlr4-llvm-ir/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

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
	return VOID
}

func (e *EvalVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	fmt.Printf("Enter - Block\n")
	scope := Scope{e.scope, make(map[string]*TLValue), false}
	for _, fctx := range ctx.AllFunctionDecl() {
		e.Visit(fctx)
	}
	for _, sctx := range ctx.AllStatement() {
		e.Visit(sctx)
	}
	if ectx := ctx.Expression(); ectx != nil {
		val := e.Visit(ectx).(*TLValue)
		e.scope = scope.parent
		return val
	}
	e.scope = scope.parent
	return VOID
}

func (e *EvalVisitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	fmt.Printf("Enter - For Statement\n")
	start := e.Visit(ctx.Expression(0)).(*TLValue).asInt()
	end := e.Visit(ctx.Expression(1)).(*TLValue).asInt()
	for i := start; i < end; i += 1 {
		e.scope.Assign(ctx.Identifier().GetText(), &TLValue{i})
		r := e.Visit(ctx.Block()).(*TLValue)
		if r != VOID {
			return r
		}
	}
	return VOID
}

func (e *EvalVisitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	fmt.Printf("Enter - While Statement\n")
	for e.Visit(ctx.Expression()).(*TLValue).asBool() {
		r := e.Visit(ctx.Block())
		if r != VOID {
			return r
		}
	}
	return VOID
}

func (e *EvalVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	fmt.Printf("Enter IfStatement\n")
	if e.Visit(ctx.IfStat().(*parser.IfStatContext).Expression()).(*TLValue).asBool() {
		return e.Visit(ctx.IfStat().(*parser.IfStatContext).Block())
	}
	for _, elseIfStat := range ctx.AllElseIfStat() {
		if e.Visit(elseIfStat.(*parser.ElseIfStatContext).Expression()).(*TLValue).asBool() {
			return e.Visit(elseIfStat.(*parser.ElseIfStatContext).Block())
		}
	}
	if elseContext := ctx.ElseStat().(*parser.ElseStatContext); elseContext != nil {
		return e.Visit(elseContext.Block())
	}
	return VOID
}

func (e *EvalVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	fmt.Printf("Enter VisitFunctionDecl\n")
	var params []antlr.TerminalNode
	if idList := ctx.IdList().(*parser.IdListContext); idList != nil {
		params = idList.AllIdentifier()
	}
	block := ctx.Block()
	id := ctx.Identifier().GetText() + "-" + strconv.Itoa(len(params))
	e.functions[id] = &Function{e.scope, params, block}
	return VOID
}

func (e *EvalVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Printf("Enter VisitStatement\n")
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) setAtIndex(ctx antlr.ParserRuleContext, indexes []parser.IExpressionContext, val *TLValue, newVal *TLValue) *TLValue {
	if !val.isList() {
		return INVALID
	}
	for i := 0; i < len(indexes)-1; i += 1 {
		idx := e.Visit(indexes[i]).(*TLValue)
		if !idx.isNumber() {
			return INVALID
		}
		val = val.asList()[idx.asInt()]
	}
	idx := e.Visit(indexes[len(indexes)-1]).(*TLValue)
	if !idx.isNumber() {
		return INVALID
	}
	val.asList()[idx.asInt()] = newVal
	return val
}

func (e *EvalVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Printf("Enter VisitAssignment\n")
	newVal := e.Visit(ctx.Expression()).(*TLValue)
	if ctx.Indexes() != nil {
		val := e.scope.resolve(ctx.Identifier().GetText())
		exps := ctx.Indexes().(*parser.IndexesContext).AllExpression()
		e.setAtIndex(ctx, exps, val, newVal)
	} else {
		name := ctx.Identifier().GetText()
		e.scope.Assign(name, newVal)
		fmt.Printf("In VisitAssignment - name: %v value: %v\n", name, newVal.value)
	}
	return VOID
}

func (e *EvalVisitor) VisitIdentifierFunctionCall(ctx *parser.IdentifierFunctionCallContext) interface{} {
	fmt.Printf("Enter VisitIdentifierFunctionCall\n")
	var params []parser.IExpressionContext
	if exprList := ctx.ExprList(); exprList != nil {
		params = exprList.(*parser.ExprListContext).AllExpression()
	}
	id := ctx.Identifier().GetText() + "-" + strconv.Itoa(len(params))
	if function := e.functions[id]; function != nil {
		var args []*TLValue
		for _, param := range params {
			args = append(args, e.Visit(param).(*TLValue))
		}
		return function.invoke(args, e.functions)
	}
	return VOID
}

func (e *EvalVisitor) VisitPrintlnFunctionCall(ctx *parser.PrintlnFunctionCallContext) interface{} {
	fmt.Printf("Enter VisitPrintlnFunctionCall\n")
	if expr := ctx.Expression(); expr != nil {
		fmt.Printf("%s\n", e.Visit(expr).(*TLValue).String())
	} else {
		fmt.Println()
	}
	return VOID
}

func (e *EvalVisitor) VisitAssertFunctionCall(ctx *parser.AssertFunctionCallContext) interface{} {
	fmt.Printf("Enter Visit Assert Function\n")
	val := e.Visit(ctx.Expression()).(*TLValue)
	if !val.isBool() || !val.asBool() {
		return INVALID
	}
	return VOID
}

func (e *EvalVisitor) VisitSizeFunctionCall(ctx *parser.SizeFunctionCallContext) interface{} {
	fmt.Printf("Enter Visit Size Function\n")
	val := e.Visit(ctx.Expression()).(*TLValue)
	if val.isString() {
		return &TLValue{len(val.asString())}
	}
	if val.isList() {
		return &TLValue{len(val.asList())}
	}
	return INVALID
}

func (e *EvalVisitor) VisitExpressionExpression(ctx *parser.ExpressionExpressionContext) interface{} {
	fmt.Printf("Enter - Expression Expression\n")
	val := e.Visit(ctx.Expression())
	return val.(*TLValue)
}

func (e *EvalVisitor) VisitUnaryMinusExpression(ctx *parser.UnaryMinusExpressionContext) interface{} {
	val := e.Visit(ctx.Expression()).(*TLValue)
	if val.isInt() {
		return &TLValue{value: -val.asInt()}
	}
	if val.isDouble() {
		return &TLValue{value: -val.asDouble()}
	}
	return NULL
}

func (e *EvalVisitor) lt(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() < right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &TLValue{left.asString() < right.asString()}
	}
	return false
}

func (e *EvalVisitor) ltEq(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() <= right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &TLValue{left.asString() <= right.asString()}
	}
	return false
}

func (e *EvalVisitor) gt(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() > right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &TLValue{left.asString() > right.asString()}
	}
	return false
}

func (e *EvalVisitor) gtEq(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() >= right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &TLValue{left.asString() >= right.asString()}
	}
	return false
}

func (e *EvalVisitor) add(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() + right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &TLValue{left.asString() + right.asString()}
	}
	return NULL
}

func (e *EvalVisitor) sub(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() - right.asDouble()}
}

func (e *EvalVisitor) mult(left *TLValue, right *TLValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &TLValue{left.asDouble() * right.asDouble()}
	}
	if left.isString() && right.isNumber() {
		times := right.asInt()
		sb := strings.Builder{}
		for i := 0; i < times; i++ {
			sb.WriteString(left.asString())
		}
		return &TLValue{sb.String()}
	}
	return NULL
}

func (e *EvalVisitor) div(left *TLValue, right *TLValue) interface{} {
	return &TLValue{left.asDouble() / right.asDouble()}
}

func (e *EvalVisitor) VisitEqExpression(ctx *parser.EqExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.TinyLanguageLexerEquals:
		return &TLValue{value: left.equals(right)}
	case parser.TinyLanguageLexerNEquals:
		return &TLValue{value: !left.equals(right)}
	}
	return NULL
}

func (e *EvalVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	return &TLValue{left.asBool() && right.asBool()}
}

func (e *EvalVisitor) VisitOrExpression(ctx *parser.OrExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	return &TLValue{left.asBool() || right.asBool()}
}

func (e *EvalVisitor) VisitTernaryExpression(ctx *parser.TernaryExpressionContext) interface{} {
	condition := e.Visit(ctx.Expression(0)).(*TLValue)
	if condition.asBool() {
		return e.Visit(ctx.Expression(1)).(*TLValue)
	} else {
		return e.Visit(ctx.Expression(2)).(*TLValue)
	}
}

func (e *EvalVisitor) VisitInExpression(ctx *parser.InExpressionContext) interface{} {
	lhs := e.Visit(ctx.Expression(0)).(*TLValue)
	rhs := e.Visit(ctx.Expression(1)).(*TLValue)
	if rhs.isList() {
		for _, c := range rhs.asList() {
			if c.equals(lhs) {
				return &TLValue{true}
			}
		}
		return &TLValue{false}
	}
	return INVALID
}

func (e *EvalVisitor) VisitCompExpression(ctx *parser.CompExpressionContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*TLValue)
	right := e.Visit(ctx.Expression(1)).(*TLValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.TinyLanguageLexerLT:
		return e.lt(left, right)
	case parser.TinyLanguageLexerLTEquals:
		return e.ltEq(left, right)
	case parser.TinyLanguageLexerGT:
		return e.gt(left, right)
	case parser.TinyLanguageLexerGTEquals:
		return e.gtEq(left, right)
	default:
		return NULL
	}
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
		return NULL
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
		return NULL
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

func (e *EvalVisitor) VisitBoolExpression(ctx *parser.BoolExpressionContext) interface{} {
	val := ctx.GetText() == "true"
	return &TLValue{val}
}

func (e *EvalVisitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	id := ctx.Identifier().GetText()
	val := e.scope.resolve(id)
	if ctx.Indexes() != nil {
		exprs := ctx.Indexes().(*parser.IndexesContext).AllExpression()
		val = e.resolveIndexes(val, exprs).(*TLValue)
	}
	return val
}

func (e *EvalVisitor) VisitStringExpression(ctx *parser.StringExpressionContext) interface{} {
	text := ctx.GetText()
	text = text[1 : len(text)-1]
	return &TLValue{text}
}

func (e *EvalVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	fmt.Printf("Enter VisitFunctionCallExpression\n")
	val := e.Visit(ctx.FunctionCall()).(*TLValue)
	return val
}

func (e *EvalVisitor) resolveIndexes(val *TLValue, indexes []parser.IExpressionContext) interface{} {
	for _, IExprCtx := range indexes {
		idx := e.Visit(IExprCtx).(*TLValue).asInt()
		if val.isString() {
			val = &TLValue{val.asString()[idx : idx+1]}
		} else {
			val = val.asList()[idx]
		}
	}
	return val
}

func (e *EvalVisitor) VisitList(ctx *parser.ListContext) interface{} {
	fmt.Printf("Enter - List\n")
	arr := make([]*TLValue, 0)
	if ctx.ExprList() != nil {
		for _, expr := range ctx.ExprList().(*parser.ExprListContext).AllExpression() {
			arr = append(arr, e.Visit(expr).(*TLValue))
		}
	}
	return &TLValue{arr}
}

func (e *EvalVisitor) VisitListExpression(ctx *parser.ListExpressionContext) interface{} {
	fmt.Printf("Enter - List Expression\n")
	val := e.Visit(ctx.List()).(*TLValue)
	if ctx.Indexes() != nil {
		exps := make([]parser.IExpressionContext, 0)
		val = e.resolveIndexes(val, exps).(*TLValue)
	}
	return val
}
