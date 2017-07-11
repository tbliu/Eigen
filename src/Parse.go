package src

import (
    "regexp"
    "strings"
)

func Transact(query string) string {
    return eval(query);
}

func eval(query string) string {
    query = strings.TrimSpace(query);
    //TODO: Add error handling and create arithmetic operation
    if (parseArithmetic(query)) {
        //if (regexp.
        return "hello";
    }
    return "goodbye";
}
// TODO: Regex not properly telling if there is an operator
// Tests if query contains an arithmetic argument
func parseArithmetic(query string) bool {
    matched, err := regexp.MatchString("+|-|*|/", query);
    if (err != nil) {
        return false;
    }
    return matched;
}
