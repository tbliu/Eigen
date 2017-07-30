package src

/** Lists all user-callable functions that return ints */

var IntFunctions = map[string]bool {
    "rank(" : true,
    "nullity(" : true,
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
        default:
            return 0, "ERROR: Invalid function call";
    }
}
