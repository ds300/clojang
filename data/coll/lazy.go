package coll

import "clojang/data/i"
import "sync"

type lazySeq struct {
  v i.IObj
  fn func () i.ISeq
  next i.ISeq
}

func (ls *lazySeq) First() i.IObj {
  if ls.fn != nil {
    res := ls.fn()
    if res == nil {
      ls.v = nil
      ls.fn = nil
      ls.next = nil
      return nil
    } else {
      ls.v = res.First()
      ls.next = res.Rest()
      ls.fn = nil
      return ls.v
    }
  } else {
    return ls.v
  }
}

func (ls *lazySeq) Rest() i.ISeq {
  ls.First()
  return ls.next
}

func lockFn(fn func () i.ISeq) func () i.ISeq {
  var mutex sync.Mutex
  return func () i.ISeq {
    mutex.Lock()
    res := fn()
    mutex.Unlock()
    return res
  }
}

func LazySeq(fn func () i.ISeq) i.ISeq {
  ls := lazySeq{nil, lockFn(fn), nil}
  return &ls
}