package ast

import "github.com/ellemouton/monkey/token"

/*
The abstract syntax tree (AST) we will construct consists entirely of
Nodes that are connected to each other. Some of these nodes implement
Statement and some implement Expression.
*/

type Node interface {
	// TokenLiteral returns the literal value of the token it is
	// associated with. For debugging and testing.
	TokenLiteral() string
}

type Statement interface {
	Node

	// statementNode is a dummy method used so that the Go compiler
	// will throw an error if we use a Statement where an
	// Expression should be used.
	statementNode()
}

type Expression interface {
	Node

	// expressionNode is a dummy method used so that the Go compiler
	// will throw an error if we use an Expression where a
	// Statement should be used.
	expressionNode() // dummy method
}

// Program is the root node of every AST. Every Monkey program is a
// series of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement lets us store things like: let x = 5;
type LetStatement struct {
	// Token is the token that this AST node is associated with.
	Token token.Token // the token.LET token

	// Name is the name of the variable
	Name *Identifier

	// Value points to the expression to the right of the equal sign.
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
