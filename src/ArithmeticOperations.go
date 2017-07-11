package main

/** Computes basic arithmetic */
type Operation struct{
    operator string;
}

func NewOperation(operator string) *Operation {
    p := new(Operation);
    p.operator = operator;
}

func applyInts(p Operation, x int, y int) int, error {
    if (p.operator == "+") {
        return x + y, nil;
    } else if (p.operator == "-") {
        return x - y, nil;
    } else if (p.operator == "*") {
        return x * y, nil;
    } else if (p.operator == "/") {
        if (y == 0) {
            err = "ERROR: Cannot divide by 0";
            return 0, err;
        }
        return x/y, nil;
    } else {
        return 0, "ERROR: Invalid operand";
    }
}

func main() {}
