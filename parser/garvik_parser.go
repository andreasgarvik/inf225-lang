// Code generated from Garvik.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Garvik

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 128,
	4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 7, 2, 8, 10, 2, 12, 2, 14, 2, 11, 11, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 22, 10, 3,
	12, 3, 14, 3, 25, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 35, 10, 3, 12, 3, 14, 3, 38, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 65, 10, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 123, 10, 3, 12, 3, 14,
	3, 126, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 2, 2, 156, 2, 9, 3, 2, 2, 2,
	4, 64, 3, 2, 2, 2, 6, 8, 5, 4, 3, 2, 7, 6, 3, 2, 2, 2, 8, 11, 3, 2, 2,
	2, 9, 7, 3, 2, 2, 2, 9, 10, 3, 2, 2, 2, 10, 3, 3, 2, 2, 2, 11, 9, 3, 2,
	2, 2, 12, 13, 8, 3, 1, 2, 13, 14, 7, 14, 2, 2, 14, 65, 5, 4, 3, 16, 15,
	16, 7, 18, 2, 2, 16, 65, 5, 4, 3, 15, 17, 18, 7, 8, 2, 2, 18, 23, 5, 4,
	3, 2, 19, 20, 7, 19, 2, 2, 20, 22, 5, 4, 3, 2, 21, 19, 3, 2, 2, 2, 22,
	25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2,
	2, 25, 23, 3, 2, 2, 2, 26, 27, 7, 9, 2, 2, 27, 65, 3, 2, 2, 2, 28, 29,
	7, 10, 2, 2, 29, 30, 5, 4, 3, 2, 30, 31, 7, 11, 2, 2, 31, 65, 3, 2, 2,
	2, 32, 36, 7, 20, 2, 2, 33, 35, 5, 4, 3, 2, 34, 33, 3, 2, 2, 2, 35, 38,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2, 2,
	38, 36, 3, 2, 2, 2, 39, 65, 7, 21, 2, 2, 40, 41, 7, 22, 2, 2, 41, 42, 5,
	4, 3, 2, 42, 43, 7, 23, 2, 2, 43, 44, 5, 4, 3, 2, 44, 45, 7, 24, 2, 2,
	45, 46, 5, 4, 3, 11, 46, 65, 3, 2, 2, 2, 47, 48, 7, 25, 2, 2, 48, 49, 5,
	4, 3, 2, 49, 50, 7, 7, 2, 2, 50, 51, 5, 4, 3, 2, 51, 52, 7, 26, 2, 2, 52,
	53, 5, 4, 3, 10, 53, 65, 3, 2, 2, 2, 54, 55, 7, 27, 2, 2, 55, 65, 5, 4,
	3, 9, 56, 57, 7, 28, 2, 2, 57, 65, 5, 4, 3, 8, 58, 59, 7, 29, 2, 2, 59,
	65, 5, 4, 3, 7, 60, 65, 7, 30, 2, 2, 61, 65, 7, 31, 2, 2, 62, 65, 7, 32,
	2, 2, 63, 65, 7, 33, 2, 2, 64, 12, 3, 2, 2, 2, 64, 15, 3, 2, 2, 2, 64,
	17, 3, 2, 2, 2, 64, 28, 3, 2, 2, 2, 64, 32, 3, 2, 2, 2, 64, 40, 3, 2, 2,
	2, 64, 47, 3, 2, 2, 2, 64, 54, 3, 2, 2, 2, 64, 56, 3, 2, 2, 2, 64, 58,
	3, 2, 2, 2, 64, 60, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 63, 3, 2, 2, 2, 65, 124, 3, 2, 2, 2, 66, 67, 12, 31, 2, 2, 67, 68,
	7, 3, 2, 2, 68, 123, 5, 4, 3, 32, 69, 70, 12, 30, 2, 2, 70, 71, 7, 4, 2,
	2, 71, 123, 5, 4, 3, 31, 72, 73, 12, 29, 2, 2, 73, 74, 7, 5, 2, 2, 74,
	123, 5, 4, 3, 30, 75, 76, 12, 28, 2, 2, 76, 77, 7, 6, 2, 2, 77, 78, 5,
	4, 3, 2, 78, 79, 7, 7, 2, 2, 79, 80, 5, 4, 3, 29, 80, 123, 3, 2, 2, 2,
	81, 82, 12, 27, 2, 2, 82, 83, 7, 6, 2, 2, 83, 123, 5, 4, 3, 28, 84, 85,
	12, 26, 2, 2, 85, 86, 7, 8, 2, 2, 86, 87, 5, 4, 3, 2, 87, 88, 7, 9, 2,
	2, 88, 89, 7, 7, 2, 2, 89, 90, 5, 4, 3, 27, 90, 123, 3, 2, 2, 2, 91, 92,
	12, 23, 2, 2, 92, 93, 7, 12, 2, 2, 93, 123, 5, 4, 3, 24, 94, 95, 12, 22,
	2, 2, 95, 96, 7, 13, 2, 2, 96, 123, 5, 4, 3, 23, 97, 98, 12, 21, 2, 2,
	98, 99, 7, 14, 2, 2, 99, 123, 5, 4, 3, 22, 100, 101, 12, 20, 2, 2, 101,
	102, 7, 15, 2, 2, 102, 123, 5, 4, 3, 21, 103, 104, 12, 19, 2, 2, 104, 105,
	7, 16, 2, 2, 105, 123, 5, 4, 3, 20, 106, 107, 12, 18, 2, 2, 107, 108, 7,
	17, 2, 2, 108, 123, 5, 4, 3, 19, 109, 110, 12, 17, 2, 2, 110, 111, 7, 7,
	2, 2, 111, 123, 5, 4, 3, 18, 112, 113, 12, 25, 2, 2, 113, 114, 7, 8, 2,
	2, 114, 115, 5, 4, 3, 2, 115, 116, 7, 9, 2, 2, 116, 123, 3, 2, 2, 2, 117,
	118, 12, 24, 2, 2, 118, 119, 7, 10, 2, 2, 119, 120, 5, 4, 3, 2, 120, 121,
	7, 11, 2, 2, 121, 123, 3, 2, 2, 2, 122, 66, 3, 2, 2, 2, 122, 69, 3, 2,
	2, 2, 122, 72, 3, 2, 2, 2, 122, 75, 3, 2, 2, 2, 122, 81, 3, 2, 2, 2, 122,
	84, 3, 2, 2, 2, 122, 91, 3, 2, 2, 2, 122, 94, 3, 2, 2, 2, 122, 97, 3, 2,
	2, 2, 122, 100, 3, 2, 2, 2, 122, 103, 3, 2, 2, 2, 122, 106, 3, 2, 2, 2,
	122, 109, 3, 2, 2, 2, 122, 112, 3, 2, 2, 2, 122, 117, 3, 2, 2, 2, 123,
	126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 5, 3,
	2, 2, 2, 126, 124, 3, 2, 2, 2, 8, 9, 23, 36, 64, 122, 124,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'=='", "'<'", "'>'", "'.'", "'='", "'['", "']'", "'('", "')'", "'/'",
	"'*'", "'-'", "'+'", "'->'", "':'", "'//'", "','", "'{'", "'}'", "'if'",
	"'then'", "'else'", "'let'", "'in'", "'pop'", "'drop'", "'len'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "BOOL", "ID", "NUM", "STR", "WS",
}

var ruleNames = []string{
	"program", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GarvikParser struct {
	*antlr.BaseParser
}

func NewGarvikParser(input antlr.TokenStream) *GarvikParser {
	this := new(GarvikParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Garvik.g4"

	return this
}

// GarvikParser tokens.
const (
	GarvikParserEOF   = antlr.TokenEOF
	GarvikParserT__0  = 1
	GarvikParserT__1  = 2
	GarvikParserT__2  = 3
	GarvikParserT__3  = 4
	GarvikParserT__4  = 5
	GarvikParserT__5  = 6
	GarvikParserT__6  = 7
	GarvikParserT__7  = 8
	GarvikParserT__8  = 9
	GarvikParserT__9  = 10
	GarvikParserT__10 = 11
	GarvikParserT__11 = 12
	GarvikParserT__12 = 13
	GarvikParserT__13 = 14
	GarvikParserT__14 = 15
	GarvikParserT__15 = 16
	GarvikParserT__16 = 17
	GarvikParserT__17 = 18
	GarvikParserT__18 = 19
	GarvikParserT__19 = 20
	GarvikParserT__20 = 21
	GarvikParserT__21 = 22
	GarvikParserT__22 = 23
	GarvikParserT__23 = 24
	GarvikParserT__24 = 25
	GarvikParserT__25 = 26
	GarvikParserT__26 = 27
	GarvikParserBOOL  = 28
	GarvikParserID    = 29
	GarvikParserNUM   = 30
	GarvikParserSTR   = 31
	GarvikParserWS    = 32
)

// GarvikParser rules.
const (
	GarvikParserRULE_program = 0
	GarvikParserRULE_expr    = 1
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GarvikParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GarvikParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ProgramContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GarvikParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GarvikParserRULE_program)
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GarvikParserT__5)|(1<<GarvikParserT__7)|(1<<GarvikParserT__11)|(1<<GarvikParserT__15)|(1<<GarvikParserT__17)|(1<<GarvikParserT__19)|(1<<GarvikParserT__22)|(1<<GarvikParserT__24)|(1<<GarvikParserT__25)|(1<<GarvikParserT__26)|(1<<GarvikParserBOOL)|(1<<GarvikParserID)|(1<<GarvikParserNUM)|(1<<GarvikParserSTR))) != 0 {
		{
			p.SetState(4)
			p.expr(0)
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GarvikParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GarvikParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrExprContext struct {
	*ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STR() antlr.TerminalNode {
	return s.GetToken(GarvikParserSTR, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DotExprContext struct {
	*ExprContext
	id    IExprContext
	field IExprContext
}

func NewDotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExprContext {
	var p = new(DotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotExprContext) GetId() IExprContext { return s.id }

func (s *DotExprContext) GetField() IExprContext { return s.field }

func (s *DotExprContext) SetId(v IExprContext) { s.id = v }

func (s *DotExprContext) SetField(v IExprContext) { s.field = v }

func (s *DotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DotExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterDotExpr(s)
	}
}

func (s *DotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitDotExpr(s)
	}
}

func (s *DotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitDotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenExprContext struct {
	*ExprContext
	id IExprContext
}

func NewLenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenExprContext {
	var p = new(LenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LenExprContext) GetId() IExprContext { return s.id }

func (s *LenExprContext) SetId(v IExprContext) { s.id = v }

func (s *LenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLenExpr(s)
	}
}

func (s *LenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLenExpr(s)
	}
}

func (s *LenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegativeExprContext struct {
	*ExprContext
}

func NewNegativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeExprContext {
	var p = new(NegativeExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NegativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterNegativeExpr(s)
	}
}

func (s *NegativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitNegativeExpr(s)
	}
}

func (s *NegativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitNegativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PopExprContext struct {
	*ExprContext
	id IExprContext
}

func NewPopExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PopExprContext {
	var p = new(PopExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PopExprContext) GetId() IExprContext { return s.id }

func (s *PopExprContext) SetId(v IExprContext) { s.id = v }

func (s *PopExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PopExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PopExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterPopExpr(s)
	}
}

func (s *PopExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitPopExpr(s)
	}
}

func (s *PopExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitPopExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubExprContext) GetLeft() IExprContext { return s.left }

func (s *SubExprContext) GetRight() IExprContext { return s.right }

func (s *SubExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *SubExprContext) SetRight(v IExprContext) { s.right = v }

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

func (s *SubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	*ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExprContext struct {
	*ExprContext
	id         IExprContext
	value      IExprContext
	expression IExprContext
}

func NewLetExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetExprContext {
	var p = new(LetExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LetExprContext) GetId() IExprContext { return s.id }

func (s *LetExprContext) GetValue() IExprContext { return s.value }

func (s *LetExprContext) GetExpression() IExprContext { return s.expression }

func (s *LetExprContext) SetId(v IExprContext) { s.id = v }

func (s *LetExprContext) SetValue(v IExprContext) { s.value = v }

func (s *LetExprContext) SetExpression(v IExprContext) { s.expression = v }

func (s *LetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LetExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLetExpr(s)
	}
}

func (s *LetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLetExpr(s)
	}
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterExprContext {
	var p = new(GreaterExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterExprContext) GetLeft() IExprContext { return s.left }

func (s *GreaterExprContext) GetRight() IExprContext { return s.right }

func (s *GreaterExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterExprContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterGreaterExpr(s)
	}
}

func (s *GreaterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitGreaterExpr(s)
	}
}

func (s *GreaterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitGreaterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExprContext struct {
	*ExprContext
	param IExprContext
	body  IExprContext
}

func NewLambdaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExprContext {
	var p = new(LambdaExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LambdaExprContext) GetParam() IExprContext { return s.param }

func (s *LambdaExprContext) GetBody() IExprContext { return s.body }

func (s *LambdaExprContext) SetParam(v IExprContext) { s.param = v }

func (s *LambdaExprContext) SetBody(v IExprContext) { s.body = v }

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LambdaExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLambdaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructExprContext struct {
	*ExprContext
}

func NewStructExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExprContext {
	var p = new(StructExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StructExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterStructExpr(s)
	}
}

func (s *StructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitStructExpr(s)
	}
}

func (s *StructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitStructExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivExprContext) GetLeft() IExprContext { return s.left }

func (s *DivExprContext) GetRight() IExprContext { return s.right }

func (s *DivExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *DivExprContext) SetRight(v IExprContext) { s.right = v }

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

func (s *DivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LookupAssignExprContext struct {
	*ExprContext
	id    IExprContext
	key   IExprContext
	value IExprContext
}

func NewLookupAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LookupAssignExprContext {
	var p = new(LookupAssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LookupAssignExprContext) GetId() IExprContext { return s.id }

func (s *LookupAssignExprContext) GetKey() IExprContext { return s.key }

func (s *LookupAssignExprContext) GetValue() IExprContext { return s.value }

func (s *LookupAssignExprContext) SetId(v IExprContext) { s.id = v }

func (s *LookupAssignExprContext) SetKey(v IExprContext) { s.key = v }

func (s *LookupAssignExprContext) SetValue(v IExprContext) { s.value = v }

func (s *LookupAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupAssignExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LookupAssignExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LookupAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLookupAssignExpr(s)
	}
}

func (s *LookupAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLookupAssignExpr(s)
	}
}

func (s *LookupAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLookupAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	*ExprContext
	fun IExprContext
	arg IExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetFun() IExprContext { return s.fun }

func (s *CallExprContext) GetArg() IExprContext { return s.arg }

func (s *CallExprContext) SetFun(v IExprContext) { s.fun = v }

func (s *CallExprContext) SetArg(v IExprContext) { s.arg = v }

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CallExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListExprContext struct {
	*ExprContext
}

func NewListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExprContext {
	var p = new(ListExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ListExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterListExpr(s)
	}
}

func (s *ListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitListExpr(s)
	}
}

func (s *ListExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitListExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	*ExprContext
	id    IExprContext
	value IExprContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignExprContext) GetId() IExprContext { return s.id }

func (s *AssignExprContext) GetValue() IExprContext { return s.value }

func (s *AssignExprContext) SetId(v IExprContext) { s.id = v }

func (s *AssignExprContext) SetValue(v IExprContext) { s.value = v }

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AssignExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExprContext struct {
	*ExprContext
	con IExprContext
	t   IExprContext
	f   IExprContext
}

func NewIfElseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExprContext {
	var p = new(IfElseExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfElseExprContext) GetCon() IExprContext { return s.con }

func (s *IfElseExprContext) GetT() IExprContext { return s.t }

func (s *IfElseExprContext) GetF() IExprContext { return s.f }

func (s *IfElseExprContext) SetCon(v IExprContext) { s.con = v }

func (s *IfElseExprContext) SetT(v IExprContext) { s.t = v }

func (s *IfElseExprContext) SetF(v IExprContext) { s.f = v }

func (s *IfElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfElseExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterIfElseExpr(s)
	}
}

func (s *IfElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitIfElseExpr(s)
	}
}

func (s *IfElseExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitIfElseExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessExprContext {
	var p = new(LessExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessExprContext) GetLeft() IExprContext { return s.left }

func (s *LessExprContext) GetRight() IExprContext { return s.right }

func (s *LessExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessExprContext) SetRight(v IExprContext) { s.right = v }

func (s *LessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLessExpr(s)
	}
}

func (s *LessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLessExpr(s)
	}
}

func (s *LessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DotAssignExprContext struct {
	*ExprContext
	id    IExprContext
	field IExprContext
	value IExprContext
}

func NewDotAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotAssignExprContext {
	var p = new(DotAssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotAssignExprContext) GetId() IExprContext { return s.id }

func (s *DotAssignExprContext) GetField() IExprContext { return s.field }

func (s *DotAssignExprContext) GetValue() IExprContext { return s.value }

func (s *DotAssignExprContext) SetId(v IExprContext) { s.id = v }

func (s *DotAssignExprContext) SetField(v IExprContext) { s.field = v }

func (s *DotAssignExprContext) SetValue(v IExprContext) { s.value = v }

func (s *DotAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotAssignExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DotAssignExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterDotAssignExpr(s)
	}
}

func (s *DotAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitDotAssignExpr(s)
	}
}

func (s *DotAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitDotAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewMultExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExprContext {
	var p = new(MultExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultExprContext) GetLeft() IExprContext { return s.left }

func (s *MultExprContext) GetRight() IExprContext { return s.right }

func (s *MultExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultExprContext) SetRight(v IExprContext) { s.right = v }

func (s *MultExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterMultExpr(s)
	}
}

func (s *MultExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitMultExpr(s)
	}
}

func (s *MultExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitMultExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldExprContext struct {
	*ExprContext
	id    IExprContext
	value IExprContext
}

func NewFieldExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldExprContext {
	var p = new(FieldExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FieldExprContext) GetId() IExprContext { return s.id }

func (s *FieldExprContext) GetValue() IExprContext { return s.value }

func (s *FieldExprContext) SetId(v IExprContext) { s.id = v }

func (s *FieldExprContext) SetValue(v IExprContext) { s.value = v }

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FieldExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterFieldExpr(s)
	}
}

func (s *FieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitFieldExpr(s)
	}
}

func (s *FieldExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitFieldExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommentExprContext struct {
	*ExprContext
}

func NewCommentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentExprContext {
	var p = new(CommentExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CommentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CommentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterCommentExpr(s)
	}
}

func (s *CommentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitCommentExpr(s)
	}
}

func (s *CommentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitCommentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	*ExprContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) NUM() antlr.TerminalNode {
	return s.GetToken(GarvikParserNUM, 0)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DropExprContext struct {
	*ExprContext
	id IExprContext
}

func NewDropExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DropExprContext {
	var p = new(DropExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DropExprContext) GetId() IExprContext { return s.id }

func (s *DropExprContext) SetId(v IExprContext) { s.id = v }

func (s *DropExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DropExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterDropExpr(s)
	}
}

func (s *DropExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitDropExpr(s)
	}
}

func (s *DropExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitDropExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetLeft() IExprContext { return s.left }

func (s *AddExprContext) GetRight() IExprContext { return s.right }

func (s *AddExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddExprContext) SetRight(v IExprContext) { s.right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LookupExprContext struct {
	*ExprContext
	id  IExprContext
	key IExprContext
}

func NewLookupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LookupExprContext {
	var p = new(LookupExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LookupExprContext) GetId() IExprContext { return s.id }

func (s *LookupExprContext) GetKey() IExprContext { return s.key }

func (s *LookupExprContext) SetId(v IExprContext) { s.id = v }

func (s *LookupExprContext) SetKey(v IExprContext) { s.key = v }

func (s *LookupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LookupExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LookupExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterLookupExpr(s)
	}
}

func (s *LookupExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitLookupExpr(s)
	}
}

func (s *LookupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitLookupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExprContext struct {
	*ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GarvikParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	*ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GarvikParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualExprContext) GetLeft() IExprContext { return s.left }

func (s *EqualExprContext) GetRight() IExprContext { return s.right }

func (s *EqualExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualExprContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.EnterEqualExpr(s)
	}
}

func (s *EqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GarvikListener); ok {
		listenerT.ExitEqualExpr(s)
	}
}

func (s *EqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GarvikVisitor:
		return t.VisitEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GarvikParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GarvikParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GarvikParserRULE_expr, _p)
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
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GarvikParserT__11:
		localctx = NewNegativeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)
			p.Match(GarvikParserT__11)
		}
		{
			p.SetState(12)
			p.expr(14)
		}

	case GarvikParserT__15:
		localctx = NewCommentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(GarvikParserT__15)
		}
		{
			p.SetState(14)
			p.expr(13)
		}

	case GarvikParserT__5:
		localctx = NewListExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(GarvikParserT__5)
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GarvikParserT__16 {
			{
				p.SetState(17)
				p.Match(GarvikParserT__16)
			}
			{
				p.SetState(18)
				p.expr(0)
			}

			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(24)
			p.Match(GarvikParserT__6)
		}

	case GarvikParserT__7:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(GarvikParserT__7)
		}
		{
			p.SetState(27)
			p.expr(0)
		}
		{
			p.SetState(28)
			p.Match(GarvikParserT__8)
		}

	case GarvikParserT__17:
		localctx = NewStructExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(GarvikParserT__17)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GarvikParserT__5)|(1<<GarvikParserT__7)|(1<<GarvikParserT__11)|(1<<GarvikParserT__15)|(1<<GarvikParserT__17)|(1<<GarvikParserT__19)|(1<<GarvikParserT__22)|(1<<GarvikParserT__24)|(1<<GarvikParserT__25)|(1<<GarvikParserT__26)|(1<<GarvikParserBOOL)|(1<<GarvikParserID)|(1<<GarvikParserNUM)|(1<<GarvikParserSTR))) != 0 {
			{
				p.SetState(31)
				p.expr(0)
			}

			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(37)
			p.Match(GarvikParserT__18)
		}

	case GarvikParserT__19:
		localctx = NewIfElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(GarvikParserT__19)
		}
		{
			p.SetState(39)

			var _x = p.expr(0)

			localctx.(*IfElseExprContext).con = _x
		}
		{
			p.SetState(40)
			p.Match(GarvikParserT__20)
		}
		{
			p.SetState(41)

			var _x = p.expr(0)

			localctx.(*IfElseExprContext).t = _x
		}
		{
			p.SetState(42)
			p.Match(GarvikParserT__21)
		}
		{
			p.SetState(43)

			var _x = p.expr(9)

			localctx.(*IfElseExprContext).f = _x
		}

	case GarvikParserT__22:
		localctx = NewLetExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(GarvikParserT__22)
		}
		{
			p.SetState(46)

			var _x = p.expr(0)

			localctx.(*LetExprContext).id = _x
		}
		{
			p.SetState(47)
			p.Match(GarvikParserT__4)
		}
		{
			p.SetState(48)

			var _x = p.expr(0)

			localctx.(*LetExprContext).value = _x
		}
		{
			p.SetState(49)
			p.Match(GarvikParserT__23)
		}
		{
			p.SetState(50)

			var _x = p.expr(8)

			localctx.(*LetExprContext).expression = _x
		}

	case GarvikParserT__24:
		localctx = NewPopExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(GarvikParserT__24)
		}
		{
			p.SetState(53)

			var _x = p.expr(7)

			localctx.(*PopExprContext).id = _x
		}

	case GarvikParserT__25:
		localctx = NewDropExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(GarvikParserT__25)
		}
		{
			p.SetState(55)

			var _x = p.expr(6)

			localctx.(*DropExprContext).id = _x
		}

	case GarvikParserT__26:
		localctx = NewLenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.Match(GarvikParserT__26)
		}
		{
			p.SetState(57)

			var _x = p.expr(5)

			localctx.(*LenExprContext).id = _x
		}

	case GarvikParserBOOL:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(GarvikParserBOOL)
		}

	case GarvikParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(GarvikParserID)
		}

	case GarvikParserNUM:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(GarvikParserNUM)
		}

	case GarvikParserSTR:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(GarvikParserSTR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
				}
				{
					p.SetState(65)
					p.Match(GarvikParserT__0)
				}
				{
					p.SetState(66)

					var _x = p.expr(30)

					localctx.(*EqualExprContext).right = _x
				}

			case 2:
				localctx = NewLessExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
				}
				{
					p.SetState(68)
					p.Match(GarvikParserT__1)
				}
				{
					p.SetState(69)

					var _x = p.expr(29)

					localctx.(*LessExprContext).right = _x
				}

			case 3:
				localctx = NewGreaterExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
				}
				{
					p.SetState(71)
					p.Match(GarvikParserT__2)
				}
				{
					p.SetState(72)

					var _x = p.expr(28)

					localctx.(*GreaterExprContext).right = _x
				}

			case 4:
				localctx = NewDotAssignExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotAssignExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(74)
					p.Match(GarvikParserT__3)
				}
				{
					p.SetState(75)

					var _x = p.expr(0)

					localctx.(*DotAssignExprContext).field = _x
				}
				{
					p.SetState(76)
					p.Match(GarvikParserT__4)
				}
				{
					p.SetState(77)

					var _x = p.expr(27)

					localctx.(*DotAssignExprContext).value = _x
				}

			case 5:
				localctx = NewDotExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(80)
					p.Match(GarvikParserT__3)
				}
				{
					p.SetState(81)

					var _x = p.expr(26)

					localctx.(*DotExprContext).field = _x
				}

			case 6:
				localctx = NewLookupAssignExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LookupAssignExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
				}
				{
					p.SetState(83)
					p.Match(GarvikParserT__5)
				}
				{
					p.SetState(84)

					var _x = p.expr(0)

					localctx.(*LookupAssignExprContext).key = _x
				}
				{
					p.SetState(85)
					p.Match(GarvikParserT__6)
				}
				{
					p.SetState(86)
					p.Match(GarvikParserT__4)
				}
				{
					p.SetState(87)

					var _x = p.expr(25)

					localctx.(*LookupAssignExprContext).value = _x
				}

			case 7:
				localctx = NewDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(90)
					p.Match(GarvikParserT__9)
				}
				{
					p.SetState(91)

					var _x = p.expr(22)

					localctx.(*DivExprContext).right = _x
				}

			case 8:
				localctx = NewMultExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(93)
					p.Match(GarvikParserT__10)
				}
				{
					p.SetState(94)

					var _x = p.expr(21)

					localctx.(*MultExprContext).right = _x
				}

			case 9:
				localctx = NewSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(96)
					p.Match(GarvikParserT__11)
				}
				{
					p.SetState(97)

					var _x = p.expr(20)

					localctx.(*SubExprContext).right = _x
				}

			case 10:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(99)
					p.Match(GarvikParserT__12)
				}
				{
					p.SetState(100)

					var _x = p.expr(19)

					localctx.(*AddExprContext).right = _x
				}

			case 11:
				localctx = NewLambdaExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LambdaExprContext).param = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(102)
					p.Match(GarvikParserT__13)
				}
				{
					p.SetState(103)

					var _x = p.expr(18)

					localctx.(*LambdaExprContext).body = _x
				}

			case 12:
				localctx = NewFieldExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*FieldExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(105)
					p.Match(GarvikParserT__14)
				}
				{
					p.SetState(106)

					var _x = p.expr(17)

					localctx.(*FieldExprContext).value = _x
				}

			case 13:
				localctx = NewAssignExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AssignExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(108)
					p.Match(GarvikParserT__4)
				}
				{
					p.SetState(109)

					var _x = p.expr(16)

					localctx.(*AssignExprContext).value = _x
				}

			case 14:
				localctx = NewLookupExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LookupExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(111)
					p.Match(GarvikParserT__5)
				}
				{
					p.SetState(112)

					var _x = p.expr(0)

					localctx.(*LookupExprContext).key = _x
				}
				{
					p.SetState(113)
					p.Match(GarvikParserT__6)
				}

			case 15:
				localctx = NewCallExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallExprContext).fun = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GarvikParserRULE_expr)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(116)
					p.Match(GarvikParserT__7)
				}
				{
					p.SetState(117)

					var _x = p.expr(0)

					localctx.(*CallExprContext).arg = _x
				}
				{
					p.SetState(118)
					p.Match(GarvikParserT__8)
				}

			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GarvikParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GarvikParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 22)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
