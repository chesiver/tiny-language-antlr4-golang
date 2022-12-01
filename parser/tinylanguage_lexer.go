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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 145,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	7, 11, 82, 10, 11, 12, 11, 14, 11, 85, 11, 11, 5, 11, 87, 10, 11, 3, 12,
	3, 12, 7, 12, 91, 10, 12, 12, 12, 14, 12, 94, 11, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 7, 13, 100, 10, 13, 12, 13, 14, 13, 103, 11, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 7, 13, 110, 10, 13, 12, 13, 14, 13, 113, 11, 13, 3,
	13, 5, 13, 116, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15,
	124, 10, 15, 12, 15, 14, 15, 127, 11, 15, 3, 15, 5, 15, 130, 10, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 22, 3, 22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 2, 31, 2, 33, 16,
	35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 3, 2, 12, 4, 2, 67, 92, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 6, 2, 12, 12, 15,
	15, 36, 36, 94, 94, 4, 2, 12, 12, 15, 15, 3, 2, 41, 41, 6, 2, 12, 12, 15,
	15, 41, 41, 94, 94, 5, 2, 11, 12, 14, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50,
	59, 2, 152, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 51, 3,
	2, 2, 2, 11, 53, 3, 2, 2, 2, 13, 55, 3, 2, 2, 2, 15, 63, 3, 2, 2, 2, 17,
	67, 3, 2, 2, 2, 19, 74, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23, 88, 3, 2, 2,
	2, 25, 115, 3, 2, 2, 2, 27, 117, 3, 2, 2, 2, 29, 129, 3, 2, 2, 2, 31, 131,
	3, 2, 2, 2, 33, 133, 3, 2, 2, 2, 35, 135, 3, 2, 2, 2, 37, 137, 3, 2, 2,
	2, 39, 139, 3, 2, 2, 2, 41, 141, 3, 2, 2, 2, 43, 143, 3, 2, 2, 2, 45, 46,
	7, 61, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 63, 2, 2, 48, 6, 3, 2, 2, 2,
	49, 50, 7, 42, 2, 2, 50, 8, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 10, 3,
	2, 2, 2, 53, 54, 7, 46, 2, 2, 54, 12, 3, 2, 2, 2, 55, 56, 7, 114, 2, 2,
	56, 57, 7, 116, 2, 2, 57, 58, 7, 107, 2, 2, 58, 59, 7, 112, 2, 2, 59, 60,
	7, 118, 2, 2, 60, 61, 7, 110, 2, 2, 61, 62, 7, 112, 2, 2, 62, 14, 3, 2,
	2, 2, 63, 64, 7, 102, 2, 2, 64, 65, 7, 103, 2, 2, 65, 66, 7, 104, 2, 2,
	66, 16, 3, 2, 2, 2, 67, 68, 7, 116, 2, 2, 68, 69, 7, 103, 2, 2, 69, 70,
	7, 118, 2, 2, 70, 71, 7, 119, 2, 2, 71, 72, 7, 116, 2, 2, 72, 73, 7, 112,
	2, 2, 73, 18, 3, 2, 2, 2, 74, 75, 7, 103, 2, 2, 75, 76, 7, 112, 2, 2, 76,
	77, 7, 102, 2, 2, 77, 20, 3, 2, 2, 2, 78, 86, 5, 29, 15, 2, 79, 83, 7,
	48, 2, 2, 80, 82, 5, 31, 16, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2,
	83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3,
	2, 2, 2, 86, 79, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 22, 3, 2, 2, 2, 88,
	92, 9, 2, 2, 2, 89, 91, 9, 3, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2,
	2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 24, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 95, 101, 9, 4, 2, 2, 96, 100, 10, 5, 2, 2, 97, 98, 7, 94, 2,
	2, 98, 100, 10, 6, 2, 2, 99, 96, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 103,
	3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2,
	2, 2, 103, 101, 3, 2, 2, 2, 104, 116, 9, 4, 2, 2, 105, 111, 9, 7, 2, 2,
	106, 110, 10, 8, 2, 2, 107, 108, 7, 94, 2, 2, 108, 110, 10, 6, 2, 2, 109,
	106, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2,
	2, 2, 114, 116, 9, 7, 2, 2, 115, 95, 3, 2, 2, 2, 115, 105, 3, 2, 2, 2,
	116, 26, 3, 2, 2, 2, 117, 118, 9, 9, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120,
	8, 14, 2, 2, 120, 28, 3, 2, 2, 2, 121, 125, 9, 10, 2, 2, 122, 124, 5, 31,
	16, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2,
	125, 126, 3, 2, 2, 2, 126, 130, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128,
	130, 7, 50, 2, 2, 129, 121, 3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 30,
	3, 2, 2, 2, 131, 132, 9, 11, 2, 2, 132, 32, 3, 2, 2, 2, 133, 134, 7, 96,
	2, 2, 134, 34, 3, 2, 2, 2, 135, 136, 7, 45, 2, 2, 136, 36, 3, 2, 2, 2,
	137, 138, 7, 47, 2, 2, 138, 38, 3, 2, 2, 2, 139, 140, 7, 44, 2, 2, 140,
	40, 3, 2, 2, 2, 141, 142, 7, 49, 2, 2, 142, 42, 3, 2, 2, 2, 143, 144, 7,
	39, 2, 2, 144, 44, 3, 2, 2, 2, 13, 2, 83, 86, 92, 99, 101, 109, 111, 115,
	125, 129, 3, 8, 2, 2,
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
	"", "';'", "'='", "'('", "')'", "','", "'println'", "'def'", "'return'",
	"'end'", "", "", "", "", "'^'", "'+'", "'-'", "'*'", "'/'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "Println", "Def", "Return", "End", "Number", "Identifier",
	"StringLiteral", "Space", "Power", "Add", "Subtract", "Mult", "Divide",
	"Mod",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "Println", "Def", "Return", "End",
	"Number", "Identifier", "StringLiteral", "Space", "Int", "Digit", "Power",
	"Add", "Subtract", "Mult", "Divide", "Mod",
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
	TinyLanguageLexerT__4          = 5
	TinyLanguageLexerPrintln       = 6
	TinyLanguageLexerDef           = 7
	TinyLanguageLexerReturn        = 8
	TinyLanguageLexerEnd           = 9
	TinyLanguageLexerNumber        = 10
	TinyLanguageLexerIdentifier    = 11
	TinyLanguageLexerStringLiteral = 12
	TinyLanguageLexerSpace         = 13
	TinyLanguageLexerPower         = 14
	TinyLanguageLexerAdd           = 15
	TinyLanguageLexerSubtract      = 16
	TinyLanguageLexerMult          = 17
	TinyLanguageLexerDivide        = 18
	TinyLanguageLexerMod           = 19
)
