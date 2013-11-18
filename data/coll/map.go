package coll

import "clojang/data/i"
import "clojang/data/coll/hamt"
import "bufio"

const nilhash = 29320394
const starthash = 5381

type Map interface {
  With (key i.IObj, val i.IObj) Map
  Without(key i.IObj) Map
  Get(key i.IObj) i.IObj
  Contains(key i.IObj) bool
  EntryAt(key i.IObj) *hamt.Entry
  Size() uint
  Hash() uint
  String() string
  Write(w bufio.Writer)
  Equals(other i.IObj) bool
  Seq() i.ISeq
}

type hamtMap struct {
  count uint
  hash uint
  hasNil bool
  nilEntry *hamt.Entry
  root hamt.INode
}

func (hmap *hamtMap) Size() uint {
  return hmap.count
}

func (hmap *hamtMap) Hash() uint {
  return hmap.hash
} 

func (hmap *hamtMap) String() string {
  return ""
}

func (hmap *hamtMap) Write(w bufio.Writer) {

}

func (hmap *hamtMap) Equals(other i.IObj) bool {
  return true
}

func (hmap *hamtMap) Seq() i.ISeq {
  return nil
}


func (hmap *hamtMap) Contains(key i.IObj) bool {
  if key == nil {
    return hmap.hasNil
  } else if hmap.root == nil {
    return false
  } else {
    return hmap.root.EntryAt(key, key.Hash(), 0) != nil
  }
}

func (hmap *hamtMap) Get(key i.IObj) i.IObj {
  if key == nil {
    if hmap.hasNil {
      return hmap.nilEntry.Val
    } else {
      return nil
    }
  } else {
    e := hmap.root.EntryAt(key, key.Hash(), 0) 
    if e != nil {
      return e.Val
    } else {
      return nil
    }
  }
}

func (hmap *hamtMap) EntryAt(key i.IObj) *hamt.Entry {
  if key == nil {
    if hmap.hasNil {
      return hmap.nilEntry
    } else {
      return nil
    }
  } else {
    return hmap.root.EntryAt(key, key.Hash(), 0)
  }
}

func clone (m hamtMap) *hamtMap {
  return &m
}

func (hmap *hamtMap) With(key i.IObj, val i.IObj) Map {
  if key == nil {
    newMap := clone(*hmap)
    if !hmap.hasNil {
      newMap.count += 1
      newMap.hasNil = true
      newMap.hash *= nilhash
    }
    newMap.nilEntry = hamt.NewEntry(key, val)
    return newMap

  } else {
    hash := key.Hash()
    var newRoot hamt.INode
    var incCount bool
    if hmap.root == nil {
      newRoot = hamt.NewEntry(key, val)
      incCount = true
    } else {
      newRoot, incCount = hmap.root.With(hamt.NewEntry(key, val), hash, 0)
    }
    newMap := clone(*hmap)
    if incCount {
      newMap.count += 1
      if hash != 0 {
        newMap.hash *= hash
      }
    }
    newMap.root = newRoot
    return newMap
  }
}

func (hmap *hamtMap) Without(key i.IObj) Map {
  if key == nil {
    if hmap.hasNil {
      newMap := clone(*hmap)
      newMap.hasNil = false
      newMap.hash /= nilhash
      newMap.count -= 1
      return newMap
    } else {
      return hmap
    }
  } else {
    if hmap.root == nil {
      return hmap
    } else {
      hash := key.Hash()
      newRoot, decCount := hmap.root.Without(key, hash, 0)
      newMap := clone(*hmap)
      if decCount {
       newMap.count -= 1
       if hash != 0 {
         newMap.hash /= hash
       }
      }
      newMap.root = newRoot
      return newMap
    }
  }
}

func NewMap() Map {
  m := new(hamtMap)
  m.hash = starthash
  return m
}
