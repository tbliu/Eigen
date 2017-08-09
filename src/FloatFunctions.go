package src

/** Lists all user-callable functions that return floats */

import (
    "strconv"
)

var FloatFunctions = map[string]bool {
    "det(" : true,
    "sqrt(" : true,
    "dot(" : true,
}

func ApplyFloatFunction(query string) (float64, string) {
    function, params := getFunctionArgs(query);
    if (params == "") {
        return 0.0, "ERROR: Must pass in a parameter";
    }
    switch {
        case function == "det":
            args := paramToMatrix(params);
            if (args == nil) {
                return 0.0, "ERROR: Parameter must be a valid matrix";
            }
            return det(args);
        case function == "sqrt":
            args, err := strconv.ParseFloat(params, 64);
            if (err != nil) {
                return 0.0, "ERROR: Invalid parameter";
            }
            return sqrt(args);
        case function == "dot":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return 0.0, "ERROR: Invalid parameters";
            }
            return dot(arg1, arg2);
        default:
            return 0.0, "ERROR: Invalid function call";
    }
}
