// Code generated from ScienceRule.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // ScienceRule

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseScienceRuleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseScienceRuleVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitLong(ctx *LongContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScienceRuleVisitor) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}
