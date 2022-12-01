// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 120,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 57, 10, 7,
	12, 7, 14, 7, 60, 11, 7, 5, 7, 62, 10, 7, 3, 8, 3, 8, 7, 8, 66, 10, 8,
	12, 8, 14, 8, 69, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 75, 10, 9, 12, 9,
	14, 9, 78, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 85, 10, 9, 12, 9,
	14, 9, 88, 11, 9, 3, 9, 5, 9, 91, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 7, 11, 99, 10, 11, 12, 11, 14, 11, 102, 11, 11, 3, 11, 5, 11,
	105, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 2, 23, 2, 25, 12, 27, 13, 29,
	14, 31, 15, 33, 16, 35, 17, 3, 2, 12, 4, 2, 67, 92, 99, 124, 6, 2, 50,
	59, 67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 6, 2, 12, 12, 15, 15, 36, 36,
	94, 94, 4, 2, 12, 12, 15, 15, 3, 2, 41, 41, 6, 2, 12, 12, 15, 15, 41, 41,
	94, 94, 5, 2, 11, 12, 14, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50, 59, 2, 127,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3,
	2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 43, 3, 2, 2, 2, 11, 45,
	3, 2, 2, 2, 13, 53, 3, 2, 2, 2, 15, 63, 3, 2, 2, 2, 17, 90, 3, 2, 2, 2,
	19, 92, 3, 2, 2, 2, 21, 104, 3, 2, 2, 2, 23, 106, 3, 2, 2, 2, 25, 108,
	3, 2, 2, 2, 27, 110, 3, 2, 2, 2, 29, 112, 3, 2, 2, 2, 31, 114, 3, 2, 2,
	2, 33, 116, 3, 2, 2, 2, 35, 118, 3, 2, 2, 2, 37, 38, 7, 61, 2, 2, 38, 4,
	3, 2, 2, 2, 39, 40, 7, 63, 2, 2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 42, 2, 2,
	42, 8, 3, 2, 2, 2, 43, 44, 7, 43, 2, 2, 44, 10, 3, 2, 2, 2, 45, 46, 7,
	114, 2, 2, 46, 47, 7, 116, 2, 2, 47, 48, 7, 107, 2, 2, 48, 49, 7, 112,
	2, 2, 49, 50, 7, 118, 2, 2, 50, 51, 7, 110, 2, 2, 51, 52, 7, 112, 2, 2,
	52, 12, 3, 2, 2, 2, 53, 61, 5, 21, 11, 2, 54, 58, 7, 48, 2, 2, 55, 57,
	5, 23, 12, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2,
	2, 58, 59, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 54,
	3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 14, 3, 2, 2, 2, 63, 67, 9, 2, 2, 2,
	64, 66, 9, 3, 2, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 16, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70,
	76, 9, 4, 2, 2, 71, 75, 10, 5, 2, 2, 72, 73, 7, 94, 2, 2, 73, 75, 10, 6,
	2, 2, 74, 71, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74,
	3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2,
	79, 91, 9, 4, 2, 2, 80, 86, 9, 7, 2, 2, 81, 85, 10, 8, 2, 2, 82, 83, 7,
	94, 2, 2, 83, 85, 10, 6, 2, 2, 84, 81, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2,
	85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 91, 9, 7, 2, 2, 90, 70, 3, 2, 2, 2, 90,
	80, 3, 2, 2, 2, 91, 18, 3, 2, 2, 2, 92, 93, 9, 9, 2, 2, 93, 94, 3, 2, 2,
	2, 94, 95, 8, 10, 2, 2, 95, 20, 3, 2, 2, 2, 96, 100, 9, 10, 2, 2, 97, 99,
	5, 23, 12, 2, 98, 97, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2,
	2, 100, 101, 3, 2, 2, 2, 101, 105, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103,
	105, 7, 50, 2, 2, 104, 96, 3, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 22, 3,
	2, 2, 2, 106, 107, 9, 11, 2, 2, 107, 24, 3, 2, 2, 2, 108, 109, 7, 96, 2,
	2, 109, 26, 3, 2, 2, 2, 110, 111, 7, 45, 2, 2, 111, 28, 3, 2, 2, 2, 112,
	113, 7, 47, 2, 2, 113, 30, 3, 2, 2, 2, 114, 115, 7, 44, 2, 2, 115, 32,
	3, 2, 2, 2, 116, 117, 7, 49, 2, 2, 117, 34, 3, 2, 2, 2, 118, 119, 7, 39,
	2, 2, 119, 36, 3, 2, 2, 2, 13, 2, 58, 61, 67, 74, 76, 84, 86, 90, 100,
	104, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'='", "'('", "')'", "'println'", "", "", "", "", "'^'", "'+'",
	"'-'", "'*'", "'/'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "Println", "Number", "Identifier", "StringLiteral",
	"Space", "Power", "Add", "Subtract", "Mult", "Divide", "Mod",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "Println", "Number", "Identifier", "StringLiteral",
	"Space", "Int", "Digit", "Power", "Add", "Subtract", "Mult", "Divide",
	"Mod",
}

type TinyLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTinyLanguageLexer(input antlr.CharStream) *TinyLanguageLexer {

	l := new(TinyLanguageLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TinyLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TinyLanguageLexer tokens.
const (
	TinyLanguageLexerT__0          = 1
	TinyLanguageLexerT__1          = 2
	TinyLanguageLexerT__2          = 3
	TinyLanguageLexerT__3          = 4
	TinyLanguageLexerPrintln       = 5
	TinyLanguageLexerNumber        = 6
	TinyLanguageLexerIdentifier    = 7
	TinyLanguageLexerStringLiteral = 8
	TinyLanguageLexerSpace         = 9
	TinyLanguageLexerPower         = 10
	TinyLanguageLexerAdd           = 11
	TinyLanguageLexerSubtract      = 12
	TinyLanguageLexerMult          = 13
	TinyLanguageLexerDivide        = 14
	TinyLanguageLexerMod           = 15
)
