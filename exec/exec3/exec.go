package exec3

// import . "clojang/data/interfaces"
import "fmt"

const (
  ADD uint32 = iota // r1, r2, r3 :: r1 = r2 + r3
  SUB             // r1, r2, r3 :: r1 = r2 - r3
  DIV             // r1, r2, r3 :: r1 = r2 / r3
  MULT            // r1, r2, r3 :: r1 = r2 * r3
  LOAD16          // r1, uint24 :: r1 = uint16
  LOAD32          // r1, _, _ -> uint32 :: r1 = uint32
  JMPF
  JMPB
  JMPFEQ          // r1, r2, offset :: PC += r1==r2 ? offset : 1  
  JMPBEQ          // r1, r2, offset :: PC += r1==r2 ? -offset : 1
  JMPFNEQ          // r1, r2, offset :: PC += r1==r2 ? offset : 1  
  JMPBNEQ          // r1, r2, offset :: PC += r1==r2 ? -offset : 1
  PRINT            // r1 :: println(r1)
  END
)

func interpret (instructions []uint32) {
  var reg [256]int32
  var r1 uint32
  var r2 uint32
  var r3 uint32
  var offset uint32

  var PC uint32
  var currentInstruction uint32
  // var SP uint16


  defer func() {
        if r := recover(); r != nil {
            fmt.Println("program exit")
        }
  }()

  inst := [...]func(){
    ADD: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] + reg[r3]
      PC++
    },
    SUB: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] - reg[r3]
      PC++
    },
    DIV: func () {
      r3 = (currentInstruction & 0xFF000000) >> 24
      if r3 != 0 {
        r1 = (currentInstruction & 0x0000FF00) >> 8
        r2 = (currentInstruction & 0x00FF0000) >> 16
        reg[r1] = reg[r2] / reg[r3]
      }
      PC++
    },
    MULT: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] * reg[r3]
      PC++
    },
    LOAD16: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      reg[r1] = int32((currentInstruction & 0xFFFF0000) >> 16)
      PC++
    },
    LOAD32: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      PC++
      reg[r1] = int32(instructions[PC])
      PC++
    },
    JMPF: func () {
      offset = (currentInstruction & 0xFFFFFF00) >> 8
      PC += offset
    },
    JMPB: func () {
      offset = (currentInstruction & 0xFFFFFF00) >> 8
      PC -= offset
    },
    JMPFEQ: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] == reg[r2] {
        PC += offset
      } else {
        PC += 1
      }
    },
    JMPBEQ: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] == reg[r2] {
        PC -= offset
      } else {
        PC += 1
      }
    },
    JMPFNEQ: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] != reg[r2] {
        PC += offset
      } else {
        PC += 1
      }
    },
    JMPBNEQ: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] != reg[r2] {
        PC -= offset
      } else {
        PC += 1
      }
    },
    PRINT: func () {
      r1 = (currentInstruction & 0x0000FF00) >> 8
      println(reg[r1])
      PC++
    },
    END: func () {
      PC = 0xFFFFFFFF
    },
  }

  for {

    currentInstruction = instructions[PC]

    inst[(currentInstruction & 0x000000FF)]()
    // fmt.Println(reg[0:6])
  }
}

func interpret2 (instructions []uint32) {
  var reg [256]int32
  var r1 uint32
  var r2 uint32
  var r3 uint32
  var val int32
  var offset uint32

  var PC uint32
  var currentInstruction uint32
  // var SP uint16



  for {

    currentInstruction = instructions[PC]

    switch (currentInstruction & 0x000000FF){
    case ADD:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] + reg[r3]
      PC++
    
    case SUB:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] - reg[r3]
      PC++
    
    case DIV:
      r3 = (currentInstruction & 0xFF000000) >> 24
      if r3 != 0 {
        r1 = (currentInstruction & 0x0000FF00) >> 8
        r2 = (currentInstruction & 0x00FF0000) >> 16
        reg[r1] = reg[r2] / reg[r3]
      }
      PC++
    
    case MULT:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      r3 = (currentInstruction & 0xFF000000) >> 24
      reg[r1] = reg[r2] * reg[r3]
      PC++
    
    case LOAD16:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      reg[r1] = int32((currentInstruction & 0xFFFF0000) >> 16)
      PC++
    
    case LOAD32:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      PC++
      val = int32(instructions[PC])
      reg[r1] = val
      PC++
    
    case JMPF:
      offset = (currentInstruction & 0xFFFFFF00) >> 8
      PC += offset
    
    case JMPB:
      offset = (currentInstruction & 0xFFFFFF00) >> 8
      PC -= offset
    
    case JMPFEQ:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] == reg[r2] {
        PC += offset
      } else {
        PC += 1
      }
    
    case JMPBEQ:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] == reg[r2] {
        PC -= offset
      } else {
        PC += 1
      }
    
    case JMPFNEQ:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] != reg[r2] {
        PC += offset
      } else {
        PC += 1
      }
    
    case JMPBNEQ:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      r2 = (currentInstruction & 0x00FF0000) >> 16
      offset = (currentInstruction & 0xFF000000) >> 24
      if reg[r1] != reg[r2] {
        PC -= offset
      } else {
        PC += 1
      }
    
    case PRINT:
      r1 = (currentInstruction & 0x0000FF00) >> 8
      println(reg[r1])
      PC++
    
    case END:
      return
    }
    // fmt.Println(reg[0:6])
  }
}

// this program sums the natural numbers to ten million (without catching overflow)
var exampleProgram []uint32 = []uint32 {
  LOAD16 ^ 0x00010500,          // r[5] = 1
  LOAD32 ^ 0x00000100,
  10000000,                               // put the number 5000 in reg[1]
  JMPFEQ ^ 0x04000100,              // jump forward somewhere if reg[1] == reg[0]
  ADD ^ 0x02010200,                  // r[2] = r[1] + r[2]
  SUB ^ 0x05010100,                  // r[1]-=r[5]
  JMPB ^ 0x00000300,
  PRINT ^ 0x00000200,                // print r[2]
  END,
}


