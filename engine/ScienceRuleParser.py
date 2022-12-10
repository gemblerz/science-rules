# Generated from ScienceRule.g4 by ANTLR 4.9
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO


def serializedATN():
    with StringIO() as buf:
        buf.write("\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\35")
        buf.write("@\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\3\2\3\2\3\2\3\2\3\2")
        buf.write("\3\2\3\2\3\2\3\2\3\2\5\2\25\n\2\3\2\3\2\3\2\3\2\3\2\3")
        buf.write("\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\7\2&\n\2\f\2\16")
        buf.write("\2)\13\2\3\3\3\3\3\3\3\3\3\3\5\3\60\n\3\3\3\3\3\5\3\64")
        buf.write("\n\3\3\3\5\3\67\n\3\3\4\3\4\5\4;\n\4\3\5\3\5\3\5\3\5\2")
        buf.write("\3\2\6\2\4\6\b\2\4\3\2\7\13\3\2\f\22\2H\2\24\3\2\2\2\4")
        buf.write("\66\3\2\2\2\68\3\2\2\2\b<\3\2\2\2\n\13\b\2\1\2\13\f\7")
        buf.write("\3\2\2\f\r\5\2\2\2\r\16\7\4\2\2\16\25\3\2\2\2\17\20\5")
        buf.write("\6\4\2\20\21\7\35\2\2\21\22\7\5\2\2\22\25\3\2\2\2\23\25")
        buf.write("\5\4\3\2\24\n\3\2\2\2\24\17\3\2\2\2\24\23\3\2\2\2\25\'")
        buf.write("\3\2\2\2\26\27\f\7\2\2\27\30\7\35\2\2\30\31\7\23\2\2\31")
        buf.write("\32\7\35\2\2\32&\5\2\2\b\33\34\f\5\2\2\34\35\7\35\2\2")
        buf.write("\35\36\t\2\2\2\36\37\7\35\2\2\37&\5\2\2\6 !\f\4\2\2!\"")
        buf.write("\7\35\2\2\"#\t\3\2\2#$\7\35\2\2$&\5\2\2\5%\26\3\2\2\2")
        buf.write("%\33\3\2\2\2% \3\2\2\2&)\3\2\2\2\'%\3\2\2\2\'(\3\2\2\2")
        buf.write("(\3\3\2\2\2)\'\3\2\2\2*\67\7\24\2\2+\67\7\25\2\2,\67\7")
        buf.write("\27\2\2-\67\7\30\2\2.\60\7\b\2\2/.\3\2\2\2/\60\3\2\2\2")
        buf.write("\60\61\3\2\2\2\61\63\7\31\2\2\62\64\7\32\2\2\63\62\3\2")
        buf.write("\2\2\63\64\3\2\2\2\64\67\3\2\2\2\65\67\5\6\4\2\66*\3\2")
        buf.write("\2\2\66+\3\2\2\2\66,\3\2\2\2\66-\3\2\2\2\66/\3\2\2\2\66")
        buf.write("\65\3\2\2\2\67\5\3\2\2\28:\7\26\2\29;\5\b\5\2:9\3\2\2")
        buf.write("\2:;\3\2\2\2;\7\3\2\2\2<=\7\6\2\2=>\5\6\4\2>\t\3\2\2\2")
        buf.write("\t\24%\'/\63\66:")
        return buf.getvalue()


class ScienceRuleParser ( Parser ):

    grammarFileName = "ScienceRule.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'('", "')'", "'pr'", "'.'", "'+'", "'-'", 
                     "'*'", "'/'", "'%'", "<INVALID>", "<INVALID>", "<INVALID>", 
                     "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                     "<INVALID>", "<INVALID>", "'null'", "<INVALID>", "<INVALID>", 
                     "<INVALID>", "<INVALID>", "<INVALID>", "'\n'" ]

    symbolicNames = [ "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "ADD", "SUB", "MUL", "DIV", "MOD", "EQ", 
                      "NE", "GT", "LT", "GE", "LE", "CO", "LOGICAL_OPERATOR", 
                      "BOOLEAN", "NULL", "ATTRNAME", "STRING", "DOUBLE", 
                      "INT", "EXP", "NEWLINE", "COMMA", "SP" ]

    RULE_sciencerule = 0
    RULE_value = 1
    RULE_attrPath = 2
    RULE_subAttr = 3

    ruleNames =  [ "sciencerule", "value", "attrPath", "subAttr" ]

    EOF = Token.EOF
    T__0=1
    T__1=2
    T__2=3
    T__3=4
    ADD=5
    SUB=6
    MUL=7
    DIV=8
    MOD=9
    EQ=10
    NE=11
    GT=12
    LT=13
    GE=14
    LE=15
    CO=16
    LOGICAL_OPERATOR=17
    BOOLEAN=18
    NULL=19
    ATTRNAME=20
    STRING=21
    DOUBLE=22
    INT=23
    EXP=24
    NEWLINE=25
    COMMA=26
    SP=27

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.9")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class ScienceruleContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return ScienceRuleParser.RULE_sciencerule

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)


    class CompareExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.lhs = None # ScienceruleContext
            self.op = None # Token
            self.rhs = None # ScienceruleContext
            self.copyFrom(ctx)

        def SP(self, i:int=None):
            if i is None:
                return self.getTokens(ScienceRuleParser.SP)
            else:
                return self.getToken(ScienceRuleParser.SP, i)
        def sciencerule(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ScienceRuleParser.ScienceruleContext)
            else:
                return self.getTypedRuleContext(ScienceRuleParser.ScienceruleContext,i)

        def EQ(self):
            return self.getToken(ScienceRuleParser.EQ, 0)
        def NE(self):
            return self.getToken(ScienceRuleParser.NE, 0)
        def GT(self):
            return self.getToken(ScienceRuleParser.GT, 0)
        def LT(self):
            return self.getToken(ScienceRuleParser.LT, 0)
        def GE(self):
            return self.getToken(ScienceRuleParser.GE, 0)
        def LE(self):
            return self.getToken(ScienceRuleParser.LE, 0)
        def CO(self):
            return self.getToken(ScienceRuleParser.CO, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCompareExp" ):
                listener.enterCompareExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCompareExp" ):
                listener.exitCompareExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCompareExp" ):
                return visitor.visitCompareExp(self)
            else:
                return visitor.visitChildren(self)


    class ArithmeticExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.lhs = None # ScienceruleContext
            self.op = None # Token
            self.rhs = None # ScienceruleContext
            self.copyFrom(ctx)

        def SP(self, i:int=None):
            if i is None:
                return self.getTokens(ScienceRuleParser.SP)
            else:
                return self.getToken(ScienceRuleParser.SP, i)
        def sciencerule(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ScienceRuleParser.ScienceruleContext)
            else:
                return self.getTypedRuleContext(ScienceRuleParser.ScienceruleContext,i)

        def ADD(self):
            return self.getToken(ScienceRuleParser.ADD, 0)
        def SUB(self):
            return self.getToken(ScienceRuleParser.SUB, 0)
        def MUL(self):
            return self.getToken(ScienceRuleParser.MUL, 0)
        def DIV(self):
            return self.getToken(ScienceRuleParser.DIV, 0)
        def MOD(self):
            return self.getToken(ScienceRuleParser.MOD, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterArithmeticExp" ):
                listener.enterArithmeticExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitArithmeticExp" ):
                listener.exitArithmeticExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitArithmeticExp" ):
                return visitor.visitArithmeticExp(self)
            else:
                return visitor.visitChildren(self)


    class ParenExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def sciencerule(self):
            return self.getTypedRuleContext(ScienceRuleParser.ScienceruleContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterParenExp" ):
                listener.enterParenExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitParenExp" ):
                listener.exitParenExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitParenExp" ):
                return visitor.visitParenExp(self)
            else:
                return visitor.visitChildren(self)


    class PresentExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def attrPath(self):
            return self.getTypedRuleContext(ScienceRuleParser.AttrPathContext,0)

        def SP(self):
            return self.getToken(ScienceRuleParser.SP, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPresentExp" ):
                listener.enterPresentExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPresentExp" ):
                listener.exitPresentExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitPresentExp" ):
                return visitor.visitPresentExp(self)
            else:
                return visitor.visitChildren(self)


    class LogicalExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def sciencerule(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(ScienceRuleParser.ScienceruleContext)
            else:
                return self.getTypedRuleContext(ScienceRuleParser.ScienceruleContext,i)

        def SP(self, i:int=None):
            if i is None:
                return self.getTokens(ScienceRuleParser.SP)
            else:
                return self.getToken(ScienceRuleParser.SP, i)
        def LOGICAL_OPERATOR(self):
            return self.getToken(ScienceRuleParser.LOGICAL_OPERATOR, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLogicalExp" ):
                listener.enterLogicalExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLogicalExp" ):
                listener.exitLogicalExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLogicalExp" ):
                return visitor.visitLogicalExp(self)
            else:
                return visitor.visitChildren(self)


    class ValueExpContext(ScienceruleContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ScienceruleContext
            super().__init__(parser)
            self.atom = None # ValueContext
            self.copyFrom(ctx)

        def value(self):
            return self.getTypedRuleContext(ScienceRuleParser.ValueContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterValueExp" ):
                listener.enterValueExp(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitValueExp" ):
                listener.exitValueExp(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitValueExp" ):
                return visitor.visitValueExp(self)
            else:
                return visitor.visitChildren(self)



    def sciencerule(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = ScienceRuleParser.ScienceruleContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 0
        self.enterRecursionRule(localctx, 0, self.RULE_sciencerule, _p)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 18
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,0,self._ctx)
            if la_ == 1:
                localctx = ScienceRuleParser.ParenExpContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 9
                self.match(ScienceRuleParser.T__0)
                self.state = 10
                self.sciencerule(0)
                self.state = 11
                self.match(ScienceRuleParser.T__1)
                pass

            elif la_ == 2:
                localctx = ScienceRuleParser.PresentExpContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 13
                self.attrPath()
                self.state = 14
                self.match(ScienceRuleParser.SP)
                self.state = 15
                self.match(ScienceRuleParser.T__2)
                pass

            elif la_ == 3:
                localctx = ScienceRuleParser.ValueExpContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 17
                localctx.atom = self.value()
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 37
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,2,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 35
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,1,self._ctx)
                    if la_ == 1:
                        localctx = ScienceRuleParser.LogicalExpContext(self, ScienceRuleParser.ScienceruleContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_sciencerule)
                        self.state = 20
                        if not self.precpred(self._ctx, 5):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 5)")
                        self.state = 21
                        self.match(ScienceRuleParser.SP)
                        self.state = 22
                        self.match(ScienceRuleParser.LOGICAL_OPERATOR)
                        self.state = 23
                        self.match(ScienceRuleParser.SP)
                        self.state = 24
                        self.sciencerule(6)
                        pass

                    elif la_ == 2:
                        localctx = ScienceRuleParser.ArithmeticExpContext(self, ScienceRuleParser.ScienceruleContext(self, _parentctx, _parentState))
                        localctx.lhs = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_sciencerule)
                        self.state = 25
                        if not self.precpred(self._ctx, 3):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 3)")
                        self.state = 26
                        self.match(ScienceRuleParser.SP)
                        self.state = 27
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & ((1 << ScienceRuleParser.ADD) | (1 << ScienceRuleParser.SUB) | (1 << ScienceRuleParser.MUL) | (1 << ScienceRuleParser.DIV) | (1 << ScienceRuleParser.MOD))) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 28
                        self.match(ScienceRuleParser.SP)
                        self.state = 29
                        localctx.rhs = self.sciencerule(4)
                        pass

                    elif la_ == 3:
                        localctx = ScienceRuleParser.CompareExpContext(self, ScienceRuleParser.ScienceruleContext(self, _parentctx, _parentState))
                        localctx.lhs = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_sciencerule)
                        self.state = 30
                        if not self.precpred(self._ctx, 2):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 2)")
                        self.state = 31
                        self.match(ScienceRuleParser.SP)
                        self.state = 32
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & ((1 << ScienceRuleParser.EQ) | (1 << ScienceRuleParser.NE) | (1 << ScienceRuleParser.GT) | (1 << ScienceRuleParser.LT) | (1 << ScienceRuleParser.GE) | (1 << ScienceRuleParser.LE) | (1 << ScienceRuleParser.CO))) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 33
                        self.match(ScienceRuleParser.SP)
                        self.state = 34
                        localctx.rhs = self.sciencerule(3)
                        pass

             
                self.state = 39
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,2,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class ValueContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return ScienceRuleParser.RULE_value

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)



    class BooleanContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def BOOLEAN(self):
            return self.getToken(ScienceRuleParser.BOOLEAN, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBoolean" ):
                listener.enterBoolean(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBoolean" ):
                listener.exitBoolean(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBoolean" ):
                return visitor.visitBoolean(self)
            else:
                return visitor.visitChildren(self)


    class NullContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def NULL(self):
            return self.getToken(ScienceRuleParser.NULL, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNull" ):
                listener.enterNull(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNull" ):
                listener.exitNull(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNull" ):
                return visitor.visitNull(self)
            else:
                return visitor.visitChildren(self)


    class StringContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def STRING(self):
            return self.getToken(ScienceRuleParser.STRING, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterString" ):
                listener.enterString(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitString" ):
                listener.exitString(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitString" ):
                return visitor.visitString(self)
            else:
                return visitor.visitChildren(self)


    class DoubleContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def DOUBLE(self):
            return self.getToken(ScienceRuleParser.DOUBLE, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDouble" ):
                listener.enterDouble(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDouble" ):
                listener.exitDouble(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitDouble" ):
                return visitor.visitDouble(self)
            else:
                return visitor.visitChildren(self)


    class TopicContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def attrPath(self):
            return self.getTypedRuleContext(ScienceRuleParser.AttrPathContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTopic" ):
                listener.enterTopic(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTopic" ):
                listener.exitTopic(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitTopic" ):
                return visitor.visitTopic(self)
            else:
                return visitor.visitChildren(self)


    class LongContext(ValueContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a ScienceRuleParser.ValueContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def INT(self):
            return self.getToken(ScienceRuleParser.INT, 0)
        def SUB(self):
            return self.getToken(ScienceRuleParser.SUB, 0)
        def EXP(self):
            return self.getToken(ScienceRuleParser.EXP, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLong" ):
                listener.enterLong(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLong" ):
                listener.exitLong(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLong" ):
                return visitor.visitLong(self)
            else:
                return visitor.visitChildren(self)



    def value(self):

        localctx = ScienceRuleParser.ValueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_value)
        self._la = 0 # Token type
        try:
            self.state = 52
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [ScienceRuleParser.BOOLEAN]:
                localctx = ScienceRuleParser.BooleanContext(self, localctx)
                self.enterOuterAlt(localctx, 1)
                self.state = 40
                self.match(ScienceRuleParser.BOOLEAN)
                pass
            elif token in [ScienceRuleParser.NULL]:
                localctx = ScienceRuleParser.NullContext(self, localctx)
                self.enterOuterAlt(localctx, 2)
                self.state = 41
                self.match(ScienceRuleParser.NULL)
                pass
            elif token in [ScienceRuleParser.STRING]:
                localctx = ScienceRuleParser.StringContext(self, localctx)
                self.enterOuterAlt(localctx, 3)
                self.state = 42
                self.match(ScienceRuleParser.STRING)
                pass
            elif token in [ScienceRuleParser.DOUBLE]:
                localctx = ScienceRuleParser.DoubleContext(self, localctx)
                self.enterOuterAlt(localctx, 4)
                self.state = 43
                self.match(ScienceRuleParser.DOUBLE)
                pass
            elif token in [ScienceRuleParser.SUB, ScienceRuleParser.INT]:
                localctx = ScienceRuleParser.LongContext(self, localctx)
                self.enterOuterAlt(localctx, 5)
                self.state = 45
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if _la==ScienceRuleParser.SUB:
                    self.state = 44
                    self.match(ScienceRuleParser.SUB)


                self.state = 47
                self.match(ScienceRuleParser.INT)
                self.state = 49
                self._errHandler.sync(self)
                la_ = self._interp.adaptivePredict(self._input,4,self._ctx)
                if la_ == 1:
                    self.state = 48
                    self.match(ScienceRuleParser.EXP)


                pass
            elif token in [ScienceRuleParser.ATTRNAME]:
                localctx = ScienceRuleParser.TopicContext(self, localctx)
                self.enterOuterAlt(localctx, 6)
                self.state = 51
                self.attrPath()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class AttrPathContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ATTRNAME(self):
            return self.getToken(ScienceRuleParser.ATTRNAME, 0)

        def subAttr(self):
            return self.getTypedRuleContext(ScienceRuleParser.SubAttrContext,0)


        def getRuleIndex(self):
            return ScienceRuleParser.RULE_attrPath

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAttrPath" ):
                listener.enterAttrPath(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAttrPath" ):
                listener.exitAttrPath(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAttrPath" ):
                return visitor.visitAttrPath(self)
            else:
                return visitor.visitChildren(self)




    def attrPath(self):

        localctx = ScienceRuleParser.AttrPathContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_attrPath)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 54
            self.match(ScienceRuleParser.ATTRNAME)
            self.state = 56
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,6,self._ctx)
            if la_ == 1:
                self.state = 55
                self.subAttr()


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class SubAttrContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def attrPath(self):
            return self.getTypedRuleContext(ScienceRuleParser.AttrPathContext,0)


        def getRuleIndex(self):
            return ScienceRuleParser.RULE_subAttr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSubAttr" ):
                listener.enterSubAttr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSubAttr" ):
                listener.exitSubAttr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSubAttr" ):
                return visitor.visitSubAttr(self)
            else:
                return visitor.visitChildren(self)




    def subAttr(self):

        localctx = ScienceRuleParser.SubAttrContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_subAttr)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 58
            self.match(ScienceRuleParser.T__3)
            self.state = 59
            self.attrPath()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx



    def sempred(self, localctx:RuleContext, ruleIndex:int, predIndex:int):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[0] = self.sciencerule_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def sciencerule_sempred(self, localctx:ScienceruleContext, predIndex:int):
            if predIndex == 0:
                return self.precpred(self._ctx, 5)
         

            if predIndex == 1:
                return self.precpred(self._ctx, 3)
         

            if predIndex == 2:
                return self.precpred(self._ctx, 2)
         




