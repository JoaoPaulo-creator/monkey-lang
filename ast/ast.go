package ast

import "joaopaulo-creator/monkey-lang/token"


type Node interface {
  TokenLiteral() string // usado apenas para debugging e tests
}

type Statement interface {
  Node
  statementNode()
}

type Expression interface {
  Node
  expressionNode()
}


// Program eh o root de AST produzida pelo parser
type Program struct {
  Statements []Statement
}


func (p *Program) TokenLiteral() string {
  if len(p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  } else {
    return ""
  }
}


type ReturnStatement struct {
  Token token.Token
  ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type LetStatement struct {
  Token token.Token
  Name *Identifier
  Value token.Token
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
  Token token.Token
  Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

