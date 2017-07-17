package src

import (
    "strings"
    "strconv"
    "regexp"
    //"fmt"
)

var Variables = make(map[string]Variable);
var Reserved = []string{"clear", "clean", "exit"};

func Transact(query string) string {
    return eval(query);
}

// Evaluates user input to return
func eval(query string) string {
    query = strings.Replace(query, " ", "", -1);
    if (checkPrimitive(query)) {
        return query;
    }
    if (isVariable(query)) {
        val, contains := Variables[query];
        if (!contains) {
            return "ERROR: Undefined variable: '" + query + "'";
        }
        return Get(val);
    }
    query = replaceSubtraction(query);
    if (strings.Contains(query, "=")) {
        return assignVariable(query);
    } else if (parseArithmetic(query) == 1) {
        return ApplyArithmetic(query);
    } else if (parseArithmetic(query) > 1) {
        return ApplyMultipleArithmetic(query);
    } else if (strings.Contains(query, "[[")) {
        return "";
    } else {
        return "ERROR: Malformed query";
    }
}

// Tests if query contains an arithmetic argument
func parseArithmetic(query string) int {
    matched := CountAny(query, "+","~", "*", "/");
    return matched;
}

// Currently, primitives are defined as integers and floats
func checkPrimitive(query string) bool {
    query = strings.TrimSpace(query);
    _, err := strconv.Atoi(query);
    if (err != nil) {
        _, err2 := strconv.ParseFloat(query, 64)
            if (err2 != nil) {
                return false;
            }
            return true;
        }
    return true;
}

func CountAny(str string, seps ...string) (i int) {
    for _, sep := range(seps) {
        i += strings.Count(str, sep);
    }
    return i;
}

// Replaces "-" with "~" as the subtraction operator so as not to confuse subtraction with a negative sign
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

func isVariable(query string) bool {
    if (len(query) == 1) {
        return true;
    }
    var letter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
    firstLetter := string(query[0]);
    if _, err := strconv.Atoi(firstLetter); err == nil {
        return false; //first letter in a variable cannot be a digit
    }
    return letter(query[1:len(query)]);
}

func assignVariable(query string) string {
    args := strings.Split(query, "=");
    if (len(args) > 2) {
        return "ERROR: Cannot assign more than one variable per statement.";
    }
    LHS := args[0];
    for i := 0; i < len(Reserved); i++ {
        if (Reserved[i] == LHS) {
            return "ERROR: " + LHS + " cannot be used as a variable name"
        }
    }
    RHS := args[1];
    if (!isVariable(LHS)) {
        return "ERROR: Invalid variable name: '" + LHS + "'";
    }
    RHS = eval(RHS);
    ans, err1 := strconv.Atoi(RHS);
    if (err1 != nil) {
        ans1, err2 := strconv.ParseFloat(RHS, 64);
        if (err2 != nil) {
            //TODO: Check if RHS is a matrix
            return "ERROR: Cannot assign variable " + LHS + " to value.";
        } else {
            v := NewVariable("float", 0, ans1, nil);
            Variables[LHS] = v;
        }
    } else {
        v := NewVariable("int", ans, 0, nil);
        Variables[LHS] = v;
    }
    return "";
}
