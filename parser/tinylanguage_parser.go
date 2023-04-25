// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // TinyLanguage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TinyLanguageParser struct {
	*antlr.BaseParser
}

var tinylanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinylanguageParserInit() {
	staticData := &tinylanguageParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'='", "'('", "')'", "','", "'?'", "':'", "'['", "']'", "'println'",
		"'assert'", "'size'", "'def'", "'if'", "'else'", "'do'", "'return'",
		"'for'", "'while'", "'to'", "'end'", "'in'", "", "", "", "", "", "'||'",
		"'&&'", "'=='", "'!='", "'>='", "'<='", "'^'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'>'", "'<'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "Println", "Assert", "Size",
		"Def", "If", "Else", "Do", "Return", "For", "While", "To", "End", "In",
		"Bool", "Number", "Identifier", "StringLiteral", "Space", "Or", "And",
		"Equals", "NEquals", "GTEquals", "LTEquals", "Power", "Add", "Subtract",
		"Mult", "Divide", "Mod", "GT", "LT",
	}
	staticData.ruleNames = []string{
		"prog", "block", "statement", "assignment", "functionDecl", "forStatement",
		"whileStatement", "functionCall", "ifStatement", "ifStat", "elseIfStat",
		"elseStat", "idList", "exprList", "expression", "list", "indexes",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 231, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 39, 8, 1, 10, 1, 12, 1, 42,
		9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 3, 3, 63, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 72, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 97, 8, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 116, 8, 7, 1, 8, 1, 8, 5, 8, 120, 8,
		8, 10, 8, 12, 8, 123, 9, 8, 1, 8, 3, 8, 126, 8, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 148, 8, 12, 10, 12, 12, 12,
		151, 9, 12, 1, 13, 1, 13, 1, 13, 5, 13, 156, 8, 13, 10, 13, 12, 13, 159,
		9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 169,
		8, 14, 1, 14, 1, 14, 3, 14, 173, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 180, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 212, 8, 14, 10, 14, 12, 14, 215, 9, 14, 1, 15, 1,
		15, 3, 15, 219, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16,
		227, 8, 16, 11, 16, 12, 16, 228, 1, 16, 0, 1, 28, 17, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 4, 1, 0, 37, 39, 1, 0, 35,
		36, 2, 0, 32, 33, 40, 41, 1, 0, 30, 31, 251, 0, 34, 1, 0, 0, 0, 2, 40,
		1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 60, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10,
		77, 1, 0, 0, 0, 12, 87, 1, 0, 0, 0, 14, 115, 1, 0, 0, 0, 16, 117, 1, 0,
		0, 0, 18, 129, 1, 0, 0, 0, 20, 134, 1, 0, 0, 0, 22, 140, 1, 0, 0, 0, 24,
		144, 1, 0, 0, 0, 26, 152, 1, 0, 0, 0, 28, 179, 1, 0, 0, 0, 30, 216, 1,
		0, 0, 0, 32, 226, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 1, 1, 0, 0, 0, 36,
		39, 3, 4, 2, 0, 37, 39, 3, 8, 4, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0,
		0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 47,
		1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 17, 0, 0, 44, 45, 3, 28, 14,
		0, 45, 46, 5, 1, 0, 0, 46, 48, 1, 0, 0, 0, 47, 43, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 3, 1, 0, 0, 0, 49, 50, 3, 6, 3, 0, 50, 51, 5, 1, 0, 0,
		51, 59, 1, 0, 0, 0, 52, 53, 3, 14, 7, 0, 53, 54, 5, 1, 0, 0, 54, 59, 1,
		0, 0, 0, 55, 59, 3, 16, 8, 0, 56, 59, 3, 10, 5, 0, 57, 59, 3, 12, 6, 0,
		58, 49, 1, 0, 0, 0, 58, 52, 1, 0, 0, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1,
		0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60, 62, 5, 25, 0, 0, 61,
		63, 3, 32, 16, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 1, 0,
		0, 0, 64, 65, 5, 2, 0, 0, 65, 66, 3, 28, 14, 0, 66, 7, 1, 0, 0, 0, 67,
		68, 5, 13, 0, 0, 68, 69, 5, 25, 0, 0, 69, 71, 5, 3, 0, 0, 70, 72, 3, 24,
		12, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73,
		74, 5, 4, 0, 0, 74, 75, 3, 2, 1, 0, 75, 76, 5, 21, 0, 0, 76, 9, 1, 0, 0,
		0, 77, 78, 5, 18, 0, 0, 78, 79, 5, 25, 0, 0, 79, 80, 5, 2, 0, 0, 80, 81,
		3, 28, 14, 0, 81, 82, 5, 20, 0, 0, 82, 83, 3, 28, 14, 0, 83, 84, 5, 16,
		0, 0, 84, 85, 3, 2, 1, 0, 85, 86, 5, 21, 0, 0, 86, 11, 1, 0, 0, 0, 87,
		88, 5, 19, 0, 0, 88, 89, 3, 28, 14, 0, 89, 90, 5, 16, 0, 0, 90, 91, 3,
		2, 1, 0, 91, 92, 5, 21, 0, 0, 92, 13, 1, 0, 0, 0, 93, 94, 5, 25, 0, 0,
		94, 96, 5, 3, 0, 0, 95, 97, 3, 26, 13, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 116, 5, 4, 0, 0, 99, 100, 5, 10, 0, 0,
		100, 102, 5, 3, 0, 0, 101, 103, 3, 28, 14, 0, 102, 101, 1, 0, 0, 0, 102,
		103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 116, 5, 4, 0, 0, 105, 106,
		5, 11, 0, 0, 106, 107, 5, 3, 0, 0, 107, 108, 3, 28, 14, 0, 108, 109, 5,
		4, 0, 0, 109, 116, 1, 0, 0, 0, 110, 111, 5, 12, 0, 0, 111, 112, 5, 3, 0,
		0, 112, 113, 3, 28, 14, 0, 113, 114, 5, 4, 0, 0, 114, 116, 1, 0, 0, 0,
		115, 93, 1, 0, 0, 0, 115, 99, 1, 0, 0, 0, 115, 105, 1, 0, 0, 0, 115, 110,
		1, 0, 0, 0, 116, 15, 1, 0, 0, 0, 117, 121, 3, 18, 9, 0, 118, 120, 3, 20,
		10, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		126, 3, 22, 11, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 128, 5, 21, 0, 0, 128, 17, 1, 0, 0, 0, 129, 130, 5, 14,
		0, 0, 130, 131, 3, 28, 14, 0, 131, 132, 5, 16, 0, 0, 132, 133, 3, 2, 1,
		0, 133, 19, 1, 0, 0, 0, 134, 135, 5, 15, 0, 0, 135, 136, 5, 14, 0, 0, 136,
		137, 3, 28, 14, 0, 137, 138, 5, 16, 0, 0, 138, 139, 3, 2, 1, 0, 139, 21,
		1, 0, 0, 0, 140, 141, 5, 15, 0, 0, 141, 142, 5, 16, 0, 0, 142, 143, 3,
		2, 1, 0, 143, 23, 1, 0, 0, 0, 144, 149, 5, 25, 0, 0, 145, 146, 5, 5, 0,
		0, 146, 148, 5, 25, 0, 0, 147, 145, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 25, 1, 0, 0, 0, 151, 149, 1,
		0, 0, 0, 152, 157, 3, 28, 14, 0, 153, 154, 5, 5, 0, 0, 154, 156, 3, 28,
		14, 0, 155, 153, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 27, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161,
		6, 14, -1, 0, 161, 162, 5, 36, 0, 0, 162, 180, 3, 28, 14, 17, 163, 180,
		5, 24, 0, 0, 164, 180, 5, 23, 0, 0, 165, 180, 3, 14, 7, 0, 166, 168, 3,
		30, 15, 0, 167, 169, 3, 32, 16, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0,
		0, 0, 169, 180, 1, 0, 0, 0, 170, 172, 5, 25, 0, 0, 171, 173, 3, 32, 16,
		0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 180, 1, 0, 0, 0, 174,
		180, 5, 26, 0, 0, 175, 176, 5, 3, 0, 0, 176, 177, 3, 28, 14, 0, 177, 178,
		5, 4, 0, 0, 178, 180, 1, 0, 0, 0, 179, 160, 1, 0, 0, 0, 179, 163, 1, 0,
		0, 0, 179, 164, 1, 0, 0, 0, 179, 165, 1, 0, 0, 0, 179, 166, 1, 0, 0, 0,
		179, 170, 1, 0, 0, 0, 179, 174, 1, 0, 0, 0, 179, 175, 1, 0, 0, 0, 180,
		213, 1, 0, 0, 0, 181, 182, 10, 16, 0, 0, 182, 183, 5, 34, 0, 0, 183, 212,
		3, 28, 14, 16, 184, 185, 10, 15, 0, 0, 185, 186, 7, 0, 0, 0, 186, 212,
		3, 28, 14, 16, 187, 188, 10, 14, 0, 0, 188, 189, 7, 1, 0, 0, 189, 212,
		3, 28, 14, 15, 190, 191, 10, 13, 0, 0, 191, 192, 7, 2, 0, 0, 192, 212,
		3, 28, 14, 14, 193, 194, 10, 12, 0, 0, 194, 195, 7, 3, 0, 0, 195, 212,
		3, 28, 14, 13, 196, 197, 10, 11, 0, 0, 197, 198, 5, 29, 0, 0, 198, 212,
		3, 28, 14, 12, 199, 200, 10, 10, 0, 0, 200, 201, 5, 28, 0, 0, 201, 212,
		3, 28, 14, 11, 202, 203, 10, 9, 0, 0, 203, 204, 5, 6, 0, 0, 204, 205, 3,
		28, 14, 0, 205, 206, 5, 7, 0, 0, 206, 207, 3, 28, 14, 10, 207, 212, 1,
		0, 0, 0, 208, 209, 10, 8, 0, 0, 209, 210, 5, 22, 0, 0, 210, 212, 3, 28,
		14, 9, 211, 181, 1, 0, 0, 0, 211, 184, 1, 0, 0, 0, 211, 187, 1, 0, 0, 0,
		211, 190, 1, 0, 0, 0, 211, 193, 1, 0, 0, 0, 211, 196, 1, 0, 0, 0, 211,
		199, 1, 0, 0, 0, 211, 202, 1, 0, 0, 0, 211, 208, 1, 0, 0, 0, 212, 215,
		1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 29, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 216, 218, 5, 8, 0, 0, 217, 219, 3, 26, 13,
		0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220,
		221, 5, 9, 0, 0, 221, 31, 1, 0, 0, 0, 222, 223, 5, 8, 0, 0, 223, 224, 3,
		28, 14, 0, 224, 225, 5, 9, 0, 0, 225, 227, 1, 0, 0, 0, 226, 222, 1, 0,
		0, 0, 227, 228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 33, 1, 0, 0, 0, 20, 38, 40, 47, 58, 62, 71, 96, 102, 115, 121, 125,
		149, 157, 168, 172, 179, 211, 213, 218, 228,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TinyLanguageParserInit initializes any static state used to implement TinyLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTinyLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinyLanguageParserInit() {
	staticData := &tinylanguageParserStaticData
	staticData.once.Do(tinylanguageParserInit)
}

// NewTinyLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTinyLanguageParser(input antlr.TokenStream) *TinyLanguageParser {
	TinyLanguageParserInit()
	this := new(TinyLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tinylanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
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
	TinyLanguageParserT__7          = 8
	TinyLanguageParserT__8          = 9
	TinyLanguageParserPrintln       = 10
	TinyLanguageParserAssert        = 11
	TinyLanguageParserSize          = 12
	TinyLanguageParserDef           = 13
	TinyLanguageParserIf            = 14
	TinyLanguageParserElse          = 15
	TinyLanguageParserDo            = 16
	TinyLanguageParserReturn        = 17
	TinyLanguageParserFor           = 18
	TinyLanguageParserWhile         = 19
	TinyLanguageParserTo            = 20
	TinyLanguageParserEnd           = 21
	TinyLanguageParserIn            = 22
	TinyLanguageParserBool          = 23
	TinyLanguageParserNumber        = 24
	TinyLanguageParserIdentifier    = 25
	TinyLanguageParserStringLiteral = 26
	TinyLanguageParserSpace         = 27
	TinyLanguageParserOr            = 28
	TinyLanguageParserAnd           = 29
	TinyLanguageParserEquals        = 30
	TinyLanguageParserNEquals       = 31
	TinyLanguageParserGTEquals      = 32
	TinyLanguageParserLTEquals      = 33
	TinyLanguageParserPower         = 34
	TinyLanguageParserAdd           = 35
	TinyLanguageParserSubtract      = 36
	TinyLanguageParserMult          = 37
	TinyLanguageParserDivide        = 38
	TinyLanguageParserMod           = 39
	TinyLanguageParserGT            = 40
	TinyLanguageParserLT            = 41
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

	// Getter signatures
	Block() IBlockContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFunctionDecl() []IFunctionDeclContext
	FunctionDecl(i int) IFunctionDeclContext
	Return() antlr.TerminalNode
	Expression() IExpressionContext

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllFunctionDecl() []IFunctionDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclContext); ok {
			tst[i] = t.(IFunctionDeclContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) FunctionDecl(i int) IFunctionDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *BlockContext) Return() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserReturn, 0)
}

func (s *BlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34372608) != 0 {
		p.SetState(38)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TinyLanguageParserPrintln, TinyLanguageParserAssert, TinyLanguageParserSize, TinyLanguageParserIf, TinyLanguageParserFor, TinyLanguageParserWhile, TinyLanguageParserIdentifier:
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

	// Getter signatures
	Assignment() IAssignmentContext
	FunctionCall() IFunctionCallContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	WhileStatement() IWhileStatementContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	// Getter signatures
	Identifier() antlr.TerminalNode
	Expression() IExpressionContext
	Indexes() IIndexesContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) Indexes() IIndexesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	if _la == TinyLanguageParserT__7 {
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

	// Getter signatures
	Def() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Block() IBlockContext
	End() antlr.TerminalNode
	IdList() IIdListContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *FunctionDeclContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	// Getter signatures
	For() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	To() antlr.TerminalNode
	Do() antlr.TerminalNode
	Block() IBlockContext
	End() antlr.TerminalNode

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	// Getter signatures
	While() antlr.TerminalNode
	Expression() IExpressionContext
	Do() antlr.TerminalNode
	Block() IBlockContext
	End() antlr.TerminalNode

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

type AssertFunctionCallContext struct {
	*FunctionCallContext
}

func NewAssertFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssertFunctionCallContext {
	var p = new(AssertFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *AssertFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertFunctionCallContext) Assert() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserAssert, 0)
}

func (s *AssertFunctionCallContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssertFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitAssertFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type SizeFunctionCallContext struct {
	*FunctionCallContext
}

func NewSizeFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SizeFunctionCallContext {
	var p = new(SizeFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *SizeFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeFunctionCallContext) Size() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserSize, 0)
}

func (s *SizeFunctionCallContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SizeFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitSizeFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	p.SetState(115)
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68845313288) != 0 {
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68845313288) != 0 {
			{
				p.SetState(101)
				p.expression(0)
			}

		}
		{
			p.SetState(104)
			p.Match(TinyLanguageParserT__3)
		}

	case TinyLanguageParserAssert:
		localctx = NewAssertFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(TinyLanguageParserAssert)
		}
		{
			p.SetState(106)
			p.Match(TinyLanguageParserT__2)
		}
		{
			p.SetState(107)
			p.expression(0)
		}
		{
			p.SetState(108)
			p.Match(TinyLanguageParserT__3)
		}

	case TinyLanguageParserSize:
		localctx = NewSizeFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(TinyLanguageParserSize)
		}
		{
			p.SetState(111)
			p.Match(TinyLanguageParserT__2)
		}
		{
			p.SetState(112)
			p.expression(0)
		}
		{
			p.SetState(113)
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

	// Getter signatures
	IfStat() IIfStatContext
	End() antlr.TerminalNode
	AllElseIfStat() []IElseIfStatContext
	ElseIfStat(i int) IElseIfStatContext
	ElseStat() IElseStatContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *IfStatementContext) End() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserEnd, 0)
}

func (s *IfStatementContext) AllElseIfStat() []IElseIfStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStatContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStatContext); ok {
			tst[i] = t.(IElseIfStatContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStat(i int) IElseIfStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfStatContext)
}

func (s *IfStatementContext) ElseStat() IElseStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(117)
		p.IfStat()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				p.ElseIfStat()
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLanguageParserElse {
		{
			p.SetState(124)
			p.ElseStat()
		}

	}
	{
		p.SetState(127)
		p.Match(TinyLanguageParserEnd)
	}

	return localctx
}

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If() antlr.TerminalNode
	Expression() IExpressionContext
	Do() antlr.TerminalNode
	Block() IBlockContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *IfStatContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(129)
		p.Match(TinyLanguageParserIf)
	}
	{
		p.SetState(130)
		p.expression(0)
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

// IElseIfStatContext is an interface to support dynamic dispatch.
type IElseIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Else() antlr.TerminalNode
	If() antlr.TerminalNode
	Expression() IExpressionContext
	Do() antlr.TerminalNode
	Block() IBlockContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStatContext) Do() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserDo, 0)
}

func (s *ElseIfStatContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(134)
		p.Match(TinyLanguageParserElse)
	}
	{
		p.SetState(135)
		p.Match(TinyLanguageParserIf)
	}
	{
		p.SetState(136)
		p.expression(0)
	}
	{
		p.SetState(137)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(138)
		p.Block()
	}

	return localctx
}

// IElseStatContext is an interface to support dynamic dispatch.
type IElseStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Else() antlr.TerminalNode
	Do() antlr.TerminalNode
	Block() IBlockContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(140)
		p.Match(TinyLanguageParserElse)
	}
	{
		p.SetState(141)
		p.Match(TinyLanguageParserDo)
	}
	{
		p.SetState(142)
		p.Block()
	}

	return localctx
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

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
	this := p
	_ = this

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
		p.SetState(144)
		p.Match(TinyLanguageParserIdentifier)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLanguageParserT__4 {
		{
			p.SetState(145)
			p.Match(TinyLanguageParserT__4)
		}
		{
			p.SetState(146)
			p.Match(TinyLanguageParserIdentifier)
		}

		p.SetState(151)
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

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
		p.SetState(152)
		p.expression(0)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLanguageParserT__4 {
		{
			p.SetState(153)
			p.Match(TinyLanguageParserT__4)
		}
		{
			p.SetState(154)
			p.expression(0)
		}

		p.SetState(159)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(TinyLanguageParserIn, 0)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitInExpression(s)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ListExpressionContext) Indexes() IIndexesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

type TernaryExpressionContext struct {
	*ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLanguageVisitor:
		return t.VisitTernaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLanguageParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TinyLanguageParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(161)
			p.Match(TinyLanguageParserSubtract)
		}
		{
			p.SetState(162)
			p.expression(17)
		}

	case 2:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.Match(TinyLanguageParserNumber)
		}

	case 3:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(TinyLanguageParserBool)
		}

	case 4:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.FunctionCall()
		}

	case 5:
		localctx = NewListExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(166)
			p.List()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(167)
				p.Indexes()
			}

		}

	case 6:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.Match(TinyLanguageParserIdentifier)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(171)
				p.Indexes()
			}

		}

	case 7:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(174)
			p.Match(TinyLanguageParserStringLiteral)
		}

	case 8:
		localctx = NewExpressionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(175)
			p.Match(TinyLanguageParserT__2)
		}
		{
			p.SetState(176)
			p.expression(0)
		}
		{
			p.SetState(177)
			p.Match(TinyLanguageParserT__3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(182)
					p.Match(TinyLanguageParserPower)
				}
				{
					p.SetState(183)
					p.expression(16)
				}

			case 2:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(185)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&962072674304) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(186)
					p.expression(16)
				}

			case 3:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(188)

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
					p.SetState(189)
					p.expression(15)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(191)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3311419785216) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(192)
					p.expression(14)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(194)

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
					p.SetState(195)
					p.expression(13)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(197)
					p.Match(TinyLanguageParserAnd)
				}
				{
					p.SetState(198)
					p.expression(12)
				}

			case 7:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(200)
					p.Match(TinyLanguageParserOr)
				}
				{
					p.SetState(201)
					p.expression(11)
				}

			case 8:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(202)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
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
				{
					p.SetState(206)
					p.expression(10)
				}

			case 9:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinyLanguageParserRULE_expression)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(209)
					p.Match(TinyLanguageParserIn)
				}
				{
					p.SetState(210)
					p.expression(9)
				}

			}

		}
		p.SetState(215)
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

	// Getter signatures
	ExprList() IExprListContext

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(216)
		p.Match(TinyLanguageParserT__7)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68845313288) != 0 {
		{
			p.SetState(217)
			p.ExprList()
		}

	}
	{
		p.SetState(220)
		p.Match(TinyLanguageParserT__8)
	}

	return localctx
}

// IIndexesContext is an interface to support dynamic dispatch.
type IIndexesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexesContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(222)
				p.Match(TinyLanguageParserT__7)
			}
			{
				p.SetState(223)
				p.expression(0)
			}
			{
				p.SetState(224)
				p.Match(TinyLanguageParserT__8)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(228)
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
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
