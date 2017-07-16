package src

import (
    "strconv"
    "strings"
    //"fmt"
)

const EPSILON = 0.01; // For floating point precision

/** Computes basic arithmetic */
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

func findArgs(query string) (string, string, string, string) {
    if (strings.ContainsAny(query, "+")) {
        args := strings.Split(query, "+");
        for i := 0; i < len(args); i++ {
            strings.TrimSpace(args[i]);
        }
        return args[0], "+", args[1], "";
    }
    if (strings.ContainsAny(query, "-")) {
        args := strings.Split(query, "-");
        for i := 0; i < len(args); i++ {
            strings.TrimSpace(args[i]);
        }
        return args[0], "-", args[1], "";
    }
    if (strings.ContainsAny(query, "*")) {
        args := strings.Split(query, "*");
        for i := 0; i < len(args); i++ {
            strings.TrimSpace(args[i]);
        }
        return args[0], "*", args[1], "";
    }
    if (strings.ContainsAny(query, "/")) {
        args := strings.Split(query, "/");
        for i := 0; i < len(args); i++ {
            strings.TrimSpace(args[i]);
        }
        return args[0], "/", args[1], "";
    }
    return "","","","ERROR: Invalid Number";
}

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
