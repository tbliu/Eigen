package src

import "strconv"

/** Computes basic arithmetic */
type Operation struct{
    operator string;
}

func NewOperation(operator string) *Operation {
    p := new(Operation);
    p.operator = operator;
    return p;
}

func Apply(p *Operation, x int, y int) (string, string) {
    if (p.operator == "+") {
        return strconv.Itoa(x + y), "";
    } else if (p.operator == "-") {
        return strconv.Itoa(x - y), "";
    } else if (p.operator == "*") {
        return strconv.Itoa(x * y), "";
    } else if (p.operator == "/") {
        if (y == 0) {
            err := "ERROR: Cannot divide by 0";
            return "0", err;
        }
        return strconv.Itoa(x / y), "";
    } else {
        return "0", "ERROR: Invalid operand";
    }
}
