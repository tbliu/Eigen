package src

/** Returns a matrix B whose columns span the column space of A */

func col(m *Matrix) (*Matrix, string) {
    pivots := getPivotIndices(m);
    numCols := 0;
    for i := 0; i < len(pivots); i++ {
        if (pivots[i] == -1) {
            break;
        }
        numCols += 1;
    }
    values := make([][]float64, numCols);
    // Get the columns
    for i := 0; i < len(pivots); i++ {
        if (pivots[i] == -1) {
            break;
        }
        arr := make([]float64, m.M);
        copy(arr, m.cols[pivots[i]]);
        values[i] = arr;
    }

    // Then transpose since we used the column vectors as row values
    mat, _ := NewMatrix(values);
    colSpace, _ := transpose(mat);
    return colSpace, "";
}


// Returns the indices of the pivot columns of m. Last element is -1 to indicate no more pivots
func getPivotIndices(m *Matrix) []int {
    cols := make([]int, m.N);
    for i := 0; i < len(cols); i++ {
        cols[i] = -1;
    }
    index := 0;
    reduced, _ := rref(m);
    for i := 0; i < len(reduced.rows); i++ {
        for j := 0; j < len(reduced.rows[i]); j++ {
            if (reduced.rows[i][j] == 1) {
                cols[index] = j;
                index += 1;
                break;
            }
        }
    }
    return cols;
}
