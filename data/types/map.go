package types

const nilhash = 29320394

type MapEntry interface {
  Key() IObj
  Val() IObj
}

func (e *keyVal) Key() IObj {
  return e.key
}

func (e *keyVal) Val() IObj {
  return e.val
}

type Map interface {
  With (key IObj, val IObj) Map
  Without(key IObj) Map
  Get(key IObj) IObj
  Contains(key IObj) bool
  EntryAt(key IObj) MapEntry
  Size() uint
}

type hamtMap struct {
  count uint
  hash uint
  hasNil bool
  nilEntry *keyVal
  root inode
}

func (hmap *hamtMap) Size() uint {
  return hmap.count
}

func (hmap *hamtMap) Hash() uint {
  return hmap.hash
} 

func (hmap *hamtMap) Contains(key IObj) bool {
  if key == nil {
    return hmap.hasNil
  } else {
    return hmap.root.entryAt(key, key.Hash(), 0) != nil
  }
}

func (hmap *hamtMap) Get(key IObj) IObj {
  if key == nil {
    if hmap.hasNil {
      return hmap.nilEntry.val
    } else {
      return nil
    }
  } else {
    return hmap.root.entryAt(key, key.Hash(), 0).val
  }
}

func (hmap *hamtMap) EntryAt(key IObj) MapEntry {
  if key == nil {
    if hmap.hasNil {
      return hmap.nilEntry
    } else {
      return nil
    }
  } else {
    return hmap.root.entryAt(key, key.Hash(), 0)
  }
}

func clone (m hamtMap) *hamtMap {
  return &m
}

func (hmap *hamtMap) With(key IObj, val interface{}) Map {
  if key == nil {
    newMap := clone(*hmap)
    if !hmap.hasNil {
      newMap.count += 1
      newMap.hasNil = true
      newMap.hash *= nilhash
    }
    newMap.nilEntry = newEntry(key, val)
    return newMap

  } else {
    hash := key.Hash()
    newRoot, incCount := hmap.root.with(newEntry(key, val), hash, 0)
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
    hash := key.Hash()
    newRoot, decCount := hmap.root.without(key, hash, 0)
    newMap := clone(*hmap)
    if decCount {
      newMap.count -= 1
      if hash != 0 {
        newMap.hash /= hash
      }
    }
    if newRoot == nil {
      newRoot = new(hamtNode)
    }
    newMap.root = newRoot
    return newMap
  }
}

func NewMap() Map {
  m := new(hamtMap)
  m.root = new(hamtNode)
  m.hash = 1
  return m
}
