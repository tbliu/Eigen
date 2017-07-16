package src

import (
    "strings"
    "strconv"
    //"fmt"
)

func Transact(query string) string {
    return eval(query);
}

// Evaluates user input to return
func eval(query string) string {
    query = strings.Replace(query, " ", "", -1);
    if (checkPrimitive(query)) {
        return query;
    }
    query = replaceSubtraction(query);
    if (parseArithmetic(query) == 1) {
        return ApplyArithmetic(query);
    }
    if (parseArithmetic(query) > 1) {
        return ApplyMultipleArithmetic(query);
    }
    return "ERROR: Malformed query";
}

// Tests if query contains an arithmetic argument
func parseArithmetic(query string) int {
    matched := CountAny(query, "+","~", "*", "/");
    return matched;
}

// Currently, primitives are defined as integers
func checkPrimitive(query string) bool {
    query = strings.TrimSpace(query);
    _, err := strconv.Atoi(query);
    if (err != nil) {
        return false;
    }
    return true;
}

func CountAny(str string, seps ...string) (i int) {
    for _, sep := range(seps) {
        i += strings.Count(str, sep);
    }
    return i;
}

// Replaces "-" with "~" as the subtraction operator
// so as not to confuse subtraction with a negative sign
func replaceSubtraction(query string) string {
    firstInst := strings.Index(query, "-");
    if (firstInst == -1) {
        return query;
    }
    if (firstInst != 0 && !strings.ContainsAny(string(query[firstInst - 1]), "+ & * & ~ & /")) {
        query = strings.Replace(query, "-", "~", 1);
        return replaceSubtraction(query);
    } else if (firstInst == 0) { // must be a number
        replaced := []string{"-", replaceSubtraction(query[1:len(query)])};
        return strings.Join(replaced, "");
    } else if (strings.ContainsAny(string(query[firstInst - 1]), "+ & * & ~ & /")) {
        // If prev byte is an operator, current byte must be a number
        replaced := []string{string(query[0]), replaceSubtraction(query[1:len(query)])};
        return strings.Join(replaced, "");
    } else {
        return query;
    }
}
