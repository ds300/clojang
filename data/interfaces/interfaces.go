package interfaces
import "bufio"

type IObj interface {
  String() string
  Hash() uint
  Equals(other IObj) bool
  Write(w bufio.Writer)
}

type ICounted interface {
  Count() uint
}

type ISeq interface {
  ISeqable
  First() IObj
  Rest() ISeq
}

type IReversible interface {
  ISeqable
  RSeq() ISeq
}

type ISeqable interface {
  Seq() ISeq
}

type IColl interface {
  ICounted
  Conj(o IObj) IColl
  Disj(o IObj) IColl
  Contains(o IObj) bool
  ValAt(k IObj) IObj
  ValAtOr(k, notFount IObj) IObj
}

type IAssoc interface {
  IColl
  Assoc(k IObj, v IObj) IAssoc
  Dissoc(k IObj) IAssoc
  EntryAt(k IObj) IMapEntry
  EntryAtOr(k, notFound IObj) IMapEntry
}

type IMap interface {
  IAssoc
  ISeqable
}

type IMapEntry interface {
  Key() IObj
  Val() IObj
}

type IVector interface {
  IAssoc
  ISeqable
}

type IFn interface {
  Arity() uint
  Invoke(...IObj) IObj
}