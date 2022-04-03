// Code generated from ScienceRule.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // ScienceRule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ScienceRuleParser.
type ScienceRuleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ScienceRuleParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#parenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#presentExp.
	VisitPresentExp(ctx *PresentExpContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#logicalExp.
	VisitLogicalExp(ctx *LogicalExpContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#long.
	VisitLong(ctx *LongContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#listStrings.
	VisitListStrings(ctx *ListStringsContext) interface{}

	// Visit a parse tree produced by ScienceRuleParser#subListOfStrings.
	VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{}
}
