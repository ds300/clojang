// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package hamt

import . "clojang/data/interfaces"

const NILHASH = 29320394

type INode interface {
  EntryAt (key IObj, hash, shift uint32) *Entry
  With (entry *Entry, hash, shift uint32) (INode, *Entry)
  Without (key IObj, hash, shift uint32) (INode, *Entry)
}

type NodeIterator interface {
  HasNext() bool
  Next() INode
}

func popcount(x uint32) byte {
  x -= ((x >> 1) & 0x55555555)
  x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
  return byte((((x + (x >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24)
}

func ipopcount(x, offset uint32) byte {
  var mask uint32 = 0xFFFFFFFF << (32 - offset)
  return popcount(x & mask)
}

func idxMask(hash, shift uint32) (uint32, uint32) {
  var idx, mask uint32
  idx = (hash >> shift) & 31
  mask = 0x80000000 >> idx
  return idx, mask
}

type hamtNode struct {
  index uint32
  kids []INode
}

type hamtNodeIterator struct {
  index uint32
  kids *[]INode
}

func (hni *hamtNodeIterator) HasNext() bool {
  return hni.index < uint32(len(*hni.kids))
}

func (hni *hamtNodeIterator) Next() INode {
  ret := (*hni.kids)[hni.index]
  hni.index += 1
  return ret
}

func (node *hamtNode) Nodes() NodeIterator {
  ret := hamtNodeIterator{0, &node.kids}
  return &ret
}

func (node *hamtNode) EntryAt(key IObj, hash, shift uint32) *Entry {
  idx, mask := idxMask(hash, shift)
  
  if mask & node.index > 0 {
    // mayhaps a match?
    pos := ipopcount(node.index, idx)
    return node.kids[pos].EntryAt(key, hash, shift+5)

  } else {
    return nil
  }
}

func withoutKidAt(node *hamtNode, pos byte, mask uint32) *hamtNode {
  if len(node.kids) == 1 {
    return nil
  } else {
    newNode := new(hamtNode )
    newNode.index = (mask ^ 0xFFFFFFFF) & node.index
    newKids := make([]INode, len(node.kids) - 1)
    copy(newKids[0:], node.kids[0:pos])
    copy(newKids[pos:], node.kids[pos+1:])
    newNode.kids = newKids
    return newNode
  }
}

func replaceKidAt(node *hamtNode, newKid INode, pos byte) *hamtNode {
  newNode := new(hamtNode)
  newKids := make([]INode, len(node.kids))
  copy(newKids, node.kids)
  newKids[pos] = newKid
  newNode.kids = newKids
  newNode.index = node.index
  return newNode
}

func (node *hamtNode) Without(key IObj, hash, shift uint32) (INode, *Entry) {
  idx, mask := idxMask(hash, shift)

  if mask & node.index > 0 {
    // mayhaps a match?

    pos := ipopcount(node.index, idx)

    result, removed := node.kids[pos].Without(key, hash, shift + 5)

    if removed != nil {
      // if something was actually removed
      if result == nil {
        switch len(node.kids) {
        case 1:
          return nil, removed
        case 2:
          goodKidPos := pos ^ 1
          _, dontGetFancy := node.kids[goodKidPos].(*hamtNode)
          if !dontGetFancy {
            // if this node only has one kid left which is not another
            // hamtNode, then we can get fancy and just return the remaining
            // kid since it doesn't use bitfiddling to figure out what's where.
            return node.kids[goodKidPos], removed
          }
          fallthrough
        default:
          return withoutKidAt(node, pos, mask), removed
        }
      } else {
        if len(node.kids) == 1 {
          _, dontGetFancy := result.(*hamtNode)
          if !dontGetFancy {
            // if this node only has one kid, and the result is not another
            // hamtNode, then we can get fancy and just return the result node
            // since it doesn't use bitfiddling to figure out what's where.
            return result, removed
          }
        }
        return replaceKidAt(node, result, pos), removed
      } 
    } 
  } 
  return node, nil
}

func (node *hamtNode) With(entry *Entry, hash, shift uint32) (INode, *Entry) {
  idx, mask := idxMask(hash, shift)

  pos := ipopcount(node.index, idx)

  if mask & node.index > 0 {
    // this hamtnode already contains something for hashes that start with the
    // same shift bits as hash, so just copy the kids straight up and replace
    // the kid at pos with a new node, then make a new hamtnode and return it
    newKids := make([]INode, len(node.kids))
    copy(newKids, node.kids) 

    newKid, replaced := newKids[pos].With(entry, hash, shift + 5)

    newKids[pos] = newKid

    newNode := hamtNode{node.index, newKids}

    return &newNode, replaced

  } else {
    // D'oh! This hamtnode doesn't already have a key whose hash has the same
    // first shift bits as hash. So we need to add the entry as a kid, and
    // change the index.

    newKids := make([]INode, 1 + len(node.kids))
    copy(newKids[0:pos], node.kids[0:pos])
    copy(newKids[pos+1:], node.kids[pos:])

    newKids[pos] = entry

    newNode := hamtNode{mask | node.index, newKids}
    return &newNode, nil
  }
}

// this thing should only be called when the hashes are different
func distinguishingNode(e1, e2 INode, h1, h2, shift uint32) *hamtNode {
  // oh dang, only the first shift-5 bits of h1 and h2 are the same.
  // so we're gonna have to create and return a new hamtnode (possibly nested)
  // to deal with that. Bummer
  newNode := new(hamtNode)
  maybeSubNewNode := newNode
  var mask1, mask2 uint32

  mask1 = 1 << (32 - ((h1 >> shift) & 31))
  mask2 = 1 << (32 - ((h2 >> shift) & 31))

  // if the masks are the same, that means we need to create nested nodes.
  // gosh dang again man.
  for mask1 == mask2 {
    maybeSubNewNode.index = mask1
    maybeSubNewNode.kids = []INode {new(hamtNode)}
    maybeSubNewNode = maybeSubNewNode.kids[0].(*hamtNode)
    shift += 5
    mask1 = 1 << (32 - ((h1 >> shift) & 31))
    mask2 = 1 << (32 - ((h2 >> shift) & 31))
  }

  // ok. so now we got to the bottom node which will contain the entry
    // and this collisionNode.
  maybeSubNewNode.index = mask1 | mask2

  // but wait. We need to figure out their positions in the kids slice
  if mask1 > mask2 {
    maybeSubNewNode.kids = []INode {e1, e2}
  } else {
    maybeSubNewNode.kids = []INode {e2, e1}
  }

  return newNode
}
