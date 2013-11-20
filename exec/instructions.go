// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

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