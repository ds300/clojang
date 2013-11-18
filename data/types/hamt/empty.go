package hamt

import "clojang/data/types/i"

type EmptyNode bool

func (e EmptyNode) EntryAt (key i.IObj, hash, shift uint) *Entry {
  return nil
}

func (e EmptyNode) With (entry *Entry, hash, shift uint) (INode, bool) {
  return entry, true
}

func (e EmptyNode) Without (key i.IObj, hash, shift uint) (INode, bool) {
  return e, false
}