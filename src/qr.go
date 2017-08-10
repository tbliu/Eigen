package src
/** qr(m *Matrix, _) returns the QR Decomposition of a matrix m.
    * setting _ to 'q' returns Q, while setting _ to 'r' returns R
*/

func qr(A *Matrix, tag string) (*Matrix, string) {
    if (tag != "q" && tag != "r" && tag != "\"q\"" && tag != "'q'" && tag != "\"r\"" && tag != "'r'") {
        return nil, "ERROR: Invalid tag. Please use either q or r";
    }

    Q, _ := gs(A);
    if (tag == "q") {
        return Q, "";
    } else {
        // Create the transposed R matrix
        Rvals := make([][]float64, Q.N);
        for i := 0; i < len(Rvals); i++ {
            column := make([]float64, Q.N);
            for j := 0; j <= i; j++ {
                column[j] = dotCols(A.cols[i], Q.cols[j]);
            }
            Rvals[i] = column;
        }
        Rtrans, _ := NewMatrix(Rvals);
        round(Rtrans);
        return transpose(Rtrans);
    }
}

func dotCols(a []float64, q []float64) float64 {
    result := 0.0;
    for i := 0; i < len(a); i++ {
        result += a[i] * q[i];
    }
    return result;
}
