// Code generated from ScienceRule.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // ScienceRule

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 78, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 36, 10,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 43, 10, 2, 12, 2, 14, 2, 46, 11,
	2, 3, 3, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 60, 10, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5, 5, 5, 66, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 76, 10, 7, 3, 7,
	2, 3, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 2, 9, 13, 2, 84, 2, 35, 3, 2,
	2, 2, 4, 47, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 65, 3, 2, 2, 2, 10, 67,
	3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14, 16, 8, 2, 1, 2, 15, 17, 7, 14, 2, 2,
	16, 15, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 20, 7,
	26, 2, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21,
	22, 7, 3, 2, 2, 22, 23, 5, 2, 2, 2, 23, 24, 7, 4, 2, 2, 24, 36, 3, 2, 2,
	2, 25, 26, 5, 4, 3, 2, 26, 27, 7, 26, 2, 2, 27, 28, 7, 5, 2, 2, 28, 36,
	3, 2, 2, 2, 29, 30, 5, 4, 3, 2, 30, 31, 7, 26, 2, 2, 31, 32, 9, 2, 2, 2,
	32, 33, 7, 26, 2, 2, 33, 34, 5, 8, 5, 2, 34, 36, 3, 2, 2, 2, 35, 14, 3,
	2, 2, 2, 35, 25, 3, 2, 2, 2, 35, 29, 3, 2, 2, 2, 36, 44, 3, 2, 2, 2, 37,
	38, 12, 5, 2, 2, 38, 39, 7, 26, 2, 2, 39, 40, 7, 15, 2, 2, 40, 41, 7, 26,
	2, 2, 41, 43, 5, 2, 2, 6, 42, 37, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42,
	3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2,
	47, 49, 7, 18, 2, 2, 48, 50, 5, 6, 4, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3,
	2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 6, 2, 2, 52, 53, 5, 4, 3, 2, 53,
	7, 3, 2, 2, 2, 54, 66, 7, 16, 2, 2, 55, 66, 7, 17, 2, 2, 56, 66, 7, 20,
	2, 2, 57, 66, 7, 21, 2, 2, 58, 60, 7, 10, 2, 2, 59, 58, 3, 2, 2, 2, 59,
	60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63, 7, 22, 2, 2, 62, 64, 7, 23,
	2, 2, 63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 54,
	3, 2, 2, 2, 65, 55, 3, 2, 2, 2, 65, 56, 3, 2, 2, 2, 65, 57, 3, 2, 2, 2,
	65, 59, 3, 2, 2, 2, 66, 9, 3, 2, 2, 2, 67, 68, 7, 7, 2, 2, 68, 69, 5, 12,
	7, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 20, 2, 2, 71, 72, 7, 25, 2, 2, 72,
	76, 5, 12, 7, 2, 73, 74, 7, 20, 2, 2, 74, 76, 7, 8, 2, 2, 75, 70, 3, 2,
	2, 2, 75, 73, 3, 2, 2, 2, 76, 13, 3, 2, 2, 2, 11, 16, 19, 35, 44, 49, 59,
	63, 65, 75,
}
var literalNames = []string{
	"", "'('", "')'", "'pr'", "'.'", "'['", "']'", "'+'", "'-'", "'*'", "'/'",
	"'%'", "", "", "", "'null'", "", "", "", "", "", "", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "ADD", "SUB", "MUL", "DIV", "MOD", "NOT", "LOGICAL_OPERATOR",
	"BOOLEAN", "NULL", "ATTRNAME", "VERSION", "STRING", "DOUBLE", "INT", "EXP",
	"NEWLINE", "COMMA", "SP",
}

var ruleNames = []string{
	"sciencerule", "attrPath", "subAttr", "value", "listStrings", "subListOfStrings",
}

type ScienceRuleParser struct {
	*antlr.BaseParser
}

// NewScienceRuleParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ScienceRuleParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewScienceRuleParser(input antlr.TokenStream) *ScienceRuleParser {
	this := new(ScienceRuleParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ScienceRule.g4"

	return this
}

// ScienceRuleParser tokens.
const (
	ScienceRuleParserEOF              = antlr.TokenEOF
	ScienceRuleParserT__0             = 1
	ScienceRuleParserT__1             = 2
	ScienceRuleParserT__2             = 3
	ScienceRuleParserT__3             = 4
	ScienceRuleParserT__4             = 5
	ScienceRuleParserT__5             = 6
	ScienceRuleParserADD              = 7
	ScienceRuleParserSUB              = 8
	ScienceRuleParserMUL              = 9
	ScienceRuleParserDIV              = 10
	ScienceRuleParserMOD              = 11
	ScienceRuleParserNOT              = 12
	ScienceRuleParserLOGICAL_OPERATOR = 13
	ScienceRuleParserBOOLEAN          = 14
	ScienceRuleParserNULL             = 15
	ScienceRuleParserATTRNAME         = 16
	ScienceRuleParserVERSION          = 17
	ScienceRuleParserSTRING           = 18
	ScienceRuleParserDOUBLE           = 19
	ScienceRuleParserINT              = 20
	ScienceRuleParserEXP              = 21
	ScienceRuleParserNEWLINE          = 22
	ScienceRuleParserCOMMA            = 23
	ScienceRuleParserSP               = 24
)

// ScienceRuleParser rules.
const (
	ScienceRuleParserRULE_sciencerule      = 0
	ScienceRuleParserRULE_attrPath         = 1
	ScienceRuleParserRULE_subAttr          = 2
	ScienceRuleParserRULE_value            = 3
	ScienceRuleParserRULE_listStrings      = 4
	ScienceRuleParserRULE_subListOfStrings = 5
)

// IScienceruleContext is an interface to support dynamic dispatch.
type IScienceruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScienceruleContext differentiates from other interfaces.
	IsScienceruleContext()
}

type ScienceruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScienceruleContext() *ScienceruleContext {
	var p = new(ScienceruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_sciencerule
	return p
}

func (*ScienceruleContext) IsScienceruleContext() {}

func NewScienceruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScienceruleContext {
	var p = new(ScienceruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_sciencerule

	return p
}

func (s *ScienceruleContext) GetParser() antlr.Parser { return s.parser }

func (s *ScienceruleContext) CopyFrom(ctx *ScienceruleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ScienceruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScienceruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	*ScienceruleContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.ScienceruleContext = NewEmptyScienceruleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ScienceruleContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(ScienceRuleParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) ADD() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserADD, 0)
}

func (s *CompareExpContext) SUB() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSUB, 0)
}

func (s *CompareExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserMUL, 0)
}

func (s *CompareExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserDIV, 0)
}

func (s *CompareExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserMOD, 0)
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	*ScienceruleContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	p.ScienceruleContext = NewEmptyScienceruleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ScienceruleContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Sciencerule() IScienceruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScienceruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScienceruleContext)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserNOT, 0)
}

func (s *ParenExpContext) SP() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSP, 0)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	*ScienceruleContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	p.ScienceruleContext = NewEmptyScienceruleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ScienceruleContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSP, 0)
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpContext struct {
	*ScienceruleContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	p.ScienceruleContext = NewEmptyScienceruleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ScienceruleContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllSciencerule() []IScienceruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScienceruleContext)(nil)).Elem())
	var tst = make([]IScienceruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScienceruleContext)
		}
	}

	return tst
}

func (s *LogicalExpContext) Sciencerule(i int) IScienceruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScienceruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScienceruleContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(ScienceRuleParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) Sciencerule() (localctx IScienceruleContext) {
	return p.sciencerule(0)
}

func (p *ScienceRuleParser) sciencerule(_p int) (localctx IScienceruleContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewScienceruleContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScienceruleContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ScienceRuleParserRULE_sciencerule, _p)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScienceRuleParserNOT {
			{
				p.SetState(13)
				p.Match(ScienceRuleParserNOT)
			}

		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScienceRuleParserSP {
			{
				p.SetState(16)
				p.Match(ScienceRuleParserSP)
			}

		}
		{
			p.SetState(19)
			p.Match(ScienceRuleParserT__0)
		}
		{
			p.SetState(20)
			p.sciencerule(0)
		}
		{
			p.SetState(21)
			p.Match(ScienceRuleParserT__1)
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.AttrPath()
		}
		{
			p.SetState(24)
			p.Match(ScienceRuleParserSP)
		}
		{
			p.SetState(25)
			p.Match(ScienceRuleParserT__2)
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.AttrPath()
		}
		{
			p.SetState(28)
			p.Match(ScienceRuleParserSP)
		}
		{
			p.SetState(29)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScienceRuleParserADD)|(1<<ScienceRuleParserSUB)|(1<<ScienceRuleParserMUL)|(1<<ScienceRuleParserDIV)|(1<<ScienceRuleParserMOD))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(30)
			p.Match(ScienceRuleParserSP)
		}
		{
			p.SetState(31)
			p.Value()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewScienceruleContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ScienceRuleParserRULE_sciencerule)
			p.SetState(35)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(36)
				p.Match(ScienceRuleParserSP)
			}
			{
				p.SetState(37)
				p.Match(ScienceRuleParserLOGICAL_OPERATOR)
			}
			{
				p.SetState(38)
				p.Match(ScienceRuleParserSP)
			}
			{
				p.SetState(39)
				p.sciencerule(4)
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_attrPath
	return p
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ScienceRuleParserRULE_attrPath)
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
		p.SetState(45)
		p.Match(ScienceRuleParserATTRNAME)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScienceRuleParserT__3 {
		{
			p.SetState(46)
			p.SubAttr()
		}

	}

	return localctx
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_subAttr
	return p
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScienceRuleParserRULE_subAttr)

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
		p.SetState(49)
		p.Match(ScienceRuleParserT__3)
	}
	{
		p.SetState(50)
		p.AttrPath()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanContext struct {
	*ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserBOOLEAN, 0)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserNULL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	*ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type LongContext struct {
	*ValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserINT, 0)
}

func (s *LongContext) SUB() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSUB, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserEXP, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ScienceRuleParserRULE_value)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScienceRuleParserBOOLEAN:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(ScienceRuleParserBOOLEAN)
		}

	case ScienceRuleParserNULL:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(ScienceRuleParserNULL)
		}

	case ScienceRuleParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(ScienceRuleParserSTRING)
		}

	case ScienceRuleParserDOUBLE:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Match(ScienceRuleParserDOUBLE)
		}

	case ScienceRuleParserSUB, ScienceRuleParserINT:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScienceRuleParserSUB {
			{
				p.SetState(56)
				p.Match(ScienceRuleParserSUB)
			}

		}
		{
			p.SetState(59)
			p.Match(ScienceRuleParserINT)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(60)
				p.Match(ScienceRuleParserEXP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_listStrings
	return p
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ScienceRuleParserRULE_listStrings)

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
		p.SetState(65)
		p.Match(ScienceRuleParserT__4)
	}
	{
		p.SetState(66)
		p.SubListOfStrings()
	}

	return localctx
}

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScienceRuleParserRULE_subListOfStrings
	return p
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScienceRuleParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ScienceRuleParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ScienceRuleVisitor:
		return t.VisitSubListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ScienceRuleParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ScienceRuleParserRULE_subListOfStrings)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(ScienceRuleParserSTRING)
		}
		{
			p.SetState(69)
			p.Match(ScienceRuleParserCOMMA)
		}
		{
			p.SetState(70)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(ScienceRuleParserSTRING)
		}
		{
			p.SetState(72)
			p.Match(ScienceRuleParserT__5)
		}

	}

	return localctx
}

func (p *ScienceRuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ScienceruleContext = nil
		if localctx != nil {
			t = localctx.(*ScienceruleContext)
		}
		return p.Sciencerule_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ScienceRuleParser) Sciencerule_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
