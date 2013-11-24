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
import "errors"

type emptyVector struct {}

func (ev emptyVector) String() string {
  return "[]"
}

func (ev emptyVector) Hash() uint32 {
  return 1
}

func (ev emptyVector) Equals(other IObj) bool {
  _, ok := other.(emptyVector)
  return ok
}

func (ev emptyVector) Write(w IStringWriter) error {
  _, err := w.WriteString("[]")
  return err
}

func (ev emptyVector) Type() uint32 {
  return types.VectorID
}

func (ev emptyVector) Count() uint32 {
  return 0
}

func (ev emptyVector) Seq() ISeq {
  return list.EmptyList{}
}

func (ev emptyVector) RSeq() ISeq {
  return list.EmptyList{}
}

func (ev emptyVector) Conj(o IObj) IColl {
  sev := singleElemVector{o, 0}
  return &sev
}

func (ev emptyVector) Contains(o IObj) bool {
  return false
}

func (ev emptyVector) Assoc(k IObj, v IObj) (IAssoc, error) {
  i, ok := k.(primitives.Long)
  if ok {
    if i == 0 {
      return ev.Conj(v).(IVector), nil
    } else {
      return nil, errors.New("Index out of bounds" + i.String())
    }
  } else {
    return nil, errors.New("Bad index type")
  }
}

func (ev emptyVector) Get(k IObj) IObj {
  return nil
}

func (ev emptyVector) GetOr(k, notFound IObj) IObj {
  return notFound
}


