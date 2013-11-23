// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package set

import . "clojang/data/interfaces"
import "clojang/data/coll/hamt"

// const nilhash = 29320394
// const starthash = 5381

type Set interface {
  With (key IObj) Set
  Without(key IObj) Set
  Get(key IObj) IObj
  Contains(key IObj) bool
  Size() uint
  Hash() uint
  String() string
  Equals(other IObj) bool
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

func (hset *hamtSet) Equals(other IObj) bool {
  return true
}

func (hset *hamtSet) Contains(key IObj) bool {
  if key == nil {
    return hset.hasNil
  } else {
    return hset.root.EntryAt(key, key.Hash(), 0) != nil
  }
}

func (hset *hamtSet) Get(key IObj) IObj {
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

func (hset *hamtSet) With(key IObj) Set {
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
    var newRoot hamt.INode
    var incCount bool
    if hset.root == nil {
      newRoot = hamt.NewEntry(key, nil)
      incCount = true
    } else {
      newRoot, incCount = hset.root.With(hamt.NewEntry(key, nil), hash, 0)
    }
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

func (hset *hamtSet) Without(key IObj) Set {
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
    if hset.root == nil {
      return hset
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

      newSet.root = newRoot
      return newSet
    }
      
    
  }
}

func NewSet() Set {
  m := new(hamtSet)
  m.hash = starthash
  return m
}
