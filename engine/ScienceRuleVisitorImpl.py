from .ScienceRuleVisitor import ScienceRuleVisitor

class ScienceRuleVisitorImpl(ScienceRuleVisitor):
    def __init__(self, measurements):
        self.measurements = measurements

    def visitCompareExp(self, ctx):
        lhs = self.visit(ctx.lhs)
        rhs = self.visit(ctx.rhs)
        op = ctx.op.text.lower()
        operations = {
            "==": lambda: lhs == rhs,
            "!=": lambda: lhs != rhs,
            ">": lambda: lhs > rhs,
            "<": lambda: lhs < rhs,
            ">=": lambda: lhs >= rhs,
            "<=": lambda: lhs <= rhs,
            "co": lambda: rhs in lhs,
        }
        return operations.get(op, lambda:None)()

    def visitArithmeticExp(self, ctx):
        lhs = self.visit(ctx.lhs)
        rhs = self.visit(ctx.rhs)
        op = ctx.op.text
        operations = {
            "+": lambda: lhs + rhs,
            "-": lambda: lhs - rhs,
            "*": lambda: lhs * rhs,
            "/": lambda: lhs / rhs,
            "%": lambda: rhs % rhs,
        }
        return operations.get(op, lambda: None)()

    def visitParenExp(self, ctx):
        return self.visit(ctx.sciencerule())

    def visitPresentExp(self, ctx):
        pass

    def visitLogicalExp(self, ctx):
        lhs = self.visit(ctx.sciencerule(0))
        rhs = self.visit(ctx.sciencerule(1))
        op = ctx.LOGICAL_OPERATOR()
        if op == "or":
            return lhs if lhs else rhs
        else:
            return lhs if not lhs else rhs

    def visitValueExp(self, ctx):
        return self.visit(ctx.atom)

    def visitBoolean(self, ctx):
        v = ctx.getText()
        return True if v == "true" else False

    def visitNull(self, ctx):
        return None
    
    def visitString(self, ctx):
        return ctx.STRING()
    
    def visitDouble(self, ctx):
        return float(ctx.getText())
    
    def visitLong(self, ctx):
        return int(ctx.getText())

    def visitTopic(self, ctx):
        return self.visit(ctx.attrPath())

    def visitAttrPath(self, ctx):
        if ctx.subAttr() == None:
            print(ctx.getText())
            return 1
        else:
            return self.visit(ctx.subAttr())

    def visitSubAttr(self, ctx):
        return self.visit(ctx.attrPath())
