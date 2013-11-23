// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package coll

import . "clojang/data/interfaces"
import "sync"

type lazySeq struct {
  v IObj
  next ISeq
  mutex *sync.Mutex
  fn func () ISeq
}

func (ls *lazySeq) First() IObj {
  mutex := ls.mutex
  if mutex != nil {
    mutex.Lock()
    if ls.fn != nil {
      res := ls.fn()
      if res == nil {
       ls.v = nil
       ls.next = nil
      } else {
       ls.v = res.First()
       ls.next = res.Rest()
      }
      ls.fn = nil
      ls.mutex = nil
    }
    mutex.Unlock()
  }

  return ls.v
}

func (ls *lazySeq) Rest() ISeq {
  ls.First()
  return ls.next
}

func LazySeq(fn func () ISeq) ISeq {
  ls := lazySeq{nil, nil, new(sync.Mutex), fn}
  return &ls
}