package src

import (
    "strings"
    "strconv"
    //"fmt"
)

/** Lists all user-callable functions that return matrices */

var Functions = map[string]bool {
    "zeros(" : true,
    "id(" : true,
    "rref(" : true,
    "transpose(" : true,
    "inv(" : true,
    "col(" : true,
    "row(" : true,
    "null(": true,
    "Lnull(" : true,
    "solve(" : true,
    "hsolve(" : true,
    "llse(" : true,
    "proj(" : true,
    "norm(" : true,
    "gs(" : true,
    "qr(" : true,
    "roll(": true,
    "xcorr(": true,
    "autocorr(": true,
}

func isFunctionCall(query string) bool {
    if (!strings.Contains(query, "(") || !strings.Contains(query, ")")) {
        return false;
    }
    funcNameIndex := strings.Index(query, "(");
    funcName := query[0:funcNameIndex+1];
    _, match := Functions[funcName];
    _, match2 := IntFunctions[funcName];
    _, match3 := FloatFunctions[funcName];
    if (!match && !match2 && !match3) {
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
        case function == "null":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return null(args);
        case function == "Lnull":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return Lnull(args);
        case function == "solve":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return nil, "ERROR: Invalid parameters."
            }
            return solve(arg1, arg2);
        case function == "hsolve":
            args := paramToMatrix(params);
            if (args == nil) {
                return nil, "ERROR: Parameter must be a valid matrix";
            }
            return hsolve(args);
        case function == "llse":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return nil, "ERROR: Invalid parameters";
            }
            return llse(arg1, arg2);
        case function == "proj":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return nil, "ERROR: Invalid parameters";
            }
            return proj(arg1, arg2);
        case function == "norm":
            arg := paramToMatrix(params);
            if (arg == nil) {
                return nil, "ERROR: Invalid parameter";
            }
            return norm(arg);
        case function == "gs":
            arg := paramToMatrix(params);
            if (arg == nil) {
                return nil, "ERROR: Invalid parameter";
            }
            return gs(arg);
        case function == "qr":
            arg1, arg2 := splitTag(params);
            if (arg1 == nil) {
                return nil, "ERROR: Invalid parameters";
            }
            return qr(arg1, arg2);
        case function == "roll":
            arg1, arg2 := splitInt(params);
            if (arg1 == nil) {
                return nil, "ERROR: Invalid parameters";
            }
            return roll(arg1, arg2);
        case function == "xcorr":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return nil, "ERROR: Invalid parameters";
            }
            return xcorr(arg1, arg2);
        case function == "autocorr":
            arg := paramToMatrix(params);
            if (arg == nil) {
                return nil, "ERROR: Invalid parameter";
            }
            return autocorr(arg);
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

// Splits parameters into two matrices
func splitMatrices(params string) (*Matrix, *Matrix) {
    oldArgs := strings.Split(params, ",");
    if (len(oldArgs) != 2) {
        return nil, nil;
    }
    mat1 := paramToMatrix(oldArgs[0]);
    mat2 := paramToMatrix(oldArgs[1]);
    return mat1, mat2;

}

func splitTag(params string) (*Matrix, string) {
    oldArgs := strings.Split(params, ",");
    if (len(oldArgs) != 2) {
        return nil, "";
    }
    matrix := paramToMatrix(oldArgs[0]);
    tag := oldArgs[1];
    return matrix, tag;
}

func splitInt(params string) (*Matrix, int) {
    oldArgs := strings.Split(params, ",");
    if (len(oldArgs) != 2) {
        return nil, 0;
    }
    matrix := paramToMatrix(oldArgs[0]);
    n, e := strconv.Atoi(oldArgs[1]);
    if (e != nil) {
        return nil, 0;
    }
    return matrix, n;
}
