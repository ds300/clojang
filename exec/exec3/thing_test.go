package exec3

import "testing"


func TestExec(t *testing.T) {
  interpret(&stackFrame{nil, nil, nil, "hey"}, exemplar2)
}