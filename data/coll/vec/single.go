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

type singleElemVector struct {
  elem IObj
}

func (sev *singleElemVector) String() string {
  var buf bytes.Buffer
  sev.Write(buf)
  return buf.String()
}

func (sev *singleElemVector) Hash() uint32 {
  return sequtil.HashIndexed(0, sev.elem.Hash())
}

func (sev *singleElemVector) Equals(other IObj) bool {
  seq, ok := other.(ISeq)
  return ok && sequtil.Equals(sev.Seq(),seq)
}

func (sev *singleElemVector) Write(w IStringWriter) error {
  w.WriteRune('[')
  if sev.elem == nil {
    w.WriteString("nil")
  } else {
    sev.elem.Write(w)
  }
  _, err := w.WriteRune(']')
  return err
}

func (sev *singleElemVector) Type() uint32 {
  return types.VectorID
}

func (sev *singleElemVector) Count() uint32 {
  return 1
}

func (sev *singleElemVector) Seq() ISeq {
  return list.Cons(sev.elem, nil)
}

func (sev *singleElemVector) RSeq() ISeq {
  return sev.Seq()
}

func (sev *singleElemVector) Conj(o IObj) IColl {
  return tupleVector{sev.elem, o}
}

func (sev *singleElemVector) Assoc(k IObj, v IObj) (IAssoc, error) {
  v, ok := k.(primitives.Long)
  if ok {
    switch int64(v) {
    case 0:
      newSev := singleElemVector{v}
      return &newSev, nil
    case 1:
      tuple := tupleVector{sev.elem, v, 0}
      return &tuple
    default:
      return nil, errors.New("Index out of bounds: " + v.String())
    }
  }
  return nil, errors.New("Bad index type")
}

func (sev *singleElemVector) Get(k IObj) IObj {
  v, ok := k.(primitives.Long)
  if ok && v == 0 {
    return sev.elem
  } else {
    return nil
  }
}

func (sev *singleElemVector) GetOr(k, notFound IObj) IObj {
  v, ok := k.(primitives.Long)
  if ok && v == 0 {
    return sev.elem
  } else {
    return notFound
  }
}

func (sev *singleElemVector) Contains(o IObj) bool {
  v, ok := o.(primitives.Long)
  return ok && v == 0
}
