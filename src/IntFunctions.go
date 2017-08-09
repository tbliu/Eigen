package src

/** Lists all user-callable functions that return ints */

var IntFunctions = map[string]bool {
    "rank(" : true,
    "nullity(" : true,
    "gcd(" : true,
    "orth(": true,
}

func ApplyIntFunction(query string) (int, string) {
    function, params := getFunctionArgs(query);
    if (params == "") {
        return 0, "ERROR: Must pass in a parameter";
    }
    switch {
        case function == "rank":
            args := paramToMatrix(params);
            if (args == nil) {
                return 0, "ERROR: Parameter must be a valid matrix";
            }
            return rank(args), "";
        case function == "nullity":
            args := paramToMatrix(params);
            if (args == nil) {
                return 0, "ERROR: Parameter must be a valid matrix";
            }
            return nullity(args), "";
        case function == "gcd":
            args := splitParams(params);
            if (args == nil) {
                return 0, "ERROR: Parameter must be an int";
            }

            if (len(args) != 2) {
                return 0, "ERROR: gcd must take two inputs";
            }
            return gcd(args[0], args[1]);
        case function == "orth":
            arg1, arg2 := splitMatrices(params);
            if (arg1 == nil || arg2 == nil) {
                return 0, "ERROR: Invalid vectors";
            }
            return orth(arg1, arg2);
        default:
            return 0, "ERROR: Invalid function call";
    }
}
