package interpreter

import (
	lx "BFG/lexer"
	lc "BFG/loopcount"
	"bufio"
	"fmt"
	"os"
)

type BF struct {
  Board []byte
  Pointer int
  LC *lc.LoopCount
  LX *lx.Lexer
  Reader bufio.Reader
}

func Abs(val int) int {
  if (val < 0) {
    return -1 * val
  }
  return val
}

func New(lc *lc.LoopCount, lx *lx.Lexer) BF {
  return BF{Board: make([]byte, 255), Pointer: 0, LC: lc, LX: lx, Reader: *bufio.NewReader(os.Stdin)}
}

func (bf *BF)HandleToken(token lx.Token) {
  switch token.Kind {
  case lx.PLUS:
    bf.Board[bf.Pointer] += 1
  case lx.MINUS:
    bf.Board[bf.Pointer] -= 1
  case lx.NEXT:
    bf.Pointer = Abs(bf.Pointer + 1) % len(bf.Board)
  case lx.PREVIOUS:
    bf.Pointer = Abs(bf.Pointer - 1) % len(bf.Board)

  case lx.OPENLOOP:

    if (bf.Board[bf.Pointer] == 0) {

      token := bf.LX.GetToken()
      for token.Kind != lx.CLOSELOOP {
        token = bf.LX.GetToken()
      }

    } else {

      if (!bf.LC.Contains(token.Pos)) {
        bf.LC.Push(token.Pos)
      }

    }

  case lx.CLOSELOOP:

    if (bf.Board[bf.Pointer] != 0) {
      pos, exists := bf.LC.Peek()
      if !exists {
        panic("Unable to find Previous Loop [")
      }
      bf.LX.CurPos = pos
      bf.LX.NextChar()
    } else {
      bf.LC.Pop()
    }
  case lx.OUTPUT:
    fmt.Printf("%c", bf.Board[bf.Pointer])

  case lx.INPUT:
    inputChar, err := bf.Reader.ReadByte()
    if err != nil {
      panic(err)
    }
    bf.Board[bf.Pointer] = inputChar

  default:
    fmt.Println("Unresolved")
  }
}
