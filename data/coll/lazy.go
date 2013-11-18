package coll

import "clojang/data/i"
import "sync"

type lazySeq struct {
  v i.IObj
  next i.ISeq
  mutex *sync.Mutex
  fn func () i.ISeq
}

func (ls *lazySeq) First() i.IObj {
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

func (ls *lazySeq) Rest() i.ISeq {
  ls.First()
  return ls.next
}

func LazySeq(fn func () i.ISeq) i.ISeq {
  ls := lazySeq{nil, nil, new(sync.Mutex), fn}
  return &ls
}