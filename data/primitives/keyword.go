// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package primitives

import . "clojang/data/interfaces"
import "clojang/data/types"
import "bufio"

type keyword struct {
  ns string
  name string
  hash uint32
}

func (k *keyword) String() string {
  if k.ns != "" {
    return ":" + k.ns + "/" + k.name
  } else {
    return ":" + k.name
  }
}

func (k *keyword) Hash() uint32 {
  if k.hash == 0 {
    h := NewString(k.String()).Hash()
    if h == 0 {
      h = 1
    }
    k.hash = h
  }
  return k.hash
}

func (k *keyword) Equals(other IObj) bool {
   v, ok := other.(*keyword)
   return ok && k.name == v.name && k.ns == v.ns
}

func (k *keyword) Write(w *bufio.Writer) error {
  _, err := w.WriteString(k.String())
  return err
}

func (k *keyword) Type() uint32 {
  return types.KeywordID
}

func (k *keyword) Name() string {
  return k.name
}