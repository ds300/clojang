package interfaces
import "bufio"

type IObj interface {
  String() string
  Hash() uint
  Equals(other IObj) bool
  Write(w bufio.Writer)
}

type ICounted interface {
  IObj
  Count() uint
}

type ISeq interface {
  IObj
  ISeqable
  First() IObj
  Rest() ISeq
}

type IReversible interface {
  ISeqable
  RSeq() ISeq
}

type ISeqable interface {
  IObj
  Seq() ISeq
}

type IColl interface {
  IObj
  ICounted
  Conj(o IObj) IColl
  Disj(o IObj) IColl
  Contains(o IObj) bool
  ValAt(k IObj) IObj
}

type IAssoc interface {
  IObj
  IColl
  Assoc(k IObj, v IObj) IAssoc
  Dissoc(k IObj) IAssoc
  EntryAt(k IObj) IMapEntry
  EntryAtOr(k, notFound IObj) IObj
}

type IMap interface {
  IAssoc
  ISeqable
}

type IMapEntry {
  IObj
  Key() IObj
  Val() IObj
}

type IVector interface {
  IAssoc
  ISeqable
}

type IFn interface {
  IObj
  Arity() uint
  Invoke(...IObj) IObj
}