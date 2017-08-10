package src

/** Lists all user-callable functions that return ints */

import (
    "strings"
    "strconv"
)

var IntFunctions = map[string]bool {
    "rank(" : true,
    "nullity(" : true,
    "gcd(" : true,
    "orth(": true,
    "modinv(":true,
    "mod(": true,
    "modexp(":true,
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
        case function == "modinv":
            args := splitIntegers(params);
            if (args == nil) {
                return 0, "ERROR: Invalid parameters";
            }
            return modinv(args[0], args[1]);
        case function == "mod":
            args := splitIntegers(params);
            if (args == nil) {
                return 0, "ERROR: Invalid parameters"
            }
            return mod(args[0], args[1]);
        case function == "modexp":
            args := splitIntegers(params);
            if (args == nil) {
                return 0, "ERROR: Invalid parameters";
            }
            return modexp(args[0], args[1], args[2]);
        default:
            return 0, "ERROR: Invalid function call";
    }
}

func splitIntegers(query string) []int {
    oldArgs := strings.Split(query, ",");
    if (len(oldArgs) < 2 || len(oldArgs) > 3) {
        return nil;
    }
    first, e := strconv.Atoi(oldArgs[0]);
    second, e2 := strconv.Atoi(oldArgs[1]);
    if (e != nil || e2 != nil) {
        return nil;
    }
    if (len(oldArgs) == 3) {
        third, e3 := strconv.Atoi(oldArgs[2]);
        if (e3 != nil) {
            return nil;
        }

        return []int{first, second, third};
    }
    return []int{first, second};
}
