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
import "strconv"

type String struct {
  string
  hash uint32
}

func (s *String) String() string {
  return s.string
}

func (s *String) Hash() uint32 {
  // djb2, maybe terribly implemented
  if s.hash == 0 {
    var hash uint32 = 5381
    for _, c := range s.string {
      hash = ((hash << 5) + hash) + uint32(c)
    }
    if hash == 0 {
      hash = 1
    }
    s.hash = hash
    return s.hash
  }
  return s.hash
}

func (s *String) Equals(other IObj) bool {
  v, ok := other.(*String)
  return ok && v.string == s.string
}

func (s *String) Write(w *bufio.Writer) error {
  _, err := w.WriteString(strconv.Quote(s.string))
  return err
}

func (s *String) Type() uint32 {
  return types.StringID
}


func NewString(s string) *String {
  ret := String{s, 0}
  return &ret
}