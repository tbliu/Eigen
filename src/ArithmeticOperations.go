package src

import (
    "strconv"
    "strings"
    //"fmt"
)

/** Computes basic arithmetic. Supports ints as well as 64 bit floats */

// Returns final value for arithmetic operation for multiple operators
func ApplyMultipleArithmetic(query string) string {
    firstOperator := strings.IndexAny(query, "+~/*");
    if (CountAny(query, "+", "~", "*", "/") == 1) {
        return ApplyArithmetic(query);
    }
    firstArgument := query[0:firstOperator+1];
    query = query[firstOperator+1:len(query)];
    secondArgIndex := strings.IndexAny(query, "+~/*");
    secondArgument := query[0:secondArgIndex];
    query = query[secondArgIndex:len(query)];
    totalArgs := []string{firstArgument, secondArgument};
    result := ApplyArithmetic(strings.Join(totalArgs, ""));
    newArgs := []string{result, query};
    return ApplyMultipleArithmetic(strings.Join(newArgs, ""));
}

// Returns final value for arithmetic operation
func ApplyArithmetic(query string) string {
    first, operator, second, argErr := findArgs(query);
    if (argErr != "") {
        return argErr;
    }
    if (strings.ContainsAny(first, ".") || strings.ContainsAny(second, ".")) {
        num1, err1 := strconv.ParseFloat(first, 64);
        num2, err2 := strconv.ParseFloat(second, 64);
        if (err1 != nil || err2 != nil) {
            return "Error: Invalid number";
        }
        result, err := applyOperatorFloat(operator, num1, num2);
        if (err == "") {
            return result;
        }
        return err;
    }
    arg1, err1 := strconv.Atoi(first);
    arg2, err2 := strconv.Atoi(second);
    if (err1 != nil || err2 != nil) {
        return "ERROR: Invalid number";
    }
    result, err := applyOperator(operator, arg1, arg2);
    if (err == "") {
        return result;
    }
    return err;
}

// Parses user query to find operands and operator
func findArgs(query string) (string, string, string, string) {
    if (strings.ContainsAny(query, "+")) {
        args := strings.Split(query, "+");
        if (IsVariable(args[0])) {
            if (Defined(args[0])) {
                if (Variables[args[0]].class != "matrix") {
                    args[0] = ToString(Variables[args[0]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[0] + "'";
            }
        }
        if (IsVariable(args[1])) {
            if (Defined(args[1])) {
                if (Variables[args[1]].class != "matrix") {
                    args[1] = ToString(Variables[args[1]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[1] + "'";
            }
        }
        return args[0], "+", args[1], "";
    }
    if (strings.ContainsAny(query, "~")) {
        args := strings.Split(query, "~");
        if (IsVariable(args[0])) {
            if (Defined(args[0])) {
                if (Variables[args[0]].class != "matrix") {
                    args[0] = ToString(Variables[args[0]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[0] + "'";
            }
        }
        if (IsVariable(args[1])) {
            if (Defined(args[1])) {
                if (Variables[args[1]].class != "matrix") {
                    args[1] = ToString(Variables[args[1]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[1] + "'";
            }
        }
        return args[0], "-", args[1], "";
    }
    if (strings.ContainsAny(query, "*")) {
        args := strings.Split(query, "*");
        if (IsVariable(args[0])) {
            if (Defined(args[0])) {
                if (Variables[args[0]].class != "matrix") {
                    args[0] = ToString(Variables[args[0]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[1] + "'";
            }
        }
        if (IsVariable(args[1])) {
            if (Defined(args[1])) {
                if (Variables[args[1]].class != "matrix") {
                    args[1] = ToString(Variables[args[1]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[1] + "'";
            }
        }
        return args[0], "*", args[1], "";
    }
    if (strings.ContainsAny(query, "/")) {
        args := strings.Split(query, "/");
        if (IsVariable(args[0])) {
            if (Defined(args[0])) {
                if (Variables[args[0]].class != "matrix") {
                    args[0] = ToString(Variables[args[0]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[0] + "'";
            }
        }
        if (IsVariable(args[1])) {
            if (Defined(args[1])) {
                if (Variables[args[1]].class != "matrix") {
                    args[1] = ToString(Variables[args[1]]);
                }
            } else {
                return "","","","ERROR: Invalid variable '" + args[0] + "'";
            }
        }
        return args[0], "/", args[1], "";
    }
    return "","","","ERROR: Invalid Number";
}

// Applies operator for ints
func applyOperator(operator string, x int, y int) (string, string) {
    if (operator == "+") {
        return strconv.Itoa(x + y), "";
    } else if (operator == "-") {
        return strconv.Itoa(x - y), "";
    } else if (operator == "*") {
        return strconv.Itoa(x * y), "";
    } else if (operator == "/") {
        if (y == 0) {
            err := "ERROR: Cannot divide by 0";
            return "0", err;
        }
        return strconv.Itoa(x / y), "";
    } else {
        return "0", "ERROR: Invalid operand";
    }
}

// Applies operator for floats
func applyOperatorFloat(operator string, x float64, y float64) (string, string) {
    if (operator == "+") {
        return strconv.FormatFloat(x + y, 'f', 2, 64), "";
    } else if (operator == "-") {
        return strconv.FormatFloat(x - y, 'f', 2, 64), "";
    } else if (operator == "*") {
        return strconv.FormatFloat(x * y, 'f', 2, 64), "";
    } else if (operator == "/") {
        if (y == 0) {
            err := "ERROR: Cannot divide by 0";
            return "0", err;
        }
        return strconv.FormatFloat(x / y, 'f', 2, 64), "";
    } else {
        return "0", "ERROR: Invalid operand";
    }
}
