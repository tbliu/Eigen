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
    "transpose(" : true,
    "inv(" : true,
    "col(" : true,
    "row(" : true,
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

func ApplyFunction(query string) (*Matrix, string) {
    function, params := getFunctionArgs(query);
    if (params == "") {
        return nil, "ERROR: Must pass in a parameter";
    }
    switch {
        case function == "id":
            params = eval(params);
            i, err := strconv.Atoi(params);
            if (err != nil) {
                return nil, "ERROR: Parameter must be an int";
            }
            matrix, e := id(i);
            return matrix, e;
        case function == "zeros":
            args := splitParams(params);
            if (args == nil) {
                return nil, "ERROR: zeros can only take in one or two arguments";
            }
            return zeros(args...);
        case function == "rref":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return rref(args);
        case function == "transpose":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return transpose(args);
        case function == "inv":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return inv(args);
        case function == "col":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return col(args);
        case function == "row":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return row(args);
        default:
            return nil, "ERROR: Invalid function call";
        }
}

func paramToMatrix(params string) *Matrix {
    var args *Matrix = nil;
    if (IsVariable(params)) {
        variable := Variables[params];
        if (variable.class != "matrix") {
            return nil;
        }
        args = variable.matrix;
    } else if (isFunctionCall(params)) {
        args, _ = ApplyFunction(params);
    } else {
        vals := queryToValues(params);
        args, _ = NewMatrix(vals);
    }
    return args;
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
