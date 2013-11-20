// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package exec

type execStackElem struct {
  val interface{}
  below *execStackElem
}

type ExecStack struct {
  count uint
  top *execStackElem
}

func (cs *ExecStack) Push(e interface{}) {
  newElem := execStackElem(e, cs.top)
  cs.top = &newElem
}

func (cs *ExecStack) Pop() interface{} {
  ret := cs.top
  cs.top = ret.below
  return ret.val
}

