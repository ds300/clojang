package exec

type Exception struct {
  thrown interface{}
  callStack *CallStack
}

type Instruction func (cs *CallStack, es *ExecStack, pc *uint, tail *[]Instruction, exception *Exception)


const THROW Instruction = func (cs *CallStack, es *ExecStack, pc *uint, tail *[]Instruction, exception *Exception) {
  valToThrow := ex.Pop()
  stack := CallStack{cs.count, cs.top}
  ex := Exception{valToThrow, *stack}

  *pc = EXCEPTION
  *exception = *ex
}



const PRINT_TOS instruction = func (cs *CallStack, es *ExecStack, pc *uint, tail *[]Instruction, exception *Exception) {
  *pc++
  fmt.Println(es.Pop())
  es.Push(nil)
}


func MakeTailCall = func (callSite *callStackElem, instructions *[]Instruction) instruction {
  return func (cs *CallStack, es *ExecStack, pc *uint, tail *[]Instruction, exception *Exception) {
    *pc = TAIL_CALL
    *tail = *instructions
    cs.Pop()
    cs.Push(callSite)
  }
}