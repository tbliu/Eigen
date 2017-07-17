package src

import (
    "regexp"
    "strconv"
    "strings"
)

/** General file for miscellaneous commands */
func Clean() {
    Variables = make(map[string]Variable);
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
