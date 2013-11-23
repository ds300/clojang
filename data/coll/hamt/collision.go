// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package hamt

import . "clojang/data/interfaces"

type collisionNode struct {
  collisionHash uint32
  entries []*Entry
}

type collisionNodeIterator struct {
  index uint32
  entries *[]*Entry
}

func (cni *collisionNodeIterator) HasNext() bool {
  return cni.index < uint32(len(*cni.entries))
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



func newCollisionNode (hash uint32, vals []*Entry) *collisionNode {
  node := collisionNode{hash, vals}
  return &node
}


func (node *collisionNode) EntryAt(key IObj, hash, shift uint32) *Entry {
  if hash != node.collisionHash {
    return nil
  } else {
    for _, entry := range node.entries {
      if entry.Key().Equals(key) {
        return entry
      }
    }
    return nil
  }
}

func (node *collisionNode) Without(key IObj, hash, shift uint32) (INode, *Entry) {
  if hash == node.collisionHash {
    for i, entry := range node.entries {
      if key.Equals(entry.Key()) {
        // uh-oh, we have to delete the thing
        switch len(node.entries) {
        case 1:
          // if this is the only thing left, just get rid of this node
          // I'm not sure this will ever get called given the case 2
          return nil, entry
        case 2:
          // there will only be one entry left after this, so just return that
          // entry instead
          return node.entries[i ^ 1], entry
        default:
          newEntries := make([]*Entry, len(node.entries)-1)

          copy(newEntries[0:], node.entries[0:i])
          copy(newEntries[i:], node.entries[i+1:])

          return newCollisionNode(hash, newEntries), entry
        }
      }
    }
  } 
  return node, nil
}

func (node *collisionNode) With(entry *Entry, hash, shift uint32) (INode, *Entry) {
  if hash == node.collisionHash {
    // good times, this entry belongs in this collisionNode
    // so iterate over existing entries to check that entry.Key is not already
    // present
    for i, existing := range node.entries {
      if entry.Key().Equals(existing.Key()) {
        // yay, entry.Key is present, so copy this node and put new entry in.
        newVals := make([]*Entry, len(node.entries))

        copy(newVals, node.entries)
        newVals[i] = entry

        return newCollisionNode(hash, newVals), existing
      }
    }

    // we didn't find entry.Key so copy this node and add the entry to the
    // start of the vals slice
    newArr := make([]*Entry, 1 + len(node.entries))
    copy(newArr[1:], node.entries)
    newArr[0] = entry
    return newCollisionNode(hash, newArr), nil

  } else {
    return distinguishingNode(node, entry, node.collisionHash, hash, shift), nil
  }
  
}