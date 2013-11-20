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