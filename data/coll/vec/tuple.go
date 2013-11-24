// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package vec

import . "clojang/data/interfaces"
import "clojang/data/types"
import "clojang/data/coll/list"
import "clojang/data/primitives"
import "clojang/data/coll/sequtil"
import "bytes"
import "errors"

type tupleVector struct {
  k IObj
  v IObj
  hash uint32
}

func (tv *tupleVector) String() string {
  var buf bytes.Buffer
  tv.Write(&buf)
  return buf.String()
}

func (tv *tupleVector) Hash() uint32 {
  if tv.hash == 0 {
    tv.hash = sequtil.HashSeq(tv.Seq())
  }
  return tv.hash
}

func (tv *tupleVector) Equals(other IObj) bool {
  seq, ok := other.(ISeqable)
  return ok && sequtil.Equals(tv.Seq(),seq.Seq())
}

func (tv *tupleVector) Write(w IStringWriter) error {
  return sequtil.WriteSeq(w, '[', tv.Seq(), ']')
}

func (tv *tupleVector) Type() uint32 {
  return types.VectorID
}

func (tv *tupleVector) Count() uint32 {
  return 2
}

func (tv *tupleVector) Seq() ISeq {
  return list.Cons(tv.k, list.Cons(tv.v, nil))
}

func (tv *tupleVector) RSeq() ISeq {
  return list.Cons(tv.v, list.Cons(tv.k, nil))
}

func (tv *tupleVector) Conj(o IObj) IColl {
  return nil
}

func (tv *tupleVector) Assoc(k IObj, v IObj) (IAssoc, error) {
  i, ok := k.(primitives.Long)
  if ok {
    switch int64(i) {
    case 0:
      tuple := tupleVector{v, tv.v, 0}
      return &tuple, nil
    case 1:
      tuple := tupleVector{tv.k, v, 0}
      return &tuple, nil
    case 2:
      return tv.Conj(v).(IVector), nil
    default:
      return nil, errors.New("Index out of bounds: " + v.String())
    }
  }
  return nil, errors.New("Bad index type")
}

func (tv *tupleVector) Get(k IObj) IObj {
  i, ok := k.(primitives.Long)
  if ok {
    switch i {
    case 0:
      return tv.k
    case 1:
      return tv.v
    } 
  }
  return nil
}

func (tv *tupleVector) GetOr(k, notFound IObj) IObj {
  i, ok := k.(primitives.Long)
  if ok {
    switch i {
    case 0:
      return tv.k
    case 1:
      return tv.v
    } 
  }
  return notFound
}

func (tv *tupleVector) Contains(o IObj) bool {
  v, ok := o.(primitives.Long)
  return ok && (v == 0 || v == 1)
}

func (tv *tupleVector) Key() IObj {
  return tv.k
}

func (tv *tupleVector) Val() IObj {
  return tv.v
}


func (tv *tupleVector) Peek() IObj {
  return tv.v
}

func (tv *tupleVector) Pop() IStack, error {
  return &singleElemVector{tv.k, 0}
}



