package hamt

import "clojang/data/types/i"

type Entry struct {
  Key i.IObj
  Val i.IObj
}

func NewEntry (key i.IObj, val i.IObj) *Entry {
  e := Entry{key, val}
  return &e
}

func (entry *Entry) EntryAt (key i.IObj, hash, shift uint) *Entry {
  if key.Equals(entry.Key) {
    return entry
  } else {
    return nil
  }
}

func (entry *Entry) With (other *Entry, hash, shift uint) (INode, bool) {
  ehash := entry.Key.Hash()

  if ehash == hash {
    if entry.Key.Equals(other.Key) {
      return other, false

    } else {
      colliders := []*Entry{other, entry}
      return newCollisionNode(hash, colliders), true
    }

  } else if shift > 30 {
    panic("I don't know what is happening")

  } else {
    return distinguishingNode(entry, other, ehash, hash, shift), true
  }
}

func (entry *Entry) Without (key i.IObj, hash, shift uint) (INode, bool) {
  ehash := entry.Key.Hash()
  if ehash == hash && key.Equals(entry.Key) {
    return nil, true
  } else {
    return entry, false
  }
}