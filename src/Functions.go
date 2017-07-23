package src

import (
    "strings"
    "strconv"
    //"fmt"
)

/** Lists all user-callable functions */

var Functions = map[string]bool {
    "zeros(" : true,
    "id(" : true,
    "rref(" : true,
}

func isFunctionCall(query string) bool {
    if (!strings.Contains(query, "(") || !strings.Contains(query, ")")) {
        return false;
    }
    funcNameIndex := strings.Index(query, "(");
    funcName := query[0:funcNameIndex+1];
    _, match := Functions[funcName];
    if (!match) {
        return false;
    }
    return true;
}

func getFunctionArgs(query string) (string, string) {
    funcNameIndex := strings.Index(query, "(");
    funcName := query[0:funcNameIndex];
    endCallIndex := strings.LastIndex(query, ")");
    parameters := query[funcNameIndex+1:endCallIndex];
    return funcName, parameters;
}

func ApplyFunction(query string) *Matrix {
    function, params := getFunctionArgs(query);
    switch {
        case function == "id":
            params = eval(params);
            i, err := strconv.Atoi(params);
            if (err != nil) {
                return nil;
            }
            return id(i);
        case function == "zeros":
            args := splitParams(params);
            if (args == nil) {
                return nil;
            }
            return zeros(args...);
        case function == "rref":
            return nil;
        default:
            return nil;
        }
}

// splits parameters for variadic functions
func splitParams(params string) []int {
    oldArgs := strings.Split(params, ",");
    args := make([]int, len(oldArgs));
    for i := 0; i < len(oldArgs); i++ {
        j, err := strconv.Atoi(oldArgs[i]);
        if (err != nil) {
            return nil;
        }
        args[i] = j;
    }
    return args;
}
