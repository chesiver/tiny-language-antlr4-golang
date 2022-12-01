// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TinyLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr"

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

	// Visit a parse tree produced by TinyLanguageParser#printlnFunctionCall.
	VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#expressionExpression.
	VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#numberExpression.
	VisitNumberExpression(ctx *NumberExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by TinyLanguageParser#multExpression.
	VisitMultExpression(ctx *MultExpressionContext) interface{}
}
