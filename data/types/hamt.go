package types

type keyVal struct {
  key IObj
  val IObj
}

type inode interface {
  entryAt (key IObj, hash, shift uint) *keyVal
  with (entry *keyVal, hash, shift uint) (inode, bool)
  without (key IObj, hash, shift uint) (inode, bool)
  iter () iterator
}

type iterator interface {
  hasNext() bool
  next() *keyVal
}

type entryIterator struct {
  nodeIterStack Stack
}


type hamtNode struct {
  index uint
  kids []interface{}
}

type collisionNode struct {
  collisionHash uint
  vals []*keyVal
}

func newEntry (key IObj, val IObj) *keyVal {
  e := new(keyVal)
  e.key = key
  e.val = val
  return e
}

func newCollisionNode (hash uint, vals []*keyVal) *collisionNode {
  node := new(collisionNode)
  node.collisionHash = hash
  node.vals = vals
  return node
}


func (node *collisionNode) entryAt(key IObj, hash, shift uint) *keyVal {
  if hash != node.collisionHash {
    return nil
  } else {
    for _, entry := range node.vals {
      if entry.key.Equals(key) {
        return entry
      }
    }
    return nil
  }
}

func (node *collisionNode) without(key IObj, hash, shift uint) (inode, bool) {
  if hash == node.collisionHash {
    for i, v := range node.vals {
      if key.Equals(v.key) {
        // uh-oh, we have to delete the thing
        if len(node.vals) > 1 {
          newvals := make([]*keyVal, len(node.vals)-1)

          copy(newvals[0:], node.vals[0:i])
          copy(newvals[i:], node.vals[i+1:])

          return newCollisionNode(hash, newvals), true
        } else {
          return nil, true
        }
      }
    }
  } 
  return node, false
}

func (node *collisionNode) with(entry *keyVal, hash, shift uint) (inode, bool) {
  if hash == node.collisionHash {
    // good times, this entry belongs in this collisionNode
    // so iterate over existing entries to check that entry.key is not already
    // present
    for i, v := range node.vals {
      if entry.key.Equals(v.key) {
        // yay, entry.key is present, so copy this node and put new entry in.
        newVals := make([]*keyVal, len(node.vals))

        copy(newVals, node.vals)
        newVals[i] = entry

        return newCollisionNode(hash, newVals), false
      }
    }

    // we didn't find entry.key so copy this node and add the entry to the
    // start of the vals slice
    newArr := make([]*keyVal, 1 + len(node.vals))
    copy(newArr[1:], node.vals)
    newArr[0] = entry
    return newCollisionNode(hash, newArr), true

  } else {
    // oh dang, only the first shift-5 bits of the hash are the same as
    // the hashes of the keys in this collisionNode
    // so we're gonna have to create and return a new hamtnode (possibly nested)
    // to deal with that. Bummer
    

    newNode := new(hamtNode)
    maybeSubNewNode := newNode

    // hopefully the go compiler is clever enough to get rid of all these
    // unnecessary allocations. I'm certainly not clever enough to figure out
    // whether does is or not.
    var idx1, idx2, mask1, mask2 uint
    idx1 = (hash >> shift) & 31
    idx2 = (node.collisionHash >> shift) & 31
    mask1 = 1 << idx1
    mask2 = 1 << idx2

    // if the masks are the same, that means we need to create nested nodes.
    // gosh dang again man.
    for mask1 == mask2 {
      maybeSubNewNode.index = mask1
      maybeSubNewNode.kids = []interface{} {new(hamtNode)}
      maybeSubNewNode = maybeSubNewNode.kids[0].(*hamtNode)
      shift += 5
      idx1 = (hash >> shift) & 31
      idx2 = (node.collisionHash >> shift) & 31
      mask1 = 1 << idx1
      mask2 = 1 << idx2
    }

    // ok. so now we got to the bottom node which will contain the entry
    // and this collisionNode.
    maybeSubNewNode.index = mask1 | mask2

    // but wait. We need to figure out their positions in the kids slice
    if mask1 > mask2 {
      maybeSubNewNode.kids = []interface{} {entry, node}
    } else {
      maybeSubNewNode.kids = []interface{} {node, entry}
    }

    return newNode, true
  }
  
}


func (node *hamtNode) entryAt(key IObj, hash, shift uint) *keyVal {
  var idx, mask uint
  idx = (hash >> shift) & 31
  mask = 1 << idx
  

  if mask & node.index > 0 {
    // mayhaps a match?

    pos := ipopcount(node.index, idx)
    kid := node.kids[pos]

    switch kid.(type) {
    case *keyVal:
      e := kid.(*keyVal)
      if e.key.Equals(key) {
        return e
      } else {
        return nil
      }
    case inode:
      return kid.(inode).entryAt(key, hash, shift+5)
    default:
      panic("Illegal kid node man wtf?!")
    }

  } else {
    return nil
  }
}

func withoutKidAt(node *hamtNode, pos byte, mask uint) *hamtNode {
  if len(node.kids) == 1 {
    return nil
  } else {
    newNode := new(hamtNode )
    newNode.index = (mask ^ 0xFFFFFFFF) & node.index
    newKids := make([]interface{}, len(node.kids) - 1)
    copy(newKids[0:], node.kids[0:pos])
    copy(newKids[pos:], node.kids[pos+1:])
    newNode.kids = newKids
    return newNode
  }
}

func replaceKidAt(node *hamtNode, newKid interface{}, pos byte) *hamtNode {
  newNode := new(hamtNode)
  newKids := make([]interface{}, len(node.kids))
  copy(newKids, node.kids)
  newKids[pos] = newKid
  newNode.kids = newKids
  newNode.index = node.index
  return newNode
}

func (node *hamtNode) without(key IObj, hash, shift uint) (inode, bool) {
  var idx, mask uint
  idx = (hash >> shift) & 31
  mask = 1 << idx

  if mask & node.index > 0 {
    // mayhaps a match?

    pos := ipopcount(node.index, idx)
    kid := node.kids[pos]

    switch kid.(type) {
    case *keyVal:
      e := kid.(*keyVal)
      if e.key.Equals(key) {
        // uh-oh, we need to do a little delete
        return withoutKidAt(node, pos, mask), true
      } else {
        return nil, false
      }
    case inode:
      result, decCount := kid.(inode).without(key, hash, shift + 5)
      if result == nil {
        return withoutKidAt(node, pos, mask), true
      } else {
        switch result.(type) {
        case *hamtNode:
          if len(result.(*hamtNode).kids) == 1 {
            return replaceKidAt(node, result.(*hamtNode).kids[0], pos), true
          }
        case *collisionNode:
          newKid := result.(*collisionNode)
          if len(newKid.vals) == 1 {
            return replaceKidAt(node, newKid.vals[0], pos), true
          }
        default:
          panic("Illegal kid node man wtf?!")
        }
        return replaceKidAt(node, result, pos), decCount
      }
    default:
      panic("Illegal kid node man wtf?!")
    }

  } else {
    return node, false
  }
}

func (node *hamtNode) with(entry *keyVal, hash, shift uint) (inode, bool) {
  var idx, mask uint
  idx = (hash >> shift) & 31
  mask = 1 << idx
  pos := ipopcount(node.index, idx)

  if mask & node.index > 0 {
    // this hamtnode already contains something for hashes that start with the
    // same shift bits as hash, so just copy the kids straight up and replace
    // the kid at pos with a new node, then make a new hamtnode and return it
    newKids := make([]interface{}, len(node.kids))
    copy(newKids, node.kids) 

    existing := newKids[pos]

    // this should end up as true if inserting this entry involved creating
    // a new entry as opposed to just replacing an existing one
    incCount := false

    // the kid could be an entry or another node
    switch existing.(type) {
    case inode:
      // if its another node, just call .with on that guy. He knows what's up.
      newKids[pos], incCount = existing.(inode).with(entry, hash, shift + 5)
    case *keyVal:
      // if its an entry, things get tricky
      ex := existing.(*keyVal)
      if ex.key.Equals(entry.key) {
        // if it has the same key, we're just gonna replace it 
        newKids[pos] = entry
      } else {
        // if it has a different key, we need to increment the count
        incCount = true
        if shift == 30 {
          // when shift == 30 there's some serious hash collision going on, so
          // we need to make a collisionNode to take care of that mess
          newKids[pos] = newCollisionNode(hash, []*keyVal{entry, ex})
        } else {
          // otherwise it's got a different key but there's some shifting left
          // to do, so make a new hamtNode and add them both to it.
          newKid, _ := new(hamtNode).with(entry, hash, shift+5)
          newKid, _ = newKid.with(ex, ex.key.Hash(), shift+5)
          newKids[pos] = newKid
          // this could be faster and more memory efficient but I'm feeling lazy
          // at this exact moment

          // TODO: not be lazy
        }
      }
    default:
      panic("Illegal kid node man wtf?!?!")  
    }
    newNode := hamtNode{node.index, newKids}
    return &newNode, incCount

  } else {
    // Huzzah! this hamtnode doesn't already have a key whose hash has the same
    // first shift bits as hash. That makes life much more simple. We just add
    // the entry as a kid, and change the index.

    newKids := make([]interface{}, 1 + len(node.kids))
    copy(newKids[0:pos], node.kids[0:pos])
    copy(newKids[pos+1:], node.kids[pos:])

    newKids[pos] = entry

    newNode := hamtNode{mask | node.index, newKids}
    return &newNode, true
  }
}

func popcount(x uint64) (n byte) {
  // bit population count, see
  // http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
  x -= (x >> 1) & 0x5555555555555555
  x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
  x += x >> 4
  x &= 0x0f0f0f0f0f0f0f0f
  x *= 0x0101010101010101
  return byte(x >> 56)
}

func ipopcount(x uint, offset uint) byte {
  var mask uint = 0xFFFFFFFF << (32 - offset)
  return popcount(uint64(x & mask))
}
