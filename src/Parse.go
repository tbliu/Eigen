package src

import (
    //"regexp"
    "strings"
    "strconv"
)

func transact(query string) string {
    return eval(query);
}

// Evaluates user input to return
func eval(query string) string {
    query = strings.Replace(query, " ", "", -1);
    if (checkPrimitive(query)) {
        return query;
    }
    if (parseArithmetic(query)) {
        return ApplyArithmetic(query);
    }
    return "ERROR: Malformed query";
}

// Tests if query contains an arithmetic argument
func parseArithmetic(query string) bool {
    matched := strings.ContainsAny(query, "+ & - & * & /");
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

