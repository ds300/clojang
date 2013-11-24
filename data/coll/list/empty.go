// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package list

import . "clojang/data/interfaces"
import "clojang/data/types"
import "errors"


type EmptyList struct{}

func (_ EmptyList) String() string {
  return "()"
}

func (_ EmptyList) Hash() uint32 {
  return 1
}

func (_ EmptyList) Equals(other IObj) bool {
  seq, ok := other.(ISeq)
  return ok && seq.Seq() == nil
}

func (_ EmptyList) Write(w IStringWriter) error {
  _, err := w.WriteString("()")
  return err
}

func (_ EmptyList) Type() uint32 {
  return types.ListID
}

func (_ EmptyList) First() IObj {
  return nil
}

func (_ EmptyList) Rest() ISeq {
  return nil
}

func (_ EmptyList) Nth(i uint32) (IObj, error) {
  return nil, errors.New("Index out of bounds: " + string(i))
}

func (_ EmptyList) Seq() ISeq {
  return nil
}

func (_ EmptyList) Conj(o IObj) IColl {
  return Cons(o, EmptyList{})
}

func (_ EmptyList) Count() uint32 {
  return 0
}


func (ls EmptyList) Peek() IObj {
  return nil
}

func (ls EmptyList) Pop() (IStack, error) {
  return nil, errors.New("Can't pop empty list")
}
