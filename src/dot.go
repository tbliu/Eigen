package src

/** Computes the dot product of two vectors */
func dot(v *Matrix, u *Matrix) (float64, string) {
    if (v.N != 1 || u.N != 1) {
        return nil, "ERROR: Arguments must be vectors";
    }
    if (v.M != u.M) {
        return nil, "ERROR: Vector dimensions must match";
    }
    result := 0.0;
    for (i := 0; i < v.cols; i++) {
        arr := make([]float64, 1);
        result += v.cols[i][0] * u.cols[i][0];
    }

    return result, "";
}
