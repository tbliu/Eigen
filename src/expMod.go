package src

/** Finds the result of x^y mod m using repeated squaring */
func modexp(x int, y int, m int) (int, string) {
    if (y == 0) {
        return 1, "";
    }
    z, _ := modexp(x, y / 2, m);
    if (y % 2 == 0) {
        return (z * z) % m, "";
    }
    return (x * z * z) % m, "";
}
