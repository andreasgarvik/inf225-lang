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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 182,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 151, 10, 29,
	3, 30, 3, 30, 7, 30, 155, 10, 30, 12, 30, 14, 30, 158, 11, 30, 3, 31, 6,
	31, 161, 10, 31, 13, 31, 14, 31, 162, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32,
	169, 10, 32, 12, 32, 14, 32, 172, 11, 32, 3, 32, 3, 32, 3, 33, 6, 33, 177,
	10, 33, 13, 33, 14, 33, 178, 3, 33, 3, 33, 2, 2, 34, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 3, 2, 8, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 3, 2, 50, 59, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100,
	104, 104, 112, 112, 116, 116, 118, 118, 6, 2, 12, 12, 15, 15, 36, 36, 94,
	94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 187, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 67,
	3, 2, 2, 2, 5, 70, 3, 2, 2, 2, 7, 72, 3, 2, 2, 2, 9, 74, 3, 2, 2, 2, 11,
	76, 3, 2, 2, 2, 13, 78, 3, 2, 2, 2, 15, 80, 3, 2, 2, 2, 17, 82, 3, 2, 2,
	2, 19, 84, 3, 2, 2, 2, 21, 86, 3, 2, 2, 2, 23, 88, 3, 2, 2, 2, 25, 90,
	3, 2, 2, 2, 27, 92, 3, 2, 2, 2, 29, 94, 3, 2, 2, 2, 31, 97, 3, 2, 2, 2,
	33, 99, 3, 2, 2, 2, 35, 102, 3, 2, 2, 2, 37, 104, 3, 2, 2, 2, 39, 106,
	3, 2, 2, 2, 41, 108, 3, 2, 2, 2, 43, 111, 3, 2, 2, 2, 45, 116, 3, 2, 2,
	2, 47, 121, 3, 2, 2, 2, 49, 125, 3, 2, 2, 2, 51, 128, 3, 2, 2, 2, 53, 132,
	3, 2, 2, 2, 55, 137, 3, 2, 2, 2, 57, 150, 3, 2, 2, 2, 59, 152, 3, 2, 2,
	2, 61, 160, 3, 2, 2, 2, 63, 164, 3, 2, 2, 2, 65, 176, 3, 2, 2, 2, 67, 68,
	7, 63, 2, 2, 68, 69, 7, 63, 2, 2, 69, 4, 3, 2, 2, 2, 70, 71, 7, 62, 2,
	2, 71, 6, 3, 2, 2, 2, 72, 73, 7, 64, 2, 2, 73, 8, 3, 2, 2, 2, 74, 75, 7,
	48, 2, 2, 75, 10, 3, 2, 2, 2, 76, 77, 7, 63, 2, 2, 77, 12, 3, 2, 2, 2,
	78, 79, 7, 93, 2, 2, 79, 14, 3, 2, 2, 2, 80, 81, 7, 95, 2, 2, 81, 16, 3,
	2, 2, 2, 82, 83, 7, 42, 2, 2, 83, 18, 3, 2, 2, 2, 84, 85, 7, 43, 2, 2,
	85, 20, 3, 2, 2, 2, 86, 87, 7, 49, 2, 2, 87, 22, 3, 2, 2, 2, 88, 89, 7,
	44, 2, 2, 89, 24, 3, 2, 2, 2, 90, 91, 7, 47, 2, 2, 91, 26, 3, 2, 2, 2,
	92, 93, 7, 45, 2, 2, 93, 28, 3, 2, 2, 2, 94, 95, 7, 47, 2, 2, 95, 96, 7,
	64, 2, 2, 96, 30, 3, 2, 2, 2, 97, 98, 7, 60, 2, 2, 98, 32, 3, 2, 2, 2,
	99, 100, 7, 49, 2, 2, 100, 101, 7, 49, 2, 2, 101, 34, 3, 2, 2, 2, 102,
	103, 7, 46, 2, 2, 103, 36, 3, 2, 2, 2, 104, 105, 7, 125, 2, 2, 105, 38,
	3, 2, 2, 2, 106, 107, 7, 127, 2, 2, 107, 40, 3, 2, 2, 2, 108, 109, 7, 107,
	2, 2, 109, 110, 7, 104, 2, 2, 110, 42, 3, 2, 2, 2, 111, 112, 7, 118, 2,
	2, 112, 113, 7, 106, 2, 2, 113, 114, 7, 103, 2, 2, 114, 115, 7, 112, 2,
	2, 115, 44, 3, 2, 2, 2, 116, 117, 7, 103, 2, 2, 117, 118, 7, 110, 2, 2,
	118, 119, 7, 117, 2, 2, 119, 120, 7, 103, 2, 2, 120, 46, 3, 2, 2, 2, 121,
	122, 7, 110, 2, 2, 122, 123, 7, 103, 2, 2, 123, 124, 7, 118, 2, 2, 124,
	48, 3, 2, 2, 2, 125, 126, 7, 107, 2, 2, 126, 127, 7, 112, 2, 2, 127, 50,
	3, 2, 2, 2, 128, 129, 7, 114, 2, 2, 129, 130, 7, 113, 2, 2, 130, 131, 7,
	114, 2, 2, 131, 52, 3, 2, 2, 2, 132, 133, 7, 102, 2, 2, 133, 134, 7, 116,
	2, 2, 134, 135, 7, 113, 2, 2, 135, 136, 7, 114, 2, 2, 136, 54, 3, 2, 2,
	2, 137, 138, 7, 110, 2, 2, 138, 139, 7, 103, 2, 2, 139, 140, 7, 112, 2,
	2, 140, 56, 3, 2, 2, 2, 141, 142, 7, 118, 2, 2, 142, 143, 7, 116, 2, 2,
	143, 144, 7, 119, 2, 2, 144, 151, 7, 103, 2, 2, 145, 146, 7, 104, 2, 2,
	146, 147, 7, 99, 2, 2, 147, 148, 7, 110, 2, 2, 148, 149, 7, 117, 2, 2,
	149, 151, 7, 103, 2, 2, 150, 141, 3, 2, 2, 2, 150, 145, 3, 2, 2, 2, 151,
	58, 3, 2, 2, 2, 152, 156, 9, 2, 2, 2, 153, 155, 9, 3, 2, 2, 154, 153, 3,
	2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2,
	2, 157, 60, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 161, 9, 4, 2, 2, 160,
	159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 62, 3, 2, 2, 2, 164, 170, 7, 36, 2, 2, 165, 166, 7, 94,
	2, 2, 166, 169, 9, 5, 2, 2, 167, 169, 10, 6, 2, 2, 168, 165, 3, 2, 2, 2,
	168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170,
	171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 174,
	7, 36, 2, 2, 174, 64, 3, 2, 2, 2, 175, 177, 9, 7, 2, 2, 176, 175, 3, 2,
	2, 2, 177, 178, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2,
	179, 180, 3, 2, 2, 2, 180, 181, 8, 33, 2, 2, 181, 66, 3, 2, 2, 2, 9, 2,
	150, 156, 162, 168, 170, 178, 3, 8, 2, 2,
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
	"", "'=='", "'<'", "'>'", "'.'", "'='", "'['", "']'", "'('", "')'", "'/'",
	"'*'", "'-'", "'+'", "'->'", "':'", "'//'", "','", "'{'", "'}'", "'if'",
	"'then'", "'else'", "'let'", "'in'", "'pop'", "'drop'", "'len'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "BOOL", "ID", "NUM", "STR", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "BOOL", "ID", "NUM", "STR", "WS",
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
	GarvikLexerT__22 = 23
	GarvikLexerT__23 = 24
	GarvikLexerT__24 = 25
	GarvikLexerT__25 = 26
	GarvikLexerT__26 = 27
	GarvikLexerBOOL  = 28
	GarvikLexerID    = 29
	GarvikLexerNUM   = 30
	GarvikLexerSTR   = 31
	GarvikLexerWS    = 32
)
