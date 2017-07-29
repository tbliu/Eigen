package src

/** Returns a vector that solves the equation Ax = b */
func solve(m *Matrix, b *Matrix) (*Matrix, string) {
    aug, err := augment(m, b);
    augmentIndex := m.N;
    if (err != "") {
        return nil, err;
    }
    reduced, _ := rref(aug);
    solution := make([][]float64, m.N); // Solving an MxN matrix yields a solution in R^N
    index := 0;
    for i := 0; i < len(reduced.rows); i++ {
        arr := make([]float64, 1);
        for j := 0; j < len(arr); j++ {
            arr[j] = reduced.rows[i][augmentIndex];
        }
        solution[index] = arr;
        index += 1;
    }
    return NewMatrix(solution);
}
