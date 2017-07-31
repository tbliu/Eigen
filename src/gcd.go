package src

/** Calculates the greatest common divisor of a and b */
func gcd(a int, b int) (int, string) {
    if (a == b) {
        return a, "";
    } else if (a > b) {
        return gcdHelper(a, b), "";
    } else {
        return gcdHelper(b, a), "";
    }
}

func gcdHelper(larger int, smaller int) int {
    if (smaller == 0) {
        return larger;
    }
    return gcdHelper(smaller, larger % smaller);
}
