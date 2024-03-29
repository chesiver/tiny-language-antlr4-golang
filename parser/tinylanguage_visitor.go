// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // TinyLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by TinyLanguageParser.
type TinyLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TinyLanguageParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#identifierFunctionCall.
	VisitIdentifierFunctionCall(ctx *IdentifierFunctionCallContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#printlnFunctionCall.
	VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#assertFunctionCall.
	VisitAssertFunctionCall(ctx *AssertFunctionCallContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#sizeFunctionCall.
	VisitSizeFunctionCall(ctx *SizeFunctionCallContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#elseIfStat.
	VisitElseIfStat(ctx *ElseIfStatContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#elseStat.
	VisitElseStat(ctx *ElseStatContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#boolExpression.
	VisitBoolExpression(ctx *BoolExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#numberExpression.
	VisitNumberExpression(ctx *NumberExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#unaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#powerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#eqExpression.
	VisitEqExpression(ctx *EqExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#inExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#stringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#expressionExpression.
	VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#compExpression.
	VisitCompExpression(ctx *CompExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#multExpression.
	VisitMultExpression(ctx *MultExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#listExpression.
	VisitListExpression(ctx *ListExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#ternaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#indexes.
	VisitIndexes(ctx *IndexesContext) interface{}
}
