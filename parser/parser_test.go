package parser

import (
	"testing"

	"github.com/aewens/monkey/ast"
	"github.com/aewens/monkey/lexer"
)

var PROGRAM string = `
let x = 5;
let y = 10;
let foobar = 838383;
`

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Token literal not 'let': %q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Not let statement: %T", s)
		return false
	}

	value := letStatement.Name.Value
	if value != name {
		t.Errorf("Name value is not %s: %s", name, value)
		return false
	}

	literal := letStatement.Name.TokenLiteral()
	if literal != name {
		t.Errorf("Name token literal is not %s: %s", name, literal)
		return false
	}

	return true
}

func TestLetStatements(t *testing.T) {
	l := lexer.New(PROGRAM)
	p := New(l)

	program := p.ParserProgram()
	if program == nil {
		t.Fatal("ParserProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Statements does not contain 3 statements: %d",
			len(program.Statements))
	}

	tests := []struct{ expectedIdentifier string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}
