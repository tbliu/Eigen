package src

/** Returns a vector that solves the equation Ax = b */
func solve(m *Matrix, b *Matrix) (*Matrix, string) {
    aug, err := augment(m, b);
    augmentIndex := m.N;
    if (err != "") {
        return nil, err;
    }
    reduced, _ := rref(aug);
    if (!validateSolution(reduced)) {
        return nil, "ERROR: There is no solution";
    }
    solution := make([][]float64, m.N); // Solving an MxN matrix yields a solution in R^N
    index := 0;
    for i := 0; i < len(reduced.rows); i++ {
        if (index >= len(solution)) {
            break;
        }
        arr := make([]float64, 1);
        for j := 0; j < len(arr); j++ {
            arr[j] = reduced.rows[i][augmentIndex];
        }
        solution[index] = arr;
        index += 1;
    }
    mat, err := NewMatrix(solution);
    if (err != "") {
        return nil, err;
    }
    return mat, "";
}

// If there is a row of zeros and the augmented entry is not zero, there is no solution
// ex: [0,0,0,0 | 4]
func validateSolution(reduced *Matrix) bool {
    augmentedEntry := reduced.N - 1;
    for i := 0; i < len(reduced.rows); i++ {
        for j := 0; j < len(reduced.rows[i]); j++ {
            if (reduced.rows[i][j] != 0 && j < augmentedEntry) {
                break;
            }

            if (j == augmentedEntry && reduced.rows[i][j] != 0) {
                return false;
            }
        }
    }
    return true;
}
