package src

func transpose(m *Matrix) (*Matrix, string) {
    rows := valsToCols(m.rows);
    matrix, err := NewMatrix(rows);
    return matrix, err;
}
