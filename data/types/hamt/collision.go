package hamt

import "clojang/data/types/i"

type collisionNode struct {
  collisionHash uint
  entries []*Entry
}



func newCollisionNode (hash uint, vals []*Entry) *collisionNode {
  node := new(collisionNode)
  node.collisionHash = hash
  node.entries = vals
  return node
}


func (node *collisionNode) EntryAt(key i.IObj, hash, shift uint) *Entry {
  if hash != node.collisionHash {
    return nil
  } else {
    for _, entry := range node.entries {
      if entry.Key.Equals(key) {
        return entry
      }
    }
    return nil
  }
}

func (node *collisionNode) Without(key i.IObj, hash, shift uint) (INode, bool) {
  if hash == node.collisionHash {
    for i, v := range node.entries {
      if key.Equals(v.Key) {
        // uh-oh, we have to delete the thing
        if len(node.entries) > 1 {
          newvals := make([]*Entry, len(node.entries)-1)

          copy(newvals[0:], node.entries[0:i])
          copy(newvals[i:], node.entries[i+1:])

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
    // so iterate over existing entries to check that entry.Key is not already
    // present
    for i, v := range node.entries {
      if entry.Key.Equals(v.Key) {
        // yay, entry.Key is present, so copy this node and put new entry in.
        newVals := make([]*Entry, len(node.entries))

        copy(newVals, node.entries)
        newVals[i] = entry

        return newCollisionNode(hash, newVals), false
      }
    }

    // we didn't find entry.Key so copy this node and add the entry to the
    // start of the vals slice
    newArr := make([]*Entry, 1 + len(node.entries))
    copy(newArr[1:], node.entries)
    newArr[0] = entry
    return newCollisionNode(hash, newArr), true

  } else {
    return distinguishingNode(node, entry, node.collisionHash, hash, shift), true
  }
  
}