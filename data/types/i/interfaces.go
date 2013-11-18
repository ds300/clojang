package i

type IObj interface {
  String() string
  Hash() uint
  Equals(other IObj) bool
}

type ISeq interface {
  First() IObj
  Rest() ISeq
}