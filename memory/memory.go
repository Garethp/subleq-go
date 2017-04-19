package memory

import (
)

var instructions = []int{9,-1,3,10,-1,6,0,0,-1,72,105}
var expressionPointers = []int{0, 3, 6}
var pointer = 0

func LoadFile() {
}

func GetExpression() (int, int, int) {
    var point = expressionPointers[pointer]

    return instructions[point], instructions[point+1], instructions[point+2]
}

func WriteExpression(int a, int b, int c) {
    var point = expressionPointers[pointer]

    instructions[point] = a
    instructions[point + 1] = b
    instructions[point + 2] = c
}

func SetPointer(point int) {
    for index,exprPointer := range expressionPointers {
        if (exprPointer >= point) {
            pointer = index
            break
        }
    }
}

func GetInstructionAt(point int) int {
    return instructions[point]
}
