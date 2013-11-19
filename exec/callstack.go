package exec

import "fmt"

type stackElem struct {
  line uint
  file *string
  name *string
  below *stackElem
}

func StackElem(line uint, file *string, name *string) *stackElem {
  x := stackElem{line, file, name, nil}
  return &x
}

type CallStack struct {
  count uint
  top *stackElem
}

func (cs *CallStack) Push(e *stackElem) {
  e.below = cs.top
  cs.top = e
}

func (cs *CallStack) Pop() *stackElem {
  ret := cs.top
  if ret != nil {
    cs.top = ret.below
  }
  return ret
}



















