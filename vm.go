package main

import (
    "fmt"
    "./memory"
)

var broken = false

func main() {
    fmt.Println("Running")

    for {
        evaluateExpr()
        if (broken) {
            break
        }
    }

    fmt.Println("")
}

func evaluateExpr() {
    a, b, c := memory.GetExpression()

    if a >= 0 {
        a = memory.GetInstructionAt(a)
    }

    if b >= 0 {
        b = memory.GetInstructionAt(b)
    }

    if b == -1 {
        output(a)
        memory.SetPointer(c)
        return
    }


    b -= a
    if (b <= 0) {
        if c < 0 {
            broken = true
            return
        }
        fmt.Println("Moving Pointer")
        memory.SetPointer(c)
    }

    broken = false
}

func output(char int) {
    fmt.Printf("%c", char)
}
