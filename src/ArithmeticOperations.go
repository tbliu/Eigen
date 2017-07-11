package src

/** Computes basic arithmetic */
type Operation struct{
    operator string;
}

func NewOperation(operator string) *Operation {
    p := new(Operation);
    p.operator = operator;
    return p;
}

func Apply(p *Operation, x int, y int) (int, string) {
    if (p.operator == "+") {
        return x + y, "";
    } else if (p.operator == "-") {
        return x - y, "";
    } else if (p.operator == "*") {
        return x * y, "";
    } else if (p.operator == "/") {
        if (y == 0) {
            err := "ERROR: Cannot divide by 0";
            return 0, err;
        }
        return x/y, "";
    } else {
        return 0, "ERROR: Invalid operand";
    }
}

func main(){}
