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

