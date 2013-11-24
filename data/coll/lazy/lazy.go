// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package lazy

import . "clojang/data/interfaces"
import "clojang/data/types"
import "clojang/data/coll/sequtil"
import "clojang/data/coll/list"
import "sync"
import (
  "bytes"
  "errors"
)

type lazySeq struct {
  cell ISeq
  mutex *sync.Mutex
  fn func () ISeq
  hash uint32
  count uint32
}

func (ls *lazySeq) First() IObj {
  mutex := ls.mutex
  if mutex != nil {
    mutex.Lock()
    if ls.fn != nil {
      ls.cell = ls.fn()
      ls.fn = nil
      ls.mutex = nil
    }
    mutex.Unlock()
  }

  if ls.cell != nil {
    return ls.cell.First()
  } else {
    return nil
  }
}

func (ls *lazySeq) Rest() ISeq {
  ls.First()
  if ls.cell != nil {
    return ls.cell.Rest()
  } else {
    return nil
  }
}

func LazySeq(fn func () ISeq) ISeq {
  ls := lazySeq{nil, new(sync.Mutex), fn, 0, 0}
  return &ls
}



func (ls *lazySeq) String() string {
  var buf bytes.Buffer
  ls.Write(&buf)
  return buf.String()
}

func (ls *lazySeq) Hash() uint32 {
  if ls.hash == 0 {
    ls.hash = sequtil.HashSeq(ls)
    if ls.hash == 0 {
      ls.hash = 1
    }
  }
  return ls.hash
}

func (ls *lazySeq) Equals(other IObj) bool {
  seq, ok := other.(ISeqable)
  return ok && sequtil.Equals(ls, seq.Seq())
}

func (ls *lazySeq) Write(w IStringWriter) error {
  return sequtil.WriteSeq(w, '(', ls, ')')
}

func (ls *lazySeq) Type() uint32 {
  return types.LazySeqID
}

func (ls *lazySeq) Nth(i uint32) (IObj, error) {
  if i == 0 {
    return ls.First(), nil
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

func (ls *lazySeq) Seq() ISeq {
  ls.First()
  if ls.cell == nil {
    return nil
  } else {
    return ls
  }
}

func (ls *lazySeq) Conj(o IObj) IColl {
  return list.Cons(o, ls)
}

func (ls *lazySeq) Count() uint32 {
  if ls.count == 0 {
    ls.count = sequtil.CountSeq(ls)
  }
  return ls.count
}

func (ls *lazySeq) Peek() IObj {
  return ls.First()
}

func (ls *lazySeq) Pop() (IStack, error) {
  if ls.Seq() != nil {
    return ls.Rest().(IStack), nil
  } else {
    return nil, errors.New("Can't pop from empty seq")
  }
}



