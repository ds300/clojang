package exec2

import "fmt"
import . "clojang/data/interfaces"
import "clojang/data/primitives"
import "clojang/data/coll/vec"
import "errors"

// const (
//   RETURN byte = iota
//   PUSH byte = iota
//   POP byte = iota
//   DO byte = iota
//   UNDO byte = iota
//   CALL byte = iota
//   CJMP byte = iota
//   JMP byte = iota
// )

type stack struct {
  val IObj
  below *stack
}

type stackFrame struct {
  fnName string
  callSite [2]int
  depth int
  below *stackFrame
}

type ArgCollector struct {
  Args []IObj
  at byte
  seqAt byte
}

func (ac *ArgCollector) Collect(arg IObj) error {
  if ac.at < ac.seqAt {
    ac.Args[ac.at] = arg
    ac.at++
  } else if ac.at == ac.seqAt {
    ac.Args[ac.at] = vec.EmptyVector{}.Conj(arg)
    ac.at++
  } else if ac.seqAt != 0 {
    ac.Args[ac.at - 1] = ac.Args[ac.at - 1].Conj(arg)
  } else {
    return errors.New("ArityExceptionBlah")
  }
  return nil
}

func (s *stackFrame) Error () string {
  return "oops at " + s.fnName
}

type RETURN struct {}
type PUSH struct {
  val IObj
}
type PUSHARG struct{}
type POP struct {}
type DO struct {}
type UNDO struct {}
type CALL struct {
  sf *stackFrame
}
type CJUMP struct {
  to int
}
type JUMP struct {
  to int
}
type THROW struct {
  callSite [2]int
}

const FALSE primitives.Bool = primitives.Bool(false)


func interpreter (instrs []interface{}) func (*stackFrame, ISeq) (IObj, error) {
  return func (callStack *stackFrame, args ISeq) (IObj, error) {
    instructions := instrs
    PC := 0
    SP := &stack{args.(IObj), nil}
    dosp := SP

    for {
      switch inst := instructions[PC].(type) {
      case RETURN:
        return SP.val, nil
      case PUSH:
        SP = &stack{inst.val, SP}
      case POP:
        SP = SP.below
      case DO:
        dosp = SP
      case UNDO:
        SP = dosp
      case CALL:
        inst.sf.below = callStack
        args = SP.val.(ISeq)
        SP = SP.below
        ret, err := (*inst.fn)(inst.sf, args)
        if err != nil {
          return nil, err
        } else {
          SP = &stack{ret, SP}
        }
      case CJUMP:
        if SP.val == nil || SP.val == FALSE {
          PC = inst.to
          SP = SP.below
        }
        continue
      case JUMP:
        PC = inst.to
      case THROW:
        return nil, &stackFrame{"throw", inst.callSite, callStack.depth + 1, callStack}
      }
      PC++
    }
  }
}

var plus = func (callStack *stackFrame, args ISeq) (IObj, error) {
  var result INumeric = primitives.Long(0)
  for args.Seq() != nil {
    x, ok := args.First().(INumeric)
    if ok {
      result = result.Plus(x)
    } else {
      return nil, errors.New("wut mayne")
    }
    args = args.Rest()
  }
  return result.(IObj), nil
}

var print = func (callStack *stackFrame, args ISeq) (IObj, error) {
  fmt.Println(args)
  return nil, nil
}

// func interpreter (instructions []byte, operands [][]IObj) callable {
//   return func (callStack *stackFrame, ISeq args) (IObj, error) {
//     PC := 0

//     st := &stack{args, st}


//   }
// }

// func interpreter (instructions []byte, operands [][]int) {
//   PC := 0

//   var st *stack

//   for {
//     switch instructions[PC] {
//     case RETURN:
//       return
//     case PUSH:
//       st = &stack{operands[PC][0], st}
//       PC++
//     case POP:
//     }
//   } 
  
// }