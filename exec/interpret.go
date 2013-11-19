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
