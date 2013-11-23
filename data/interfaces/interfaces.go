// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package interfaces
import "bufio"

type IObj interface {
  String() string
  Hash() uint32
  Equals(other IObj) bool
  // for serializing
  Write(w *bufio.Writer) error
  Type() uint32
}

type IMeta interface {
  WithMeta(meta IObj) IMeta
  Meta() IObj
}

type INamed interface {
  Name() string
}

type INumeric interface {
  Mult(other INumeric) INumeric
  Div(other INumeric) (INumeric, error)
  Plus(other INumeric) INumeric
  Sub(other INumeric) INumeric
  Mod(other INumeric) (INumeric, error)
}

type ICounted interface {
  Count() uint32
}

type ISeq interface {
  ISeqable
  First() IObj
  Rest() ISeq
  Nth() (ISeq, error)
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
  Contains(o IObj) bool
  ValAt(k IObj) IObj
  ValAtOr(k, notFount IObj) IObj
}

type IAssoc interface {
  IColl
  Assoc(k IObj, v IObj) (IAssoc, error)
}

type IMap interface {
  IAssoc
  ISeqable
  EntryAt(k IObj) IMapEntry
  EntryAtOr(k, notFound IObj) IMapEntry
  Dissoc(k IObj) IAssoc
}

type IMapEntry interface {
  Key() IObj
  Val() IObj
}

type ISet interface {
  IColl
  Disj(o IObj) IColl
}


type IVector interface {
  IAssoc
  ISeqable
}

type IFn interface {
  Arity() uint
  Invoke(...IObj) IObj
}

