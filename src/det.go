package src

/** Computes the determinant of a matrix 
*  Matrix must be square.
*/

func det(m *Matrix) (float64, string) {
    if (m.N != m.M) {
        return 0.0, "ERROR: Matrix must be square";
    }
    reduced := copyMatrix(m);
    det := findDet(reduced, 1);
    return det, "";
}

/** Method:
    * det of identity matrix == 1
    * Scaling a row by x scales the determinant by x
    * Swapping a row scales the determinant by -1
    * Adding a scaled row to another row doesn't change the determinant

    This function is a modified version of rref.go -- it reduces m while keeping track of what row operations have been performed.
*/
func findDet(m *Matrix, currDet float64) float64 {
    for i := 0; i < len(m.rows); i++ {
        factor := 0.0;
        m.rows[i], factor = invertRow(m.rows[i]);
        row := m.rows[i];
        currDet *= 1.0/factor;
        pivotIndex := getPivotIndex(row);
        if pivotIndex == -1 {
            continue;
        }

        for j:= 0; j < len(m.rows); j++ {
            if (i == j || m.rows[j][pivotIndex] == 0) {
                continue;
            }

            scaleFactor := m.rows[j][pivotIndex] * -1.0;
            newRow := addRows(scaleRow(m.rows[i], scaleFactor), m.rows[j]);
            m.rows[j] = newRow;
        }
    }
    currDet = swapRowsDet(m, currDet);
    // Multiply entries along the diagonal
    for i := 0; i < len(m.rows); i++ {
        currDet *= m.rows[i][i];
    }
    return currDet;
}

func swapRowsDet(m *Matrix, currDet float64) float64 {
    for i := 0; i < len(m.rows); i++ {
        pivotIndex := getPivotIndex(m.rows[i]);

        for j := i; j < len(m.rows); j++ {
            if (i == j) {
                continue;
            }

            otherPivotIndex := getPivotIndex(m.rows[j]);
            if (otherPivotIndex == -1) {
                continue;
            }
            if ((pivotIndex > otherPivotIndex) || (pivotIndex == -1)) {
                swap(m, i, j);
                i = 0;
                currDet *= -1;
            }
        }
    }
    return currDet;
}
