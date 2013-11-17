package hamt

import "clojang/data/types"

type collisionNode struct {
  collisionHash uint
  vals []*Entry
}



func newCollisionNode (hash uint, vals []*Entry) *collisionNode {
  node := new(collisionNode)
  node.collisionHash = hash
  node.vals = vals
  return node
}


func (node *collisionNode) EntryAt(key types.IObj, hash, shift uint) *Entry {
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

func (node *collisionNode) Without(key types.IObj, hash, shift uint) (INode, bool) {
  if hash == node.collisionHash {
    for i, v := range node.vals {
      if key.Equals(v.key) {
        // uh-oh, we have to delete the thing
        if len(node.vals) > 1 {
          newvals := make([]*Entry, len(node.vals)-1)

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

func (node *collisionNode) With(entry *Entry, hash, shift uint) (INode, bool) {
  if hash == node.collisionHash {
    // good times, this entry belongs in this collisionNode
    // so iterate over existing entries to check that entry.key is not already
    // present
    for i, v := range node.vals {
      if entry.key.Equals(v.key) {
        // yay, entry.key is present, so copy this node and put new entry in.
        newVals := make([]*Entry, len(node.vals))

        copy(newVals, node.vals)
        newVals[i] = entry

        return newCollisionNode(hash, newVals), false
      }
    }

    // we didn't find entry.key so copy this node and add the entry to the
    // start of the vals slice
    newArr := make([]*Entry, 1 + len(node.vals))
    copy(newArr[1:], node.vals)
    newArr[0] = entry
    return newCollisionNode(hash, newArr), true

  } else {
    return distinguishingNode(node, entry, node.collisionHash, hash, shift), true
  }
  
}