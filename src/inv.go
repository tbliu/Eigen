package src

/** Calculates the inverse of a matrix. Matrix must be full rank and square */

func inv(m *Matrix) (*Matrix, string) {
    if (m.N != m.M) {
        return nil, "ERROR: Matrix must be square";
    }

    original := copyMatrix(m);
    det, _ := det(original);
    if (det == 0) {
        return nil, "ERROR: Matrix is not invertible";
    }

    inverse, _ := id(m.N);
    for i := 0; i < len(original.rows); i++ {
        var factor float64 = 0;
        original.rows[i], factor = invertRow(original.rows[i]);
        inverse.rows[i] = scaleRow(inverse.rows[i], factor);
        row := original.rows[i];
        pivotIndex := getPivotIndex(row);

        for j := 0; j < len(original.rows); j++ {
            if (i == j || original.rows[j][pivotIndex] == 0) {
                continue;
            }

            scaleFactor := original.rows[j][pivotIndex] * -1.0;
            newRow := addRows(scaleRow(original.rows[i], scaleFactor), original.rows[j]);
            original.rows[j] = newRow;
            newInverseRow := addRows(scaleRow(inverse.rows[i], scaleFactor), inverse.rows[j]);
            inverse.rows[j] = newInverseRow;
        }
    }
    swapRowsWithInverse(original, inverse);
    round(inverse);
    inverse.cols = valsToCols(inverse.rows);
    return inverse, "";
}

func swapRowsWithInverse(original *Matrix, inverse *Matrix) {
    for i := 0; i < len(original.rows); i++ {
        pivotIndex := getPivotIndex(original.rows[i]);

        for j := i; j < len(original.rows); j++ {
            if (i == j) {
                continue
            }

            otherPivotIndex := getPivotIndex(original.rows[j]);
            if ((pivotIndex > otherPivotIndex && otherPivotIndex != -1) || pivotIndex == -1) {
                swap(original, i, j);
                swap(inverse, i, j);
                i = 0;
            }
        }
    }
}
