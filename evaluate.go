package sciencerule

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gemblerz/science-rules/parser"
)

type Evaluator struct {
	lastDebugErr error

	rule string
	tree antlr.ParseTree

	testHookPanic func()
}

func NewEvaluator(rule string) (ret *Evaluator, retErr error) {
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
		}
	}()
	input := antlr.NewInputStream(rule)
	lex := parser.NewScienceRuleLexer(input)
	lex.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewScienceRuleParser(tokens)
	// TODO: maybe log properly
	p.RemoveErrorListeners()
	tree := p.Query()

	return &Evaluator{
		rule: rule,
		tree: tree,
	}, nil
}

func (e *Evaluator) Reset() error {
	e.lastDebugErr = nil
	return nil
}

func (e *Evaluator) LastDebugErr() error {
	return e.lastDebugErr
}

func (e *Evaluator) Process(items map[string]interface{}) (ret bool, retErr error) {
	e.lastDebugErr = nil
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
			ret = false
		}
	}()

	visitor := NewJsonQueryVisitorImpl(items)
	result := visitor.Visit(e.tree)
	e.lastDebugErr = visitor.debugErr
	if e.testHookPanic != nil {
		defer e.testHookPanic()
	}
	if result == nil || visitor.err != nil {
		return false, visitor.err
	}

	return result.(bool), visitor.err
}

func Evaluate(rule string, items map[string]interface{}) bool {
	ev, err := NewEvaluator(rule)
	if err != nil {
		return false
	}
	result, _ := ev.Process(items)
	return result
}
