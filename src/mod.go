package src

/** Calculates A mod B */
func mod(A int, B int) (int, string) {
    if (A > 0) {
        return A % B, "";
    }

    A = A % B;
    return A + B, "";
}
