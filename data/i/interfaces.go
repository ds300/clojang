package i
import "bufio"

type IObj interface {
  String() string
  Hash() uint
  Equals(other IObj) bool
  Write(w bufio.Writer)
}

type ISeq interface {
  First() IObj
  Rest() ISeq
}

type Seqable interface {
  Seq() ISeq
}