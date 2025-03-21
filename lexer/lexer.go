package lexer

import (
	"bytes"
	"errors"
)

const (
  PLUS = iota
  MINUS = iota
  NEXT = iota
  PREVIOUS = iota
  OPENLOOP = iota
  CLOSELOOP = iota
  INPUT = iota
  OUTPUT = iota
  EOF = iota
)

type Token struct {
  Kind int
  Val byte
  Pos int
}

type Lexer struct {
  Source []byte
  CurPos int
  CurChar byte
}

func MapCharToken(val byte, pos int) (Token, error) {
  kind := 0
  switch val {
    case byte('+'):
      kind = PLUS
    case byte('-'):
      kind = MINUS
    case byte('>'):
      kind = NEXT
    case byte('<'):
      kind = PREVIOUS
    case byte('['):
      kind = OPENLOOP
    case byte(']'):
      kind = CLOSELOOP
    case byte('.'):
      kind = OUTPUT
    case byte(','):
      kind = INPUT
    case 26:
      kind = EOF
    default:
      return Token{Kind: 0, Val: 0, Pos: pos}, errors.New("Something went wrong in the lexer")
  }
  return Token{Kind: kind, Val: val, Pos: pos}, nil
}

func New(source []byte) Lexer {
  lex := Lexer{Source: source, CurPos: -1, CurChar: 0}
  lex.NextChar()
  return lex
}

func (lexer *Lexer) NextChar() {
  if (lexer.CurPos + 1 >= len(lexer.Source)) {
    lexer.CurChar = 26
    lexer.CurPos = len(lexer.Source)
  } else {
    lexer.CurPos += 1
    lexer.CurChar = lexer.Source[lexer.CurPos]
  }
}

func (lexer *Lexer) GetToken() Token {

  for !bytes.Contains([]byte("+-.,[]<>"), []byte{lexer.CurChar}) && lexer.CurPos < len(lexer.Source) {
    lexer.NextChar()
  }

  token, err := MapCharToken(lexer.CurChar, lexer.CurPos)

  if err != nil {
    panic(err)
  }

  lexer.NextChar()

  return token
}
