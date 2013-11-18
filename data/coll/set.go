package coll

import "clojang/data/i"
import "clojang/data/coll/hamt"

// const nilhash = 29320394
// const starthash = 5381

type Set interface {
  With (key i.IObj) Set
  Without(key i.IObj) Set
  Get(key i.IObj) i.IObj
  Contains(key i.IObj) bool
  Size() uint
  Hash() uint
  String() string
  Equals(other i.IObj) bool
}

type hamtSet struct {
  count uint
  hash uint
  hasNil bool
  root hamt.INode
}

func (hset *hamtSet) Size() uint {
  return hset.count
}

func (hset *hamtSet) Hash() uint {
  return hset.hash
} 

func (hset *hamtSet) String() string {
  return ""
}

func (hset *hamtSet) Equals(other i.IObj) bool {
  return true
}

func (hset *hamtSet) Contains(key i.IObj) bool {
  if key == nil {
    return hset.hasNil
  } else {
    return hset.root.EntryAt(key, key.Hash(), 0) != nil
  }
}

func (hset *hamtSet) Get(key i.IObj) i.IObj {
  if key == nil {
    return nil
  } else {
    e := hset.root.EntryAt(key, key.Hash(), 0) 
    if e != nil {
      return e.Key
    } else {
      return nil
    }
  }
}

func cloneSet (m hamtSet) *hamtSet {
  return &m
}

func (hset *hamtSet) With(key i.IObj) Set {
  if key == nil {
    if hset.hasNil {
      return hset
    } else {
      newSet := cloneSet(*hset)
      newSet.count += 1
      newSet.hasNil = true
      newSet.hash *= nilhash
      return newSet
    }

  } else {
    hash := key.Hash()
    newRoot, incCount := hset.root.With(hamt.NewEntry(key, nil), hash, 0)
    newSet := cloneSet(*hset)
    if incCount {
      newSet.count += 1
      if hash != 0 {
        newSet.hash *= hash
      }
    }
    newSet.root = newRoot
    return newSet
  }
}

func (hset *hamtSet) Without(key i.IObj) Set {
  if key == nil {
    if hset.hasNil {
      newSet := cloneSet(*hset)
      newSet.hasNil = false
      newSet.hash /= nilhash
      newSet.count -= 1
      return newSet
    } else {
      return hset
    }
  } else {
    hash := key.Hash()
    newRoot, decCount := hset.root.Without(key, hash, 0)
    newSet := cloneSet(*hset)
    if decCount {
      newSet.count -= 1
      if hash != 0 {
        newSet.hash /= hash
      }
    }
    if newRoot == nil {
      newRoot = hamt.EmptyNode(false)
    }
    newSet.root = newRoot
    return newSet
  }
}

func NewSet() Set {
  m := new(hamtSet)
  m.root = hamt.EmptyNode(false)
  m.hash = starthash
  return m
}
