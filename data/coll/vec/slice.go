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
import "clojang/data/primitives"
import "clojang/data/coll/sequtil"
import "clojang/data/coll/lazy"
import "clojang/data/coll/list"
import "bytes"
import "errors"

func SliceSeq(slice *[]IObj, index int) ISeq {
  if index < len(*slice) {
    return lazy.LazySeq(func () ISeq {
      return list.Cons((*slice)[index], SliceSeq(slice, index + 1))   
    })
  } else {
    return nil
  }
}

func RSliceSeq(slice *[]IObj, index int) ISeq {
  if index >= 0 {
    return lazy.LazySeq(func () ISeq {
      return list.Cons((*slice)[index], RSliceSeq(slice, index - 1))   
    })
  } else {
    return nil
  }
}



type sliceVector struct {
  slice []IObj
  hash uint32
}


func (sv *sliceVector) String() string {
  var buf bytes.Buffer
  sv.Write(&buf)
  return buf.String()
}

func (sv *sliceVector) Hash() uint32 {
  if sv.hash == 0 {
    sv.hash = sequtil.HashSeq(sv.Seq())
  }
  return sv.hash
}

func (sv *sliceVector) Equals(other IObj) bool {
  seq, ok := other.(ISeqable)
  return ok && sequtil.Equals(sv.Seq(), seq.Seq())
}

func (sv *sliceVector) Write(w IStringWriter) error {
  return sequtil.WriteSeq(w, '[', sv.Seq(), ']')
}

func (sv *sliceVector) Type() uint32 {
  return types.VectorID
}

func (sv *sliceVector) Count() uint32 {
  return uint32(len(sv.slice))
}

func (sv *sliceVector) Seq() ISeq {
  return SliceSeq(&sv.slice, 0)
}

func (sv *sliceVector) RSeq() ISeq {
  return RSliceSeq(&sv.slice, len(sv.slice) - 1)
}

func (sv *sliceVector) Conj(o IObj) IColl {
  newSlice := make([]IObj, len(sv.slice) + 1)
  copy(newSlice, sv.slice)
  return &sliceVector{newSlice, 0}
}

func (sv *sliceVector) Assoc(k IObj, v IObj) (IAssoc, error) {
  l, ok := k.(primitives.Long)
  if ok {
    i := int(l)
    if i > 0 && i < len(sv.slice) {
      newSlice := make([]IObj, len(sv.slice))
      copy(newSlice, sv.slice)
      newSlice[i] = v
      return &sliceVector{newSlice, 0}, nil
    } else if i == len(sv.slice) {
      return sv.Conj(v).(IVector), nil
    } else {
      return nil, errors.New("index out of bounds")
    }
  }

  return nil, errors.New("bad index type")
}

func (sv *sliceVector) Get(k IObj) IObj {
  i, ok := k.(primitives.Long)
  if ok && int(i) >= 0 && int(i) < len(sv.slice) {
    return sv.slice[int(i)]
  } else {
    return nil
  }
}

func (sv *sliceVector) GetOr(k, notFound IObj) IObj {
  i, ok := k.(primitives.Long)
  if ok && int(i) >= 0 && int(i) < len(sv.slice) {
    return sv.slice[int(i)]
  } else {
    return notFound
  }
}

func (sv *sliceVector) Contains(o IObj) bool {
  i, ok := o.(primitives.Long)
  return ok && int(i) >= 0 && int(i) < len(sv.slice)
}

func (sv *sliceVector) Peek() IObj {
  return sv.slice[len(sv.slice) - 1]
}

func (sv *sliceVector) Pop() (IStack, error) {
  switch len(sv.slice) {
  case 0:
    fallthrough
  case 1:
    fallthrough
  case 2:
    panic("slice vector with len < 3 is not alowed to occur")
  case 3:
    return &tupleVector{sv.slice[0], sv.slice[1], 0}, nil
  }
  return &sliceVector{sv.slice[:len(sv.slice) - 1], 0}, nil
}

