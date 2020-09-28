// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 63, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 48, 10,
	10, 13, 10, 14, 10, 49, 3, 11, 6, 11, 53, 10, 11, 13, 11, 14, 11, 54, 3,
	12, 6, 12, 58, 10, 12, 13, 12, 14, 12, 59, 3, 12, 3, 12, 2, 2, 13, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	3, 2, 5, 3, 2, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	65, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3,
	2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33,
	3, 2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 42, 3, 2, 2, 2,
	19, 47, 3, 2, 2, 2, 21, 52, 3, 2, 2, 2, 23, 57, 3, 2, 2, 2, 25, 26, 7,
	42, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 43, 2, 2, 28, 6, 3, 2, 2, 2, 29,
	30, 7, 44, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 45, 2, 2, 32, 10, 3, 2,
	2, 2, 33, 34, 7, 110, 2, 2, 34, 35, 7, 103, 2, 2, 35, 36, 7, 118, 2, 2,
	36, 12, 3, 2, 2, 2, 37, 38, 7, 63, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7,
	107, 2, 2, 40, 41, 7, 112, 2, 2, 41, 16, 3, 2, 2, 2, 42, 43, 7, 103, 2,
	2, 43, 44, 7, 112, 2, 2, 44, 45, 7, 102, 2, 2, 45, 18, 3, 2, 2, 2, 46,
	48, 9, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3, 2, 2,
	2, 49, 50, 3, 2, 2, 2, 50, 20, 3, 2, 2, 2, 51, 53, 9, 3, 2, 2, 52, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2,
	55, 22, 3, 2, 2, 2, 56, 58, 9, 4, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61,
	62, 8, 12, 2, 2, 62, 24, 3, 2, 2, 2, 6, 2, 49, 54, 59, 3, 8, 2, 2,
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
	"", "'('", "')'", "'*'", "'+'", "'let'", "'='", "'in'", "'end'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ID", "NUM", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "ID", "NUM",
	"WS",
}

type SimplLexer struct {
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

func NewSimplLexer(input antlr.CharStream) *SimplLexer {

	l := new(SimplLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Simpl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimplLexer tokens.
const (
	SimplLexerT__0 = 1
	SimplLexerT__1 = 2
	SimplLexerT__2 = 3
	SimplLexerT__3 = 4
	SimplLexerT__4 = 5
	SimplLexerT__5 = 6
	SimplLexerT__6 = 7
	SimplLexerT__7 = 8
	SimplLexerID   = 9
	SimplLexerNUM  = 10
	SimplLexerWS   = 11
)
