package primitives

import . "clojang/data/interfaces"
import "clojang/data/types"
import "bufio"

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

func (s *symbol) Write(w *bufio.Writer) error {
  _, err := w.WriteString(s.String())
  return err
}

func (s *symbol) Type() uint32 {
  return types.KeywordID
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