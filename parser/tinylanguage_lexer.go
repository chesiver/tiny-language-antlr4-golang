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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 104,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7,
	7, 49, 10, 7, 12, 7, 14, 7, 52, 11, 7, 5, 7, 54, 10, 7, 3, 8, 3, 8, 7,
	8, 58, 10, 8, 12, 8, 14, 8, 61, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 67,
	10, 9, 12, 9, 14, 9, 70, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 77,
	10, 9, 12, 9, 14, 9, 80, 11, 9, 3, 9, 5, 9, 83, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 7, 11, 91, 10, 11, 12, 11, 14, 11, 94, 11, 11,
	3, 11, 5, 11, 97, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 2,
	2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	2, 23, 2, 25, 12, 27, 13, 3, 2, 12, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59,
	67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 6, 2, 12, 12, 15, 15, 36, 36, 94,
	94, 4, 2, 12, 12, 15, 15, 3, 2, 41, 41, 6, 2, 12, 12, 15, 15, 41, 41, 94,
	94, 5, 2, 11, 12, 14, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50, 59, 2, 111, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2,
	2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 37, 3,
	2, 2, 2, 13, 45, 3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 82, 3, 2, 2, 2, 19,
	84, 3, 2, 2, 2, 21, 96, 3, 2, 2, 2, 23, 98, 3, 2, 2, 2, 25, 100, 3, 2,
	2, 2, 27, 102, 3, 2, 2, 2, 29, 30, 7, 61, 2, 2, 30, 4, 3, 2, 2, 2, 31,
	32, 7, 63, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 8, 3, 2, 2,
	2, 35, 36, 7, 43, 2, 2, 36, 10, 3, 2, 2, 2, 37, 38, 7, 114, 2, 2, 38, 39,
	7, 116, 2, 2, 39, 40, 7, 107, 2, 2, 40, 41, 7, 112, 2, 2, 41, 42, 7, 118,
	2, 2, 42, 43, 7, 110, 2, 2, 43, 44, 7, 112, 2, 2, 44, 12, 3, 2, 2, 2, 45,
	53, 5, 21, 11, 2, 46, 50, 7, 48, 2, 2, 47, 49, 5, 23, 12, 2, 48, 47, 3,
	2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 46, 3, 2, 2, 2, 53, 54, 3, 2, 2,
	2, 54, 14, 3, 2, 2, 2, 55, 59, 9, 2, 2, 2, 56, 58, 9, 3, 2, 2, 57, 56,
	3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 16, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 68, 9, 4, 2, 2, 63, 67, 10,
	5, 2, 2, 64, 65, 7, 94, 2, 2, 65, 67, 10, 6, 2, 2, 66, 63, 3, 2, 2, 2,
	66, 64, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 83, 9, 4, 2, 2, 72,
	78, 9, 7, 2, 2, 73, 77, 10, 8, 2, 2, 74, 75, 7, 94, 2, 2, 75, 77, 10, 6,
	2, 2, 76, 73, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76,
	3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2,
	81, 83, 9, 7, 2, 2, 82, 62, 3, 2, 2, 2, 82, 72, 3, 2, 2, 2, 83, 18, 3,
	2, 2, 2, 84, 85, 9, 9, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 8, 10, 2, 2, 87,
	20, 3, 2, 2, 2, 88, 92, 9, 10, 2, 2, 89, 91, 5, 23, 12, 2, 90, 89, 3, 2,
	2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 97,
	3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 97, 7, 50, 2, 2, 96, 88, 3, 2, 2, 2,
	96, 95, 3, 2, 2, 2, 97, 22, 3, 2, 2, 2, 98, 99, 9, 11, 2, 2, 99, 24, 3,
	2, 2, 2, 100, 101, 7, 45, 2, 2, 101, 26, 3, 2, 2, 2, 102, 103, 7, 47, 2,
	2, 103, 28, 3, 2, 2, 2, 13, 2, 50, 53, 59, 66, 68, 76, 78, 82, 92, 96,
	3, 8, 2, 2,
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
	"", "';'", "'='", "'('", "')'", "'println'", "", "", "", "", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "Println", "Number", "Identifier", "StringLiteral",
	"Space", "Add", "Subtract",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "Println", "Number", "Identifier", "StringLiteral",
	"Space", "Int", "Digit", "Add", "Subtract",
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
	TinyLanguageLexerAdd           = 10
	TinyLanguageLexerSubtract      = 11
)
