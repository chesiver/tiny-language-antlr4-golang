// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TinyLanguage

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 214,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 5, 5, 65, 10, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 74, 10, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 99, 10, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 5, 9, 105, 10, 9, 3, 9, 5, 9, 108, 10, 9, 3, 10, 3, 10,
	7, 10, 112, 10, 10, 12, 10, 14, 10, 115, 11, 10, 3, 10, 5, 10, 118, 10,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7,
	14, 140, 10, 14, 12, 14, 14, 14, 143, 11, 14, 3, 15, 3, 15, 3, 15, 7, 15,
	148, 10, 15, 12, 15, 14, 15, 151, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 5, 16, 161, 10, 16, 3, 16, 3, 16, 5, 16, 165,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 172, 10, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 195,
	10, 16, 12, 16, 14, 16, 198, 11, 16, 3, 17, 3, 17, 5, 17, 202, 10, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 6, 18, 210, 10, 18, 13, 18, 14,
	18, 211, 3, 18, 2, 3, 30, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 2, 6, 3, 2, 34, 36, 3, 2, 32, 33, 4, 2, 29, 30, 37,
	38, 3, 2, 27, 28, 2, 230, 2, 36, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 60,
	3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 79, 3, 2, 2, 2,
	14, 89, 3, 2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20, 121,
	3, 2, 2, 2, 22, 126, 3, 2, 2, 2, 24, 132, 3, 2, 2, 2, 26, 136, 3, 2, 2,
	2, 28, 144, 3, 2, 2, 2, 30, 171, 3, 2, 2, 2, 32, 199, 3, 2, 2, 2, 34, 209,
	3, 2, 2, 2, 36, 37, 5, 4, 3, 2, 37, 3, 3, 2, 2, 2, 38, 41, 5, 6, 4, 2,
	39, 41, 5, 10, 6, 2, 40, 38, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3,
	2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 49, 3, 2, 2, 2, 44,
	42, 3, 2, 2, 2, 45, 46, 7, 15, 2, 2, 46, 47, 5, 30, 16, 2, 47, 48, 7, 3,
	2, 2, 48, 50, 3, 2, 2, 2, 49, 45, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 5,
	3, 2, 2, 2, 51, 52, 5, 8, 5, 2, 52, 53, 7, 3, 2, 2, 53, 61, 3, 2, 2, 2,
	54, 55, 5, 16, 9, 2, 55, 56, 7, 3, 2, 2, 56, 61, 3, 2, 2, 2, 57, 61, 5,
	18, 10, 2, 58, 61, 5, 12, 7, 2, 59, 61, 5, 14, 8, 2, 60, 51, 3, 2, 2, 2,
	60, 54, 3, 2, 2, 2, 60, 57, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3,
	2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 64, 7, 22, 2, 2, 63, 65, 5, 34, 18, 2,
	64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 7,
	4, 2, 2, 67, 68, 5, 30, 16, 2, 68, 9, 3, 2, 2, 2, 69, 70, 7, 11, 2, 2,
	70, 71, 7, 22, 2, 2, 71, 73, 7, 5, 2, 2, 72, 74, 5, 26, 14, 2, 73, 72,
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 7, 6, 2, 2,
	76, 77, 5, 4, 3, 2, 77, 78, 7, 19, 2, 2, 78, 11, 3, 2, 2, 2, 79, 80, 7,
	16, 2, 2, 80, 81, 7, 22, 2, 2, 81, 82, 7, 4, 2, 2, 82, 83, 5, 30, 16, 2,
	83, 84, 7, 18, 2, 2, 84, 85, 5, 30, 16, 2, 85, 86, 7, 14, 2, 2, 86, 87,
	5, 4, 3, 2, 87, 88, 7, 19, 2, 2, 88, 13, 3, 2, 2, 2, 89, 90, 7, 17, 2,
	2, 90, 91, 5, 30, 16, 2, 91, 92, 7, 14, 2, 2, 92, 93, 5, 4, 3, 2, 93, 94,
	7, 19, 2, 2, 94, 15, 3, 2, 2, 2, 95, 96, 7, 22, 2, 2, 96, 98, 7, 5, 2,
	2, 97, 99, 5, 28, 15, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100,
	3, 2, 2, 2, 100, 108, 7, 6, 2, 2, 101, 102, 7, 10, 2, 2, 102, 104, 7, 5,
	2, 2, 103, 105, 5, 30, 16, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2,
	2, 105, 106, 3, 2, 2, 2, 106, 108, 7, 6, 2, 2, 107, 95, 3, 2, 2, 2, 107,
	101, 3, 2, 2, 2, 108, 17, 3, 2, 2, 2, 109, 113, 5, 20, 11, 2, 110, 112,
	5, 22, 12, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3,
	2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2,
	2, 116, 118, 5, 24, 13, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 120, 7, 19, 2, 2, 120, 19, 3, 2, 2, 2, 121,
	122, 7, 12, 2, 2, 122, 123, 5, 30, 16, 2, 123, 124, 7, 14, 2, 2, 124, 125,
	5, 4, 3, 2, 125, 21, 3, 2, 2, 2, 126, 127, 7, 13, 2, 2, 127, 128, 7, 12,
	2, 2, 128, 129, 5, 30, 16, 2, 129, 130, 7, 14, 2, 2, 130, 131, 5, 4, 3,
	2, 131, 23, 3, 2, 2, 2, 132, 133, 7, 13, 2, 2, 133, 134, 7, 14, 2, 2, 134,
	135, 5, 4, 3, 2, 135, 25, 3, 2, 2, 2, 136, 141, 7, 22, 2, 2, 137, 138,
	7, 7, 2, 2, 138, 140, 7, 22, 2, 2, 139, 137, 3, 2, 2, 2, 140, 143, 3, 2,
	2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 27, 3, 2, 2, 2,
	143, 141, 3, 2, 2, 2, 144, 149, 5, 30, 16, 2, 145, 146, 7, 7, 2, 2, 146,
	148, 5, 30, 16, 2, 147, 145, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147,
	3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 29, 3, 2, 2, 2, 151, 149, 3, 2,
	2, 2, 152, 153, 8, 16, 1, 2, 153, 154, 7, 33, 2, 2, 154, 172, 5, 30, 16,
	17, 155, 172, 7, 21, 2, 2, 156, 172, 7, 20, 2, 2, 157, 172, 5, 16, 9, 2,
	158, 160, 5, 32, 17, 2, 159, 161, 5, 34, 18, 2, 160, 159, 3, 2, 2, 2, 160,
	161, 3, 2, 2, 2, 161, 172, 3, 2, 2, 2, 162, 164, 7, 22, 2, 2, 163, 165,
	5, 34, 18, 2, 164, 163, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 172, 3,
	2, 2, 2, 166, 172, 7, 23, 2, 2, 167, 168, 7, 5, 2, 2, 168, 169, 5, 30,
	16, 2, 169, 170, 7, 6, 2, 2, 170, 172, 3, 2, 2, 2, 171, 152, 3, 2, 2, 2,
	171, 155, 3, 2, 2, 2, 171, 156, 3, 2, 2, 2, 171, 157, 3, 2, 2, 2, 171,
	158, 3, 2, 2, 2, 171, 162, 3, 2, 2, 2, 171, 166, 3, 2, 2, 2, 171, 167,
	3, 2, 2, 2, 172, 196, 3, 2, 2, 2, 173, 174, 12, 16, 2, 2, 174, 175, 7,
	31, 2, 2, 175, 195, 5, 30, 16, 16, 176, 177, 12, 15, 2, 2, 177, 178, 9,
	2, 2, 2, 178, 195, 5, 30, 16, 16, 179, 180, 12, 14, 2, 2, 180, 181, 9,
	3, 2, 2, 181, 195, 5, 30, 16, 15, 182, 183, 12, 13, 2, 2, 183, 184, 9,
	4, 2, 2, 184, 195, 5, 30, 16, 14, 185, 186, 12, 12, 2, 2, 186, 187, 9,
	5, 2, 2, 187, 195, 5, 30, 16, 13, 188, 189, 12, 11, 2, 2, 189, 190, 7,
	26, 2, 2, 190, 195, 5, 30, 16, 12, 191, 192, 12, 10, 2, 2, 192, 193, 7,
	25, 2, 2, 193, 195, 5, 30, 16, 11, 194, 173, 3, 2, 2, 2, 194, 176, 3, 2,
	2, 2, 194, 179, 3, 2, 2, 2, 194, 182, 3, 2, 2, 2, 194, 185, 3, 2, 2, 2,
	194, 188, 3, 2, 2, 2, 194, 191, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196,
	194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 31, 3, 2, 2, 2, 198, 196, 3,
	2, 2, 2, 199, 201, 7, 8, 2, 2, 200, 202, 5, 28, 15, 2, 201, 200, 3, 2,
	2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 7, 9, 2, 2,
	204, 33, 3, 2, 2, 2, 205, 206, 7, 8, 2, 2, 206, 207, 5, 30, 16, 2, 207,
	208, 7, 9, 2, 2, 208, 210, 3, 2, 2, 2, 209, 205, 3, 2, 2, 2, 210, 211,
	3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 35, 3, 2,
	2, 2, 22, 40, 42, 49, 60, 64, 73, 98, 104, 107, 113, 117, 141, 149, 160,
	164, 171, 194, 196, 201, 211,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'='", "'('", "')'", "','", "'['", "']'", "'println'", "'def'",
	"'if'", "'else'", "'do'", "'return'", "'for'", "'while'", "'to'", "'end'",
	"", "", "", "", "", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'^'",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "Println", "Def", "If", "Else", "Do", "Return",
	"For", "While", "To", "End", "Bool", "Number", "Identifier", "StringLiteral",
	"Space", "Or", "And", "Equals", "NEquals", "GTEquals", "LTEquals", "Power",
	"Add", "Subtract", "Mult", "Divide", "Mod", "GT", "LT",
}

var ruleNames = []string{
	"prog", "block", "statement", "assignment", "functionDecl", "forStatement",
	"whileStatement", "functionCall", "ifStatement", "ifStat", "elseIfStat",
	"elseStat", "idList", "exprList", "expression", "list", "indexes",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TinyLanguageParser struct {
	*antlr.BaseParser
}

func NewTinyLanguageParser(input antlr.TokenStream) *TinyLanguageParser {
	this := new(TinyLanguageParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TinyLanguage.g4"

	return this
}

// TinyLanguageParser tokens.
const (
	TinyLanguageParserEOF           = antlr.TokenEOF
	TinyLanguageParserT__0          = 1
	TinyLanguageParserT__1          = 2
	TinyLanguageParserT__2          = 3
	TinyLanguageParserT__3          = 4
	TinyLanguageParserT__4          = 5
	TinyLanguageParserT__5          = 6
	TinyLanguageParserT__6          = 7
	TinyLanguageParserPrintln       = 8
	TinyLanguageParserDef           = 9
	TinyLanguageParserIf            = 10
	TinyLanguageParserElse          = 11
	TinyLanguageParserDo            = 12
	TinyLanguageParserReturn        = 13
	TinyLanguageParserFor           = 14
	TinyLanguageParserWhile         = 15
	TinyLanguageParserTo            = 16
	TinyLanguageParserEnd           = 17
	TinyLanguageParserBool          = 18
	TinyLanguageParserNumber        = 19
	TinyLanguageParserIdentifier    = 20
	TinyLanguageParserStringLiteral = 21
	TinyLanguageParserSpace         = 22
	TinyLanguageParserOr            = 23
	TinyLanguageParserAnd           = 24
	TinyLanguageParserEquals        = 25
	TinyLanguageParserNEquals       = 26
	TinyLanguageParserGTEquals      = 27
	TinyLanguageParserLTEquals      = 28
	TinyLanguageParserPower         = 29
	TinyLanguageParserAdd           = 30
	TinyLanguageParserSubtract      = 31
	TinyLanguageParserMult          = 32
	TinyLanguageParserDivide        = 33
	TinyLanguageParserMod           = 34
	TinyLanguageParserGT            = 35
	TinyLanguageParserLT            = 36
)

// TinyLanguageParser rules.
const (
	TinyLanguageParserRULE_prog           = 0
	TinyLanguageParserRULE_block          = 1
	TinyLanguageParserRULE_statement      = 2
	TinyLanguageParserRULE_assignment     = 3
	TinyLanguageParserRULE_functionDecl   = 4
	TinyLanguageParserRULE_forStatement   = 5
	TinyLanguageParserRULE_whileStatement = 6
	TinyLanguageParserRULE_functionCall   = 7
	TinyLanguageParserRULE_ifStatement    = 8
	TinyLanguageParserRULE_ifStat         = 9
	TinyLanguageParserRULE_elseIfStat     = 10
	TinyLanguageParserRULE_elseStat       = 11
	TinyLanguageParserRULE_idList         = 12
	TinyLanguageParserRULE_exprList       = 13
	TinyLanguageParserRULE_expression     = 14
	TinyLanguageParserRULE_list           = 15
	TinyLanguageParserRULE_indexes        = 16
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TinyLanguageParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Block()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllFunctionDecl() []IFunctionDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclContext)(nil)).Elem())
	var tst = make([]IFunctionDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclContext)
		}
	}

	return tst
}

func (s *BlockContext) FunctionDecl(i int) IFunctionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *BlockContext) Return() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserReturn, 0)
}

func (s *BlockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TinyLanguageParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLanguageParserPrintln)|(1<<TinyLanguageParserDef)|(1<<TinyLanguageParserIf)|(1<<TinyLanguageParserFor)|(1<<TinyLanguageParserWhile)|(1<<TinyLanguageParserIdentifier))) != 0 {
		p.SetState(38)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TinyLanguageParserPrintln, TinyLanguageParserIf, TinyLanguageParserFor, TinyLanguageParserWhile, TinyLanguageParserIdentifier:
			{
				p.SetState(36)
				p.Statement()
			}

		case TinyLanguageParserDef:
			{
				p.SetState(37)
				p.FunctionDecl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLanguageParserReturn {
		{
			p.SetState(43)
			p.Match(TinyLanguageParserReturn)
		}
		{
			p.SetState(44)
			p.expression(0)
		}
		{
			p.SetState(45)
			p.Match(TinyLanguageParserT__0)
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TinyLanguageParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Assignment()
		}
		{
			p.SetState(50)
			p.Match(TinyLanguageParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.FunctionCall()
		}
		{
			p.SetState(53)
			p.Match(TinyLanguageParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.ForStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.WhileStatement()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TinyLanguageParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(TinyLanguageParserIdentifier)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLanguageParserT__5 {
		{
			p.SetState(61)
			p.Indexes()
		}

	}
	{
		p.SetState(64)
		p.Match(TinyLanguageParserT__1)
	}
	{
		p.SetState(65)
		p.expression(0)
	}

	return localctx
}

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_functionDecl
	return p
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) Def() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDef, 0)
}

func (s *FunctionDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, 0)
}

func (s *FunctionDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *FunctionDeclContext) IdList() IIdListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TinyLanguageParserRULE_functionDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(TinyLanguageParserDef)
	}
	{
		p.SetState(68)
		p.Match(TinyLanguageParserIdentifier)
	}
	{
		p.SetState(69)
		p.Match(TinyLanguageParserT__2)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLanguageParserIdentifier {
		{
			p.SetState(70)
			p.IdList()
		}

	}
	{
		p.SetState(73)
		p.Match(TinyLanguageParserT__3)
	}
	{
		p.SetState(74)
		p.Block()
	}
	{
		p.SetState(75)
		p.Match(TinyLanguageParserEnd)
	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserFor, 0)
}

func (s *ForStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, 0)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) To() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserTo, 0)
}

func (s *ForStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TinyLanguageParserRULE_forStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(TinyLanguageParserFor)
	}
	{
		p.SetState(78)
		p.Match(TinyLanguageParserIdentifier)
	}
	{
		p.SetState(79)
		p.Match(TinyLanguageParserT__1)
	}
	{
		p.SetState(80)
		p.expression(0)
	}
	{
		p.SetState(81)
		p.Match(TinyLanguageParserTo)
	}
	{
		p.SetState(82)
		p.expression(0)
	}
	{
		p.SetState(83)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(84)
		p.Block()
	}
	{
		p.SetState(85)
		p.Match(TinyLanguageParserEnd)
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TinyLanguageParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(TinyLanguageParserWhile)
	}
	{
		p.SetState(88)
		p.expression(0)
	}
	{
		p.SetState(89)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(90)
		p.Block()
	}
	{
		p.SetState(91)
		p.Match(TinyLanguageParserEnd)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) CopyFrom(ctx *FunctionCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintlnFunctionCallContext struct {
	*FunctionCallContext
}

func NewPrintlnFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnFunctionCallContext {
	var p = new(PrintlnFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *PrintlnFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnFunctionCallContext) Println() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserPrintln, 0)
}

func (s *PrintlnFunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintlnFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitPrintlnFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierFunctionCallContext struct {
	*FunctionCallContext
}

func NewIdentifierFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierFunctionCallContext {
	var p = new(IdentifierFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *IdentifierFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, 0)
}

func (s *IdentifierFunctionCallContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *IdentifierFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIdentifierFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TinyLanguageParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLanguageParserIdentifier:
		localctx = NewIdentifierFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(TinyLanguageParserIdentifier)
		}
		{
			p.SetState(94)
			p.Match(TinyLanguageParserT__2)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLanguageParserT__2)|(1<<TinyLanguageParserT__5)|(1<<TinyLanguageParserPrintln)|(1<<TinyLanguageParserBool)|(1<<TinyLanguageParserNumber)|(1<<TinyLanguageParserIdentifier)|(1<<TinyLanguageParserStringLiteral)|(1<<TinyLanguageParserSubtract))) != 0 {
			{
				p.SetState(95)
				p.ExprList()
			}

		}
		{
			p.SetState(98)
			p.Match(TinyLanguageParserT__3)
		}

	case TinyLanguageParserPrintln:
		localctx = NewPrintlnFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(TinyLanguageParserPrintln)
		}
		{
			p.SetState(100)
			p.Match(TinyLanguageParserT__2)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLanguageParserT__2)|(1<<TinyLanguageParserT__5)|(1<<TinyLanguageParserPrintln)|(1<<TinyLanguageParserBool)|(1<<TinyLanguageParserNumber)|(1<<TinyLanguageParserIdentifier)|(1<<TinyLanguageParserStringLiteral)|(1<<TinyLanguageParserSubtract))) != 0 {
			{
				p.SetState(101)
				p.expression(0)
			}

		}
		{
			p.SetState(104)
			p.Match(TinyLanguageParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfStat() IIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *IfStatementContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *IfStatementContext) AllElseIfStat() []IElseIfStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStatContext)(nil)).Elem())
	var tst = make([]IElseIfStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStatContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStat(i int) IElseIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStatContext)
}

func (s *IfStatementContext) ElseStat() IElseStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TinyLanguageParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.IfStat()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(108)
				p.ElseIfStat()
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLanguageParserElse {
		{
			p.SetState(114)
			p.ElseStat()
		}

	}
	{
		p.SetState(117)
		p.Match(TinyLanguageParserEnd)
	}

	return localctx
}

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_ifStat
	return p
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) If() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIf, 0)
}

func (s *IfStatContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *IfStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TinyLanguageParserRULE_ifStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(TinyLanguageParserIf)
	}
	{
		p.SetState(120)
		p.expression(0)
	}
	{
		p.SetState(121)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(122)
		p.Block()
	}

	return localctx
}

// IElseIfStatContext is an interface to support dynamic dispatch.
type IElseIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatContext differentiates from other interfaces.
	IsElseIfStatContext()
}

type ElseIfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatContext() *ElseIfStatContext {
	var p = new(ElseIfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_elseIfStat
	return p
}

func (*ElseIfStatContext) IsElseIfStatContext() {}

func NewElseIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatContext {
	var p = new(ElseIfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_elseIfStat

	return p
}

func (s *ElseIfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatContext) Else() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserElse, 0)
}

func (s *ElseIfStatContext) If() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIf, 0)
}

func (s *ElseIfStatContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStatContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *ElseIfStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitElseIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) ElseIfStat() (localctx IElseIfStatContext) {
	localctx = NewElseIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TinyLanguageParserRULE_elseIfStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(TinyLanguageParserElse)
	}
	{
		p.SetState(125)
		p.Match(TinyLanguageParserIf)
	}
	{
		p.SetState(126)
		p.expression(0)
	}
	{
		p.SetState(127)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(128)
		p.Block()
	}

	return localctx
}

// IElseStatContext is an interface to support dynamic dispatch.
type IElseStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatContext differentiates from other interfaces.
	IsElseStatContext()
}

type ElseStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatContext() *ElseStatContext {
	var p = new(ElseStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_elseStat
	return p
}

func (*ElseStatContext) IsElseStatContext() {}

func NewElseStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatContext {
	var p = new(ElseStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_elseStat

	return p
}

func (s *ElseStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatContext) Else() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserElse, 0)
}

func (s *ElseStatContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *ElseStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitElseStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) ElseStat() (localctx IElseStatContext) {
	localctx = NewElseStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TinyLanguageParserRULE_elseStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TinyLanguageParserElse)
	}
	{
		p.SetState(131)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(132)
		p.Block()
	}

	return localctx
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_idList
	return p
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(TinyLanguageParserIdentifier)
}

func (s *IdListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TinyLanguageParserRULE_idList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(TinyLanguageParserIdentifier)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLanguageParserT__4 {
		{
			p.SetState(135)
			p.Match(TinyLanguageParserT__4)
		}
		{
			p.SetState(136)
			p.Match(TinyLanguageParserIdentifier)
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TinyLanguageParserRULE_exprList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.expression(0)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLanguageParserT__4 {
		{
			p.SetState(143)
			p.Match(TinyLanguageParserT__4)
		}
		{
			p.SetState(144)
			p.expression(0)
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExpressionContext struct {
	*ExpressionContext
}

func NewBoolExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) Bool() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserBool, 0)
}

func (s *BoolExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitBoolExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExpressionContext struct {
	*ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserNumber, 0)
}

func (s *NumberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitNumberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpressionContext struct {
	*ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserOr, 0)
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExpressionContext struct {
	*ExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Subtract() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserSubtract, 0)
}

func (s *UnaryMinusExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitUnaryMinusExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerExpressionContext struct {
	*ExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowerExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExpressionContext) Power() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserPower, 0)
}

func (s *PowerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitPowerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewEqExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExpressionContext {
	var p = new(EqExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEquals, 0)
}

func (s *EqExpressionContext) NEquals() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserNEquals, 0)
}

func (s *EqExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitEqExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpressionContext struct {
	*ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserAnd, 0)
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExpressionContext struct {
	*ExpressionContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserStringLiteral, 0)
}

func (s *StringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionExpressionContext struct {
	*ExpressionContext
}

func NewExpressionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionExpressionContext {
	var p = new(ExpressionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitExpressionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) Add() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserAdd, 0)
}

func (s *AddExpressionContext) Subtract() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserSubtract, 0)
}

func (s *AddExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitAddExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewCompExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExpressionContext {
	var p = new(CompExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompExpressionContext) GetOp() antlr.Token { return s.op }

func (s *CompExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) GTEquals() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserGTEquals, 0)
}

func (s *CompExpressionContext) LTEquals() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserLTEquals, 0)
}

func (s *CompExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserGT, 0)
}

func (s *CompExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserLT, 0)
}

func (s *CompExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitCompExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMultExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpressionContext {
	var p = new(MultExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) Mult() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserMult, 0)
}

func (s *MultExpressionContext) Divide() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDivide, 0)
}

func (s *MultExpressionContext) Mod() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserMod, 0)
}

func (s *MultExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitMultExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListExpressionContext struct {
	*ExpressionContext
}

func NewListExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExpressionContext {
	var p = new(ListExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExpressionContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ListExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *ListExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitListExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TinyLanguageParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, TinyLanguageParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(151)
			p.Match(TinyLanguageParserSubtract)
		}
		{
			p.SetState(152)
			p.expression(15)
		}

	case 2:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(TinyLanguageParserNumber)
		}

	case 3:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Match(TinyLanguageParserBool)
		}

	case 4:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)
			p.FunctionCall()
		}

	case 5:
		localctx = NewListExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(156)
			p.List()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(157)
				p.Indexes()
			}

		}

	case 6:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(TinyLanguageParserIdentifier)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.Indexes()
			}

		}

	case 7:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(TinyLanguageParserStringLiteral)
		}

	case 8:
		localctx = NewExpressionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(TinyLanguageParserT__2)
		}
		{
			p.SetState(166)
			p.expression(0)
		}
		{
			p.SetState(167)
			p.Match(TinyLanguageParserT__3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(172)
					p.Match(TinyLanguageParserPower)
				}
				{
					p.SetState(173)
					p.expression(14)
				}

			case 2:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(175)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(TinyLanguageParserMult-32))|(1<<(TinyLanguageParserDivide-32))|(1<<(TinyLanguageParserMod-32)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(176)
					p.expression(14)
				}

			case 3:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(178)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TinyLanguageParserAdd || _la == TinyLanguageParserSubtract) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(179)
					p.expression(13)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(181)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(TinyLanguageParserGTEquals-27))|(1<<(TinyLanguageParserLTEquals-27))|(1<<(TinyLanguageParserGT-27))|(1<<(TinyLanguageParserLT-27)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(182)
					p.expression(12)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(184)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TinyLanguageParserEquals || _la == TinyLanguageParserNEquals) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(185)
					p.expression(11)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(187)
					p.Match(TinyLanguageParserAnd)
				}
				{
					p.SetState(188)
					p.expression(10)
				}

			case 7:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(190)
					p.Match(TinyLanguageParserOr)
				}
				{
					p.SetState(191)
					p.expression(9)
				}

			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TinyLanguageParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(TinyLanguageParserT__5)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLanguageParserT__2)|(1<<TinyLanguageParserT__5)|(1<<TinyLanguageParserPrintln)|(1<<TinyLanguageParserBool)|(1<<TinyLanguageParserNumber)|(1<<TinyLanguageParserIdentifier)|(1<<TinyLanguageParserStringLiteral)|(1<<TinyLanguageParserSubtract))) != 0 {
		{
			p.SetState(198)
			p.ExprList()
		}

	}
	{
		p.SetState(201)
		p.Match(TinyLanguageParserT__6)
	}

	return localctx
}

// IIndexesContext is an interface to support dynamic dispatch.
type IIndexesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexesContext differentiates from other interfaces.
	IsIndexesContext()
}

type IndexesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexesContext() *IndexesContext {
	var p = new(IndexesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLanguageParserRULE_indexes
	return p
}

func (*IndexesContext) IsIndexesContext() {}

func NewIndexesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexesContext {
	var p = new(IndexesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLanguageParserRULE_indexes

	return p
}

func (s *IndexesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexesContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IndexesContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitIndexes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Indexes() (localctx IIndexesContext) {
	localctx = NewIndexesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TinyLanguageParserRULE_indexes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(203)
				p.Match(TinyLanguageParserT__5)
			}
			{
				p.SetState(204)
				p.expression(0)
			}
			{
				p.SetState(205)
				p.Match(TinyLanguageParserT__6)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

func (p *TinyLanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TinyLanguageParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
