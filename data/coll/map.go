// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package coll

import . "clojang/data/interfaces"
import "clojang/data/coll/hamt"
import "bufio"

const nilhash = 29320394
const starthash = 5381



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

func (hmap *hamtMap) Equals(other IObj) bool {
  return true
}

func (hmap *hamtMap) Seq() ISeq {
  return nil
}


func (hmap *hamtMap) Contains(key IObj) bool {
  if key == nil {
    return hmap.hasNil
  } else if hmap.root == nil {
    return false
  } else {
    return hmap.root.EntryAt(key, key.Hash(), 0) != nil
  }
}

func (hmap *hamtMap) Get(key IObj) IObj {
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

func (hmap *hamtMap) EntryAt(key IObj) *hamt.Entry {
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

func (hmap *hamtMap) With(key IObj, val IObj) Map {
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

func (hmap *hamtMap) Without(key IObj) Map {
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
