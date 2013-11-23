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

type symbol struct {
  ns string
  name string
  hash uint32
  meta IObj
}

func (s *symbol) String() string {
  if s.ns != "" {
    return s.ns + "/" + s.name
  } else {
    return s.name
  }
}

func (s *symbol) Hash() uint32 {
  if s.hash == 0 {
    h := NewString(s.String()).Hash()
    if h == 0 {
      h = 1
    }
    s.hash = h
  }
  return s.hash
}

func (s *symbol) Equals(other IObj) bool {
   v, ok := other.(*symbol)
   return ok && s.name == v.name && s.ns == v.ns
}

func (s *symbol) Write(w IStringWriter) error {
  _, err := w.WriteString(s.String())
  return err
}

func (s *symbol) Type() uint32 {
  return types.SymbolID
}

func (s *symbol) Name() string {
  return s.name
}

func (s *symbol) WithMeta(meta IObj) IMeta {
  newSymbol := symbol{s.ns, s.name, s.hash, meta}
  return &newSymbol
}

func (s *symbol) Meta() IObj {
  return s.meta
}

func NewSymbol(ns, name string) *symbol {
  newSymbol := symbol{ns, name, 0, nil}
  return &newSymbol
}