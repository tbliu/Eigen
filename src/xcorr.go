package src
/** xcorr(x, y) returns the cross-correlation of y with respect to x */

func xcorr(x *Matrix, y *Matrix) (*Matrix, string) {
    if (x.N != 1 || y.N != 1) {
        return nil, "ERROR: Parameters must be vectors";
    }

    if (x.M != y.M) {
        return nil, "ERROR: Vectors must have the same dimensions";
    }

    vals := make([][]float64, y.M);
    for i := 0; i < y.M; i++ {
        v, _ := roll(y, i);
        vals[i] = v.cols[0];
    }
    circulant, _ := NewMatrix(vals);
    return mulMatrices(circulant, x), "";
}
