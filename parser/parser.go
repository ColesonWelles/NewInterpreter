package parser

import (
	"github.com/ColesonWelles/NewInterpreter/ast"
	"github.com/ColesonWelles/NewInterpreter/lexer"
	"github.com/ColesonWelles/NewInterpreter/token"
)

type Parser struct {
	l         *lexer.Lexer // Pointer to an instance of the lexer, on which we repeatedly call NextToken() to get the next token in the input
	curToken  token.Token  // The current token under examination
	peekToken token.Token  // Needed for the decision if curToken doesn't give us enough information
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
