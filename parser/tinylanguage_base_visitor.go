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

func (v *BaseTinyLanguageVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIdentifierFunctionCall(ctx *IdentifierFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitAssertFunctionCall(ctx *AssertFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitSizeFunctionCall(ctx *SizeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitElseIfStat(ctx *ElseIfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitElseStat(ctx *ElseStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIdList(ctx *IdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitBoolExpression(ctx *BoolExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitNumberExpression(ctx *NumberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitEqExpression(ctx *EqExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitStringExpression(ctx *StringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitCompExpression(ctx *CompExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitMultExpression(ctx *MultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitListExpression(ctx *ListExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLanguageVisitor) VisitIndexes(ctx *IndexesContext) interface{} {
	return v.VisitChildren(ctx)
}
