package hamt

import "clojang/data/types"

type Entry struct {
  key types.IObj
  val types.IObj
}

func NewEntry (key types.IObj, val types.IObj) *Entry {
  e := Entry{key, val}
  return &e
}

func (entry *Entry) EntryAt (key types.IObj, hash, shift uint) *Entry {
  if key.Equals(entry.key) {
    return entry
  } else {
    return nil
  }
}

func (entry *Entry) With (other *Entry, hash, shift uint) (INode, bool) {
  ehash := entry.key.Hash()

  if ehash == hash {
    if entry.key.Equals(other.key) {
      return other

    } else {
      colliders := []*Entry{other, entry}
      return newCollisionNode(hash, colliders)
    }

  } else if shift > 30 {
    panic("I don't know what is happening")

  } else {
    return distinguishingNode(entry, other, ehash, hash, shift), true
  }
}

func (entry *Entry) Without (key types.IObj, hash, shift uint) (INode, bool) {
  ehash := entry.key.Hash()
  if ehash == hash && key.Equals(entry.key) {
    return nil, true
  } else {
    return entry, false
  }
}