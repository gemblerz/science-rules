# Generated from ScienceRule.g4 by ANTLR 4.9
from antlr4 import *
if __name__ is not None and "." in __name__:
    from .ScienceRuleParser import ScienceRuleParser
else:
    from ScienceRuleParser import ScienceRuleParser

# This class defines a complete generic visitor for a parse tree produced by ScienceRuleParser.

class ScienceRuleVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by ScienceRuleParser#compareExp.
    def visitCompareExp(self, ctx:ScienceRuleParser.CompareExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#arithmeticExp.
    def visitArithmeticExp(self, ctx:ScienceRuleParser.ArithmeticExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#parenExp.
    def visitParenExp(self, ctx:ScienceRuleParser.ParenExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#presentExp.
    def visitPresentExp(self, ctx:ScienceRuleParser.PresentExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#logicalExp.
    def visitLogicalExp(self, ctx:ScienceRuleParser.LogicalExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#valueExp.
    def visitValueExp(self, ctx:ScienceRuleParser.ValueExpContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#boolean.
    def visitBoolean(self, ctx:ScienceRuleParser.BooleanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#null.
    def visitNull(self, ctx:ScienceRuleParser.NullContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#string.
    def visitString(self, ctx:ScienceRuleParser.StringContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#double.
    def visitDouble(self, ctx:ScienceRuleParser.DoubleContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#long.
    def visitLong(self, ctx:ScienceRuleParser.LongContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#topic.
    def visitTopic(self, ctx:ScienceRuleParser.TopicContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#attrPath.
    def visitAttrPath(self, ctx:ScienceRuleParser.AttrPathContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by ScienceRuleParser#subAttr.
    def visitSubAttr(self, ctx:ScienceRuleParser.SubAttrContext):
        return self.visitChildren(ctx)



del ScienceRuleParser