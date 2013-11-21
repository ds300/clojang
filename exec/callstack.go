// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package exec

import "fmt"

type callSite struct {
  line uint
  file *string
  name *string
  below *callSite
}

func CallSite(line uint, file *string, name *string) *callSite {
  x := callSite{line, file, name, nil}
  return &x
}

type CallStack struct {
  count uint
  top *callSite
}

func (cs *CallStack) Push(e *callSite) {
  e.below = cs.top
  cs.top = e
}

func (cs *CallStack) Pop() *callSite {
  ret := cs.top
  if ret != nil {
    cs.top = ret.below
  }
  return ret
}



















