// Code generated from ./grammar/TinyLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TinyLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var tinylanguagelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinylanguagelexerLexerInit() {
	staticData := &tinylanguagelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"Println", "Assert", "Size", "Def", "If", "Else", "Do", "Return", "For",
		"While", "To", "End", "In", "Bool", "Number", "Identifier", "StringLiteral",
		"Space", "Int", "Digit", "Or", "And", "Equals", "NEquals", "GTEquals",
		"LTEquals", "Power", "Add", "Subtract", "Mult", "Divide", "Mod", "GT",
		"LT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 267, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 177, 8, 22, 1, 23, 1, 23, 1, 23,
		5, 23, 182, 8, 23, 10, 23, 12, 23, 185, 9, 23, 3, 23, 187, 8, 23, 1, 24,
		1, 24, 5, 24, 191, 8, 24, 10, 24, 12, 24, 194, 9, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 5, 25, 200, 8, 25, 10, 25, 12, 25, 203, 9, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 5, 25, 210, 8, 25, 10, 25, 12, 25, 213, 9, 25, 1,
		25, 3, 25, 216, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27,
		224, 8, 27, 10, 27, 12, 27, 227, 9, 27, 1, 27, 3, 27, 230, 8, 27, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 42, 1, 42, 0, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 0, 57, 0, 59, 28, 61, 29, 63, 30, 65, 31, 67, 32,
		69, 33, 71, 34, 73, 35, 75, 36, 77, 37, 79, 38, 81, 39, 83, 40, 85, 41,
		1, 0, 10, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		1, 0, 34, 34, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13,
		1, 0, 39, 39, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 3, 0, 9, 10, 12, 13,
		32, 32, 1, 0, 49, 57, 1, 0, 48, 57, 275, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0,
		0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1,
		0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77,
		1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0,
		85, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3, 89, 1, 0, 0, 0, 5, 91, 1, 0, 0, 0,
		7, 93, 1, 0, 0, 0, 9, 95, 1, 0, 0, 0, 11, 97, 1, 0, 0, 0, 13, 99, 1, 0,
		0, 0, 15, 101, 1, 0, 0, 0, 17, 103, 1, 0, 0, 0, 19, 105, 1, 0, 0, 0, 21,
		113, 1, 0, 0, 0, 23, 120, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 129, 1,
		0, 0, 0, 29, 132, 1, 0, 0, 0, 31, 137, 1, 0, 0, 0, 33, 140, 1, 0, 0, 0,
		35, 147, 1, 0, 0, 0, 37, 151, 1, 0, 0, 0, 39, 157, 1, 0, 0, 0, 41, 160,
		1, 0, 0, 0, 43, 164, 1, 0, 0, 0, 45, 176, 1, 0, 0, 0, 47, 178, 1, 0, 0,
		0, 49, 188, 1, 0, 0, 0, 51, 215, 1, 0, 0, 0, 53, 217, 1, 0, 0, 0, 55, 229,
		1, 0, 0, 0, 57, 231, 1, 0, 0, 0, 59, 233, 1, 0, 0, 0, 61, 236, 1, 0, 0,
		0, 63, 239, 1, 0, 0, 0, 65, 242, 1, 0, 0, 0, 67, 245, 1, 0, 0, 0, 69, 248,
		1, 0, 0, 0, 71, 251, 1, 0, 0, 0, 73, 253, 1, 0, 0, 0, 75, 255, 1, 0, 0,
		0, 77, 257, 1, 0, 0, 0, 79, 259, 1, 0, 0, 0, 81, 261, 1, 0, 0, 0, 83, 263,
		1, 0, 0, 0, 85, 265, 1, 0, 0, 0, 87, 88, 5, 59, 0, 0, 88, 2, 1, 0, 0, 0,
		89, 90, 5, 61, 0, 0, 90, 4, 1, 0, 0, 0, 91, 92, 5, 40, 0, 0, 92, 6, 1,
		0, 0, 0, 93, 94, 5, 41, 0, 0, 94, 8, 1, 0, 0, 0, 95, 96, 5, 44, 0, 0, 96,
		10, 1, 0, 0, 0, 97, 98, 5, 63, 0, 0, 98, 12, 1, 0, 0, 0, 99, 100, 5, 58,
		0, 0, 100, 14, 1, 0, 0, 0, 101, 102, 5, 91, 0, 0, 102, 16, 1, 0, 0, 0,
		103, 104, 5, 93, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5, 112, 0, 0, 106,
		107, 5, 114, 0, 0, 107, 108, 5, 105, 0, 0, 108, 109, 5, 110, 0, 0, 109,
		110, 5, 116, 0, 0, 110, 111, 5, 108, 0, 0, 111, 112, 5, 110, 0, 0, 112,
		20, 1, 0, 0, 0, 113, 114, 5, 97, 0, 0, 114, 115, 5, 115, 0, 0, 115, 116,
		5, 115, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118, 5, 114, 0, 0, 118, 119,
		5, 116, 0, 0, 119, 22, 1, 0, 0, 0, 120, 121, 5, 115, 0, 0, 121, 122, 5,
		105, 0, 0, 122, 123, 5, 122, 0, 0, 123, 124, 5, 101, 0, 0, 124, 24, 1,
		0, 0, 0, 125, 126, 5, 100, 0, 0, 126, 127, 5, 101, 0, 0, 127, 128, 5, 102,
		0, 0, 128, 26, 1, 0, 0, 0, 129, 130, 5, 105, 0, 0, 130, 131, 5, 102, 0,
		0, 131, 28, 1, 0, 0, 0, 132, 133, 5, 101, 0, 0, 133, 134, 5, 108, 0, 0,
		134, 135, 5, 115, 0, 0, 135, 136, 5, 101, 0, 0, 136, 30, 1, 0, 0, 0, 137,
		138, 5, 100, 0, 0, 138, 139, 5, 111, 0, 0, 139, 32, 1, 0, 0, 0, 140, 141,
		5, 114, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 116, 0, 0, 143, 144,
		5, 117, 0, 0, 144, 145, 5, 114, 0, 0, 145, 146, 5, 110, 0, 0, 146, 34,
		1, 0, 0, 0, 147, 148, 5, 102, 0, 0, 148, 149, 5, 111, 0, 0, 149, 150, 5,
		114, 0, 0, 150, 36, 1, 0, 0, 0, 151, 152, 5, 119, 0, 0, 152, 153, 5, 104,
		0, 0, 153, 154, 5, 105, 0, 0, 154, 155, 5, 108, 0, 0, 155, 156, 5, 101,
		0, 0, 156, 38, 1, 0, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 111, 0,
		0, 159, 40, 1, 0, 0, 0, 160, 161, 5, 101, 0, 0, 161, 162, 5, 110, 0, 0,
		162, 163, 5, 100, 0, 0, 163, 42, 1, 0, 0, 0, 164, 165, 5, 105, 0, 0, 165,
		166, 5, 110, 0, 0, 166, 44, 1, 0, 0, 0, 167, 168, 5, 116, 0, 0, 168, 169,
		5, 114, 0, 0, 169, 170, 5, 117, 0, 0, 170, 177, 5, 101, 0, 0, 171, 172,
		5, 102, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 108, 0, 0, 174, 175,
		5, 115, 0, 0, 175, 177, 5, 101, 0, 0, 176, 167, 1, 0, 0, 0, 176, 171, 1,
		0, 0, 0, 177, 46, 1, 0, 0, 0, 178, 186, 3, 55, 27, 0, 179, 183, 5, 46,
		0, 0, 180, 182, 3, 57, 28, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0,
		0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185,
		183, 1, 0, 0, 0, 186, 179, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 48, 1,
		0, 0, 0, 188, 192, 7, 0, 0, 0, 189, 191, 7, 1, 0, 0, 190, 189, 1, 0, 0,
		0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193,
		50, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 201, 7, 2, 0, 0, 196, 200, 8,
		3, 0, 0, 197, 198, 5, 92, 0, 0, 198, 200, 8, 4, 0, 0, 199, 196, 1, 0, 0,
		0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 216,
		7, 2, 0, 0, 205, 211, 7, 5, 0, 0, 206, 210, 8, 6, 0, 0, 207, 208, 5, 92,
		0, 0, 208, 210, 8, 4, 0, 0, 209, 206, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0,
		210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212,
		214, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 216, 7, 5, 0, 0, 215, 195,
		1, 0, 0, 0, 215, 205, 1, 0, 0, 0, 216, 52, 1, 0, 0, 0, 217, 218, 7, 7,
		0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 6, 26, 0, 0, 220, 54, 1, 0, 0, 0,
		221, 225, 7, 8, 0, 0, 222, 224, 3, 57, 28, 0, 223, 222, 1, 0, 0, 0, 224,
		227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 230,
		1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 230, 5, 48, 0, 0, 229, 221, 1, 0,
		0, 0, 229, 228, 1, 0, 0, 0, 230, 56, 1, 0, 0, 0, 231, 232, 7, 9, 0, 0,
		232, 58, 1, 0, 0, 0, 233, 234, 5, 124, 0, 0, 234, 235, 5, 124, 0, 0, 235,
		60, 1, 0, 0, 0, 236, 237, 5, 38, 0, 0, 237, 238, 5, 38, 0, 0, 238, 62,
		1, 0, 0, 0, 239, 240, 5, 61, 0, 0, 240, 241, 5, 61, 0, 0, 241, 64, 1, 0,
		0, 0, 242, 243, 5, 33, 0, 0, 243, 244, 5, 61, 0, 0, 244, 66, 1, 0, 0, 0,
		245, 246, 5, 62, 0, 0, 246, 247, 5, 61, 0, 0, 247, 68, 1, 0, 0, 0, 248,
		249, 5, 60, 0, 0, 249, 250, 5, 61, 0, 0, 250, 70, 1, 0, 0, 0, 251, 252,
		5, 94, 0, 0, 252, 72, 1, 0, 0, 0, 253, 254, 5, 43, 0, 0, 254, 74, 1, 0,
		0, 0, 255, 256, 5, 45, 0, 0, 256, 76, 1, 0, 0, 0, 257, 258, 5, 42, 0, 0,
		258, 78, 1, 0, 0, 0, 259, 260, 5, 47, 0, 0, 260, 80, 1, 0, 0, 0, 261, 262,
		5, 37, 0, 0, 262, 82, 1, 0, 0, 0, 263, 264, 5, 62, 0, 0, 264, 84, 1, 0,
		0, 0, 265, 266, 5, 60, 0, 0, 266, 86, 1, 0, 0, 0, 12, 0, 176, 183, 186,
		192, 199, 201, 209, 211, 215, 225, 229, 1, 6, 0, 0,
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

// TinyLanguageLexerInit initializes any static state used to implement TinyLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTinyLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinyLanguageLexerInit() {
	staticData := &tinylanguagelexerLexerStaticData
	staticData.once.Do(tinylanguagelexerLexerInit)
}

// NewTinyLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTinyLanguageLexer(input antlr.CharStream) *TinyLanguageLexer {
	TinyLanguageLexerInit()
	l := new(TinyLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &tinylanguagelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
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
	TinyLanguageLexerT__5          = 6
	TinyLanguageLexerT__6          = 7
	TinyLanguageLexerT__7          = 8
	TinyLanguageLexerT__8          = 9
	TinyLanguageLexerPrintln       = 10
	TinyLanguageLexerAssert        = 11
	TinyLanguageLexerSize          = 12
	TinyLanguageLexerDef           = 13
	TinyLanguageLexerIf            = 14
	TinyLanguageLexerElse          = 15
	TinyLanguageLexerDo            = 16
	TinyLanguageLexerReturn        = 17
	TinyLanguageLexerFor           = 18
	TinyLanguageLexerWhile         = 19
	TinyLanguageLexerTo            = 20
	TinyLanguageLexerEnd           = 21
	TinyLanguageLexerIn            = 22
	TinyLanguageLexerBool          = 23
	TinyLanguageLexerNumber        = 24
	TinyLanguageLexerIdentifier    = 25
	TinyLanguageLexerStringLiteral = 26
	TinyLanguageLexerSpace         = 27
	TinyLanguageLexerOr            = 28
	TinyLanguageLexerAnd           = 29
	TinyLanguageLexerEquals        = 30
	TinyLanguageLexerNEquals       = 31
	TinyLanguageLexerGTEquals      = 32
	TinyLanguageLexerLTEquals      = 33
	TinyLanguageLexerPower         = 34
	TinyLanguageLexerAdd           = 35
	TinyLanguageLexerSubtract      = 36
	TinyLanguageLexerMult          = 37
	TinyLanguageLexerDivide        = 38
	TinyLanguageLexerMod           = 39
	TinyLanguageLexerGT            = 40
	TinyLanguageLexerLT            = 41
)
