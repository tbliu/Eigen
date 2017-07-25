package src

import "math"

/** Computes the row reduced echelon form of a matrix */

const epsilon = 0.005; // rounding factor for floating point precision

func rref(m *Matrix) (*Matrix, string) {
    reduced := copyMatrix(m);
    if (checkIfZeroMatrix(reduced)) {
        return reduced, "";
    }
    for i := 0; i < len(reduced.rows); i++ {
        reduced.rows[i] = invertRow(reduced.rows[i]);
        row := reduced.rows[i];
        pivotIndex := getPivotIndex(row);
        if pivotIndex == -1 {
            continue;
        }

        for j := 0; j < len(reduced.rows); j++ {
            // Don't add row to itself. If row element is 0, ignore it
            if (i == j || reduced.rows[j][pivotIndex] == 0) {
                continue;
            }

            scaleFactor := reduced.rows[j][pivotIndex] * -1.0;
            newRow := addRows(scaleRow(reduced.rows[i], scaleFactor), reduced.rows[j]);
            reduced.rows[j] = newRow;
        }
    }
    swapRows(reduced);
    round(reduced);
    reduced.cols = valsToCols(reduced.rows);
    return reduced, "";
}

func swapRows(m *Matrix) {
    for i := 0; i < len(m.rows); i++ {
        pivotIndex := getPivotIndex(m.rows[i]);

        for j := i; j < len(m.rows); j++ {
            if (i == j) {
                continue;
            }

            otherPivotIndex := getPivotIndex(m.rows[j]);
            if ((pivotIndex > otherPivotIndex && otherPivotIndex != -1) || pivotIndex == -1) {
                swap(m, i, j);
                i = 0;
            }
        }
    }
}

func invertRow(row []float64) []float64 {
    for i := 0; i < len(row); i++ {
        if (row[i] != 0) {
            factor := 1.0/row[i];
            row = scaleRow(row, factor);
            return row;
        }
    }
    return row;
}

func getPivotIndex(row []float64) int {
    for i := 0; i < len(row); i++ {
        if row[i] == 1 {
            return i;
        }
    }
    return -1; // row of zeros -- no pivot
}

func checkIfZeroMatrix(m *Matrix) bool {
    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            if (m.rows[i][j] != 0) {
                return false;
            }
        }
    }
    return true;
}

// removes -0 literal from matrix
func round(m *Matrix) {
    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            if (math.Abs(m.rows[i][j]) < epsilon) {
                m.rows[i][j] = 0;
            }
        }
    }
}
