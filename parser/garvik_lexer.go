// Code generated from Garvik.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 155,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 126, 10, 24, 3, 25, 6, 25, 129, 10, 25,
	13, 25, 14, 25, 130, 3, 26, 6, 26, 134, 10, 26, 13, 26, 14, 26, 135, 3,
	27, 3, 27, 3, 27, 3, 27, 7, 27, 142, 10, 27, 12, 27, 14, 27, 145, 11, 27,
	3, 27, 3, 27, 3, 28, 6, 28, 150, 10, 28, 13, 28, 14, 28, 151, 3, 28, 3,
	28, 2, 2, 29, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 3, 2, 7, 3, 2, 99, 124, 3, 2, 50, 59, 10, 2, 36, 36, 41, 41, 94, 94,
	100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 6, 2, 12, 12, 15, 15,
	36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 160, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 3, 57, 3, 2, 2, 2,
	5, 60, 3, 2, 2, 2, 7, 62, 3, 2, 2, 2, 9, 64, 3, 2, 2, 2, 11, 66, 3, 2,
	2, 2, 13, 68, 3, 2, 2, 2, 15, 70, 3, 2, 2, 2, 17, 72, 3, 2, 2, 2, 19, 74,
	3, 2, 2, 2, 21, 76, 3, 2, 2, 2, 23, 78, 3, 2, 2, 2, 25, 80, 3, 2, 2, 2,
	27, 83, 3, 2, 2, 2, 29, 86, 3, 2, 2, 2, 31, 88, 3, 2, 2, 2, 33, 90, 3,
	2, 2, 2, 35, 92, 3, 2, 2, 2, 37, 95, 3, 2, 2, 2, 39, 100, 3, 2, 2, 2, 41,
	105, 3, 2, 2, 2, 43, 109, 3, 2, 2, 2, 45, 112, 3, 2, 2, 2, 47, 125, 3,
	2, 2, 2, 49, 128, 3, 2, 2, 2, 51, 133, 3, 2, 2, 2, 53, 137, 3, 2, 2, 2,
	55, 149, 3, 2, 2, 2, 57, 58, 7, 63, 2, 2, 58, 59, 7, 63, 2, 2, 59, 4, 3,
	2, 2, 2, 60, 61, 7, 48, 2, 2, 61, 6, 3, 2, 2, 2, 62, 63, 7, 93, 2, 2, 63,
	8, 3, 2, 2, 2, 64, 65, 7, 95, 2, 2, 65, 10, 3, 2, 2, 2, 66, 67, 7, 63,
	2, 2, 67, 12, 3, 2, 2, 2, 68, 69, 7, 42, 2, 2, 69, 14, 3, 2, 2, 2, 70,
	71, 7, 43, 2, 2, 71, 16, 3, 2, 2, 2, 72, 73, 7, 49, 2, 2, 73, 18, 3, 2,
	2, 2, 74, 75, 7, 44, 2, 2, 75, 20, 3, 2, 2, 2, 76, 77, 7, 47, 2, 2, 77,
	22, 3, 2, 2, 2, 78, 79, 7, 45, 2, 2, 79, 24, 3, 2, 2, 2, 80, 81, 7, 47,
	2, 2, 81, 82, 7, 64, 2, 2, 82, 26, 3, 2, 2, 2, 83, 84, 7, 49, 2, 2, 84,
	85, 7, 49, 2, 2, 85, 28, 3, 2, 2, 2, 86, 87, 7, 46, 2, 2, 87, 30, 3, 2,
	2, 2, 88, 89, 7, 125, 2, 2, 89, 32, 3, 2, 2, 2, 90, 91, 7, 127, 2, 2, 91,
	34, 3, 2, 2, 2, 92, 93, 7, 107, 2, 2, 93, 94, 7, 104, 2, 2, 94, 36, 3,
	2, 2, 2, 95, 96, 7, 118, 2, 2, 96, 97, 7, 106, 2, 2, 97, 98, 7, 103, 2,
	2, 98, 99, 7, 112, 2, 2, 99, 38, 3, 2, 2, 2, 100, 101, 7, 103, 2, 2, 101,
	102, 7, 110, 2, 2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 103, 2, 2, 104,
	40, 3, 2, 2, 2, 105, 106, 7, 110, 2, 2, 106, 107, 7, 103, 2, 2, 107, 108,
	7, 118, 2, 2, 108, 42, 3, 2, 2, 2, 109, 110, 7, 107, 2, 2, 110, 111, 7,
	112, 2, 2, 111, 44, 3, 2, 2, 2, 112, 113, 7, 110, 2, 2, 113, 114, 7, 103,
	2, 2, 114, 115, 7, 112, 2, 2, 115, 46, 3, 2, 2, 2, 116, 117, 7, 118, 2,
	2, 117, 118, 7, 116, 2, 2, 118, 119, 7, 119, 2, 2, 119, 126, 7, 103, 2,
	2, 120, 121, 7, 104, 2, 2, 121, 122, 7, 99, 2, 2, 122, 123, 7, 110, 2,
	2, 123, 124, 7, 117, 2, 2, 124, 126, 7, 103, 2, 2, 125, 116, 3, 2, 2, 2,
	125, 120, 3, 2, 2, 2, 126, 48, 3, 2, 2, 2, 127, 129, 9, 2, 2, 2, 128, 127,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2,
	2, 2, 131, 50, 3, 2, 2, 2, 132, 134, 9, 3, 2, 2, 133, 132, 3, 2, 2, 2,
	134, 135, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	52, 3, 2, 2, 2, 137, 143, 7, 36, 2, 2, 138, 139, 7, 94, 2, 2, 139, 142,
	9, 4, 2, 2, 140, 142, 10, 5, 2, 2, 141, 138, 3, 2, 2, 2, 141, 140, 3, 2,
	2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2,
	144, 146, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 147, 7, 36, 2, 2, 147,
	54, 3, 2, 2, 2, 148, 150, 9, 6, 2, 2, 149, 148, 3, 2, 2, 2, 150, 151, 3,
	2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2,
	2, 153, 154, 8, 28, 2, 2, 154, 56, 3, 2, 2, 2, 9, 2, 125, 130, 135, 141,
	143, 151, 3, 8, 2, 2,
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
	"", "'=='", "'.'", "'['", "']'", "'='", "'('", "')'", "'/'", "'*'", "'-'",
	"'+'", "'->'", "'//'", "','", "'{'", "'}'", "'if'", "'then'", "'else'",
	"'let'", "'in'", "'len'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "BOOL", "ID", "NUM", "STR", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "BOOL", "ID", "NUM", "STR",
	"WS",
}

type GarvikLexer struct {
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

func NewGarvikLexer(input antlr.CharStream) *GarvikLexer {

	l := new(GarvikLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Garvik.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GarvikLexer tokens.
const (
	GarvikLexerT__0  = 1
	GarvikLexerT__1  = 2
	GarvikLexerT__2  = 3
	GarvikLexerT__3  = 4
	GarvikLexerT__4  = 5
	GarvikLexerT__5  = 6
	GarvikLexerT__6  = 7
	GarvikLexerT__7  = 8
	GarvikLexerT__8  = 9
	GarvikLexerT__9  = 10
	GarvikLexerT__10 = 11
	GarvikLexerT__11 = 12
	GarvikLexerT__12 = 13
	GarvikLexerT__13 = 14
	GarvikLexerT__14 = 15
	GarvikLexerT__15 = 16
	GarvikLexerT__16 = 17
	GarvikLexerT__17 = 18
	GarvikLexerT__18 = 19
	GarvikLexerT__19 = 20
	GarvikLexerT__20 = 21
	GarvikLexerT__21 = 22
	GarvikLexerBOOL  = 23
	GarvikLexerID    = 24
	GarvikLexerNUM   = 25
	GarvikLexerSTR   = 26
	GarvikLexerWS    = 27
)