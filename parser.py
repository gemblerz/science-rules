from antlr4 import *
from engine.ScienceRuleLexer import ScienceRuleLexer
from engine.ScienceRuleParser import ScienceRuleParser
from engine.ScienceRuleVisitorImpl import ScienceRuleVisitorImpl

from antlr4.tree.Trees import Trees

def evaluate(query, data, show_tree=False):
    q = InputStream(query)
    # lexer
    lexer = ScienceRuleLexer(q)
    stream = CommonTokenStream(lexer)
    # parser
    parser = ScienceRuleParser(stream)
    tree = parser.sciencerule()
    # evaluator
    visitor = ScienceRuleVisitorImpl(data)
    if show_tree:
        print(Trees.toStringTree(tree, None, parser))
    output = visitor.visit(tree)
    return output