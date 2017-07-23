package src

func transpose(m *Matrix) *Matrix {
    rows := valsToCols(m.rows);
    matrix, _ := NewMatrix(rows);
    return matrix;
}
