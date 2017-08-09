package src
/** gs(m *Matrix) applies Gram-Schmidt onto the matrix m to create an orthogonal matrix
  * with orthogonal columns
*/

func gs(m *Matrix) (*Matrix, string) {
    vals := make([][]float64, m.N);
    for i := 0; i < len(m.cols); i++ {
        col := make([]float64, m.M);
        copy(col, m.cols[i]);
        for j := 0; j < i; j++ {
            projection := projCols(vals[j], m.cols[i]);
            col = subtractCols(col, projection);
        }
        vals[i] = col;
    }
    // Normalize the columns
    for i := 0; i < len(vals); i++ {
        normalize(vals[i]);
    }
    // Create the transposed matrix using the new column values
    matrix, e := NewMatrix(vals);
    if (e != "") {
        return nil, e;
    }
    round(matrix);
    return transpose(matrix);
}

/** Computes the projection of w onto v
  * Note: proj.go computes the projection for the matrix struct. 
  * This helper function is designed to compute the projection for 
    two arrays to be made into a matrix later.
*/
func projCols(v []float64, w []float64) []float64 {
    if (len(v) != len(w)) {
        return nil;
    }
    projection := make([]float64, len(v));
    numer := 0.0;
    denom := 0.0;
    for i := 0; i < len(v); i++ {
        numer += v[i] * w[i];
        denom += v[i] * v[i];
    }
    factor := numer/denom;
    for i := 0; i < len(v); i++ {
        projection[i] = factor * v[i];
    }
    return projection;
}

func subtractCols(v []float64, w []float64) []float64 {
    if (len(v) != len(w)) {
        return nil;
    }
    orthVector := make([]float64, len(v));
    for i := 0; i < len(v); i++ {
        orthVector[i] = v[i] - w[i];
    }
    return orthVector;
}

func normalize(v []float64) {
    factor := 0.0;
    for i := 0; i < len(v); i++ {
        factor += v[i] * v[i];
    }

    normFactor, _ := sqrt(factor);
    for i := 0; i < len(v); i++ {
        v[i] /= normFactor;
    }
}
