package coll


// type Set interface {
//   With (key IObj) Set
//   Without(key IObj) Set
//   Get(key IObj) Iobj
//   Contains(key IObj) bool
//   Size() uint
// }

// type hamtSet struct {
//   count uint
//   hash uint
//   hasNil bool
//   root inode
// }

// func (hset *hamtSet) Size() uint {
//   return hset.count
// }

// func (hset *hamtSet) Hash() uint {
//   return hset.hash
// } 

// func (hset *hamtSet) Contains(key IObj) bool {
//   if key == nil {
//     return hset.hasNil
//   } else {
//     return hset.root.entryAt(key, key.Hash(), 0) != nil
//   }
// }

// func (hset *hamtSet) Get(key IObj) IObj {
//   if key == nil {
//     return nil
//   } else {
//     return hset.root.entryAt(key, key.Hash(), 0).key
//   }
// }


// func clone (s hamtSet) *hamtSet {
//   return &s
// }

// func (hset *hamtSet) With(key IObj, val interface{}) Set {
//   if key == nil {
//     if !hset.hasNil {
//       newSet := clone(*hset)
//       newSet.count += 1
//       newSet.hasNil = true
//       newSet.hash *= nilhash
//       return newSet
//     } else {
//       return hset
//     }
//   } else {
//     hash := key.Hash()
//     newRoot, incCount := hset.root.with(newEntry(key, true), hash, 0)
//     newSet := clone(*hset)
//     if incCount {
//       newSet.count += 1
//       if hash != 0 {
//         newSet.hash *= hash
//       }
//     }
//     newSet.root = newRoot
//     return newSet
//   }
// }

// func (hset *hamtSet) Without(key IObj) Set {
//   if key == nil {
//     if hset.hasNil {
//       newSet := clone(*hset)
//       newSet.hasNil = false
//       newSet.hash /= nilhash
//       newSet.count -= 1
//       return newSet
//     } else {
//       return hset
//     }
//   } else {
//     hash := key.Hash()
//     newRoot, decCount := hset.root.without(key, hash, 0)
//     newSet := clone(*hset)
//     if decCount {
//       newSet.count -= 1
//       if hash != 0 {
//         newSet.hash /= hash
//       }
//     }
//     if newRoot == nil {
//       newRoot = new(hamtNode)
//     }
//     newSet.root = newRoot
//     return newSet
//   }
// }

// func NewSet() Set {
//   m := new(hamtSet)
//   m.root = new(hamtNode)
//   m.hash = 1
//   return m
// }
