// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package util

type Stack interface {
  Peek() interface{}
  Pop() interface{}
  Empty() bool
  Push(item interface{})
}

type stackElem struct {
  val interface{}
  below *stackElem
}

type stack struct {
  top *stackElem
  count uint
}

func (s *stack) Peek() interface{} {
  if s.count == 0 {
    return nil
  } else {
    return s.top.val
  }
}

func (s *stack) Pop() interface{} {
  if s.count == 0 {
    return nil
  } else {
    res := s.top.val
    s.top = s.top.below
    s.count -= 1
    return res
  }
}

func (s *stack) Emtpy() bool {
  return s.count == 0
}

func (s *stack) Push(item interface{}) {
  e := new(stackElem)
  e.val = item
  e.below = s.top
  s.top = e
  s.count += 1
}

func NewStack () Stack {
  return new(stack)
}