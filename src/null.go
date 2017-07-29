package src

/** Returns a matrix B whose columns span the null space of A */

func null(m *Matrix) (*Matrix, string) {
    reduced, _ := rref(m);
    if (checkIfIdentity(reduced)) {
        return nil, "ERROR: Null space is empty";
    }
    // Rank-Nullity theorem
    size := reduced.N - rank(reduced);
    values := make([][]float64, size);
    index := 0;
    freeColumns := getFreeColumns(reduced);
    for i := 0; i < len(freeColumns); i++ {
        arr := make([]float64, m.N);
        for j := 0; j < len(m.rows); j++ {
            arr[j] = -1.0 * m.rows[j][freeColumns[i]];
        }
        arr[freeColumns[i]] = 1;
        values[index] = arr;
        index += 1;
    }
    transposedNull, _ := NewMatrix(values);
    round(transposedNull);
    return transpose(transposedNull);
}

// Return indices of free columns in a reduced matrix
func getFreeColumns(m *Matrix) []int {
    pivots := make(map[int]bool);
    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            if (m.rows[i][j] == 1) {
                pivots[j] = true;
                break;
            }
        }
    }
    size := m.N - rank(m);
    free := make([]int, size);
    index := 0;
    for i := 0; i < len(m.cols); i++ {
        _, isPivot := pivots[i];
        if (!isPivot) {
            free[index] = i;
            index += 1;
        }
    }
    return free;
}

func checkIfIdentity(m *Matrix) bool {
    if (m.N != m.M) {
        return false;
    }

    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            if (i == j) {
                if (m.rows[i][j] != 1) {
                    return false;
                }
            } else {
                if (m.rows[i][j] != 0) {
                    return false;
                }
            }
        }
    }
    return true;
}
