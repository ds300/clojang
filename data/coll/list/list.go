// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package list

import . "clojang/data/interfaces"
import "clojang/data/types"
import "bytes"
import "errors"
import "clojang/data/coll/sequtil"

type list struct {
  count uint32
  hash uint32
  first IObj
  rest ISeq
}

func Cons(val IObj, seq ISeqable) *list {
  var iseq ISeq
  if seq == nil || seq.Seq() == nil {
    iseq = EmptyList{}
  } else {
    iseq = seq.Seq()
  }

  ls := list{0, 0, val, iseq}
  return &ls
}

func (ls *list) String() string {
  var buf bytes.Buffer
  ls.Write(&buf)
  return buf.String()
}

func (ls *list) Hash() uint32 {
  if ls.hash == 0 {
    ls.hash = sequtil.HashSeq(ls)
    if ls.hash == 0 {
      ls.hash = 1
    }
  }
  return ls.hash
}

func (ls *list) Equals(other IObj) bool {
  seq, ok := other.(ISeqable)
  return ok && sequtil.Equals(ls, seq.Seq())
}

func (ls *list) Write(w IStringWriter) error {
  return sequtil.WriteSeq(w, '(', ls, ')')
}

func (ls *list) Type() uint32 {
  return types.ListID
}

func (ls *list) First() IObj {
  return ls.first
}

func (ls *list) Rest() ISeq {
  return ls.rest
}

func (ls *list) Nth(i uint32) (IObj, error) {
  if i == 0 {
    return ls.first, nil
  } else {
    seq := ls.Rest()
    j := i - 1
    for seq.Seq() != nil && j > 0 {
      seq = seq.Rest()
    }
    if j == 0 {
      return seq.First(), nil
    } else {
      return nil, errors.New("Index out of bounds: " + string(i))
    }
  }
}

func (ls *list) Seq() ISeq {
  return ls
}

func (ls *list) Conj(o IObj) IColl {
  return Cons(o, ls)
}

func (ls *list) Count() uint32 {
  if ls.count == 0 {
    ls.count = sequtil.CountSeq(ls)
  }
  return ls.count
}

func (ls *list) Peek() IObj {
  return ls.First()
}

func (ls *list) Pop() (IStack, error) {
  return ls.Rest().(IStack), nil
}
