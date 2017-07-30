package src

import (
    "regexp"
    "strconv"
    "strings"
    //"fmt"
)

/** General file for miscellaneous commands */
func Clean() {
    Variables = make(map[string]Variable);
}

func IsVariable(query string) bool {
    for i := 0; i < len(Reserved); i++ {
        if (query == Reserved[i]) {
            return false;
        }
    }
    if (len(query) == 1) {
        _, errInt := strconv.Atoi(query);
        if (errInt == nil) {
            return false;
        }
        _, errFloat := strconv.ParseFloat(query, 64);
        if (errFloat == nil) {
            return false;
        }
        return true;
    }
    var letter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
    firstLetter := string(query[0]);
    if _, err := strconv.Atoi(firstLetter); err == nil {
        return false; //first letter in a variable cannot be a digit
    }
    return letter(query[1:len(query)]);
}

func IsMatrix(query string) (*Matrix, bool) {
    if IsVariable(query) {
        v, match := Variables[query];
        if (!match) {
            return nil, false;
        } else if (v.class == "matrix") {
            return v.matrix, true;
        } else {
            return nil, false;
        }
    } else if (string(query[0]) == "[") {
        rows := queryToValues(query);
        matrix, err := NewMatrix(rows);
        if (err != "") {
            return nil, false;
        }
        return matrix, true;
    } else {
        return nil, false;
    }
}

// checks of query contains a matrix literal or variable
func containsMatrix(query string) bool {
    _, match := IsMatrix(query);
    if (match) {
        return true;
    }
    f := func(c rune) bool {
        return c == '+' || c == '~' || c == '*' || c == '/';
    }
    split := strings.FieldsFunc(query, f);
    for i := 0; i < len(split); i++ {
        _, match = IsMatrix(split[i]);
        if (match) {
            return true;
        }
    }
    return false;
}

func assignVariable(query string) string {
    count := strings.Count(query, "=");
    if (count > 1) {
        return "ERROR: Cannot have more than one '=' per statement";
    }
    equalIndex := strings.Index(query, "=");
    if (equalIndex + 1 >= len(query)) {
        return "ERROR: Malformed query";
    }
    RHS := query[equalIndex+1:len(query)];
    LHS := query[0:equalIndex];
    if (!IsVariable(LHS)) {
        return "ERROR: Invalid variable name: '" + LHS + "'";
    }
    if (isFunctionCall(RHS)) {
        funcNameIndex := strings.Index(RHS, "(");
        funcName := RHS[0:funcNameIndex+1]
        _, match := Functions[funcName];
        if (match) {
            matrix, err := ApplyFunction(RHS);
            if (err != "") {
                return "ERROR: Cannot assign variable to invalid value";
            }
            v := NewVariable("matrix", 0, 0, matrix);
            Variables[LHS] = v;
            return "";
        }
        _, match = IntFunctions[funcName];
        if (match) {
            val, err := ApplyIntFunction(RHS);
            if (err != "") {
                return "ERROR: Cannot assign variable to invalid value";
            }
            v := NewVariable("int", val, 0.0, nil);
            Variables[LHS] = v;
            return ""
        }

        _, match = FloatFunctions[funcName];
        if (match) {
            val, err := ApplyFloatFunction(RHS);
            if (err != "") {
                return "ERROR: Cannot assign variable to invalid value";
            }
            v := NewVariable("float", 0, val, nil);
            Variables[LHS] = v;
            return "";
        }
    }
    if (containsMatrix(RHS)) {
        count := CountAny(query, "+", "~", "*", "/");
        if (count == 0) {
            matrix, _ := IsMatrix(RHS);
            v := NewVariable("matrix", 0, 0, matrix);
            LHS := query[0:equalIndex];
            if (!IsVariable(LHS)) {
                return "ERROR: Invalid variable name: '" + LHS + "'";
            }
            Variables[LHS] = v;
            return "";
        } else {
            mat, _ := ApplyMultipleMatrixOperations(RHS);
            LHS := query[0:equalIndex];
            if (!IsVariable(LHS)) {
                return "ERROR: Invalid variable name: '" + LHS + "'";
            }
            v := NewVariable("matrix", 0, 0, mat);
            Variables[LHS] = v;
            return "";
        }
    }
    RHS = Transact(RHS);
    query = query[0:equalIndex+1];
    newQuery := []string{query, RHS};
    query = strings.Join(newQuery, "");
    args := strings.Split(query, "=");
    if (len(args) < 2) {
        return "ERROR: Malformed query";
    }
    LHS = args[0];
    RHS = args[1];
    if (!IsVariable(LHS)) {
        return "ERROR: Invalid variable name: '" + LHS + "'";
    }
    RHS = eval(RHS);
    ans, err1 := strconv.Atoi(RHS);
    if (err1 != nil) {
        ans1, err2 := strconv.ParseFloat(RHS, 64);
        if (err2 != nil) {
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
