package hamt

import "clojang/data/i"

type collisionNode struct {
  collisionHash uint
  entries []*Entry
}

type collisionNodeIterator struct {
  index uint
  entries *[]*Entry
}

func (cni *collisionNodeIterator) HasNext() bool {
  return cni.index < uint(len(*cni.entries))
}

func (cni *collisionNodeIterator) Next() INode {
  ret := (*cni.entries)[cni.index]
  cni.index += 1
  return ret
}

func (node *collisionNode) Nodes() NodeIterator {
  ret := collisionNodeIterator{0, &node.entries}
  return &ret
}



func newCollisionNode (hash uint, vals []*Entry) *collisionNode {
  node := collisionNode{hash, vals}
  return &node
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
        switch len(node.entries) {
        case 1:
          // if this is the only thing left, just get rid of this node
          // I'm not sure this will ever get called given the case 2
          return nil, true
        case 2:
          // there will only be one entry left after this, so just return that
          // entry instead
          return node.entries[i ^ 1], true
        default:
          newEntries := make([]*Entry, len(node.entries)-1)

          copy(newEntries[0:], node.entries[0:i])
          copy(newEntries[i:], node.entries[i+1:])

          return newCollisionNode(hash, newEntries), true
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