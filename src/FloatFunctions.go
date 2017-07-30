package src

/** Lists all user-callable functions that return floats */

var FloatFunctions = map[string]bool {
    "det(" : true,
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
        default:
            return 0.0, "ERROR: Invalid function call";
    }
}
