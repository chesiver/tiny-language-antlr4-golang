// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TinyLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTinyLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTinyLanguageVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitNumberExpression(ctx *NumberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitMultExpression(ctx *MultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
