package src

/** Computes the dot product of two vectors */
func dot(v *Matrix, u *Matrix) (float64, string) {
    if (v.N != 1 || u.N != 1) {
        return 0.0, "ERROR: Arguments must be vectors";
    }
    if (v.M != u.M) {
        return 0.0, "ERROR: Vector dimensions must match";
    }
    result := 0.0;
    for i := 0; i < len(v.rows); i++ {
        result += v.cols[0][i] * u.cols[0][i];
    }

    return result, "";
}
