// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package exec

const EXCEPTION uint = 0xFFFFFFFF
const TAIL_CALL uint = 0xEFFFFFFF

func interpreter (instructions []instruction) instruction {
  var PC *uint
  N := uint(len(instructions))
  var TAIL *[]instruction

  return func (cs *CallStack, es *ExecStack, pc *uint, tail *[]instruction) {
    *pc++
    for {
      for *PC < N {
        instructions[*PC](cs, es, PC, TAIL)
      }

      if *PC < TAIL_CALL {
        if *PC != uint(len(instructions)) { panic("What the jazz?") }
        return
      } else if *PC == TAIL_CALL {
        instructions = *TAIL
        *TAIL = nil
        *PC = 0
      } else if *PC == EXCEPTION {
        *pc = EXCEPTION
        return
      }
    }
  }
}
