package coll

import "clojang/data/i"

type list struct {
  val i.IObj
  next i.ISeq
}

func (ls *list) First() i.IObj {
  return ls.val
}

func (ls *list) Rest() i.ISeq {
  return ls.next
}

func Cons(val i.IObj, seq i.ISeq) i.ISeq {
  ls := list{val, seq}
  return &ls
}

