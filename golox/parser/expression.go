package parser

import (
	"fmt"
	"strings"
)

type ExprType int

const (
	EXPR_NIL ExprType = iota
	EXPR_LITERAL
	EXPR_UNARY
	EXPR_BINARY
	EXPR_GROUPING
)

type Expr struct {
	tag      ExprType
	token    Token
	children []Expr
}

func NewLiteralExpr(token Token) Expr {
	return Expr{tag: EXPR_LITERAL, token: token}
}

func NewUnaryExpr(child Expr, operand Token) Expr {
	return Expr{tag: EXPR_UNARY, token: operand, children: []Expr{child}}
}

func NewBinaryExpr(left Expr, operand Token, right Expr) Expr {
	return Expr{tag: EXPR_BINARY, token: operand, children: []Expr{left, right}}
}

// Created by wrapping expression in ()
func NewGroupingExpr(expr Expr) Expr {
	return Expr{tag: EXPR_GROUPING, children: []Expr{expr}}
}

func (e Expr) PrettyPrint() string {
	sb := strings.Builder{}

	type qElem struct {
		expr  Expr
		level int
	}

	stack := []qElem{{expr: e, level: 0}}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		elem := stack[lastIdx]
		stack = stack[:lastIdx]

		// Add indentation
		for _ = range elem.level {
			sb.WriteByte('-')
		}

		sb.WriteString(fmt.Sprintf("| %d: %s", elem.expr.token.line, elem.expr.token.lexeme))
		sb.WriteByte('\n')

		// Children need to be added to stack in reverse order so they will be popped in original order
		for i := len(elem.expr.children) - 1; i >= 0; i-- {
			child := elem.expr.children[i]
			stack = append(stack, qElem{expr: child, level: elem.level + 1})
		}
	}

	return sb.String()
}