package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNextTokenNumber(t *testing.T) {
	l := &Lexer{}
	l.ReadLine("123 45.67")

	tok1 := l.NextToken()
	assert.Equal(t, NUMBER, tok1.Type)
	assert.Equal(t, "123", tok1.Value)

	tok2 := l.NextToken()
	assert.Equal(t, NUMBER, tok2.Type)
	assert.Equal(t, "45.67", tok2.Value)
}

func TestNextTokenKeywordsAndIdentifiers(t *testing.T) {
	l := &Lexer{}
	l.ReadLine("int myVar = 10")

	tok1 := l.NextToken()
	assert.Equal(t, TYPEDEFF, tok1.Type)
	assert.Equal(t, "int", tok1.Value)

	tok2 := l.NextToken()
	assert.Equal(t, IDENTIFIER, tok2.Type)
	assert.Equal(t, "myVar", tok2.Value)

	tok3 := l.NextToken()
	assert.Equal(t, EQUAL, tok3.Type)
	assert.Equal(t, "=", tok3.Value)

	tok4 := l.NextToken()
	assert.Equal(t, NUMBER, tok4.Type)
	assert.Equal(t, "10", tok4.Value)
}

func TestNextTokenStrings(t *testing.T) {
	l := &Lexer{}
	l.ReadLine(`"hello world"`)

	tok := l.NextToken()
	assert.Equal(t, STRING, tok.Type)
	assert.Equal(t, "hello world", tok.Value)
}

func TestNextTokenOperators(t *testing.T) {
	l := &Lexer{}
	l.ReadLine(">= <= != == && || + - * /")

	expected := []TokenType{
		GREATER_EQUAL, LESS_EQUAL, NOT_EQUAL, EQUAL_EQUAL, AND, OR,
		PLUS, MINUS, STAR, SLASH, EOF,
	}

	for _, exp := range expected {
		tok := l.NextToken()
		assert.Equal(t, exp, tok.Type)
	}
}
