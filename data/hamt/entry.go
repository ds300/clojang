// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package hamt

import . "clojang/data/interfaces"

type Entry struct {
  IMapEntry
}

// this should never be called
func (node *Entry) Nodes() NodeIterator {
  return nil
}

func NewEntry (key IObj, val IObj) *Entry {
  e := Entry{key, val}
  return &e
}

func (entry *Entry) EntryAt (key IObj, hash, shift uint) *Entry {
  if key.Equals(entry.Key) {
    return entry
  } else {
    return nil
  }
}

func (entry *Entry) With (other *Entry, hash, shift uint) (INode, bool) {
  ehash := entry.Key.Hash()

  if ehash == hash {
    if entry.Key.Equals(other.Key) {
      return other, false

    } else {
      colliders := []*Entry{other, entry}
      return newCollisionNode(hash, colliders), true
    }

  } else if shift > 30 {
    panic("I don't know what is happening")

  } else {
    return distinguishingNode(entry, other, ehash, hash, shift), true
  }
}

func (entry *Entry) Without (key IObj, hash, shift uint) (INode, bool) {
  ehash := entry.Key.Hash()
  if ehash == hash && key.Equals(entry.Key) {
    return nil, true
  } else {
    return entry, false
  }
}