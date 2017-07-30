package src

/** Returns a matrix whose columns are the eigenvectors of m */
func eig(m *Matrix) (*Matrix, string) {
    if (m.N != m.M) {
        return nil, "ERROR: Matrix must be square";
    }
    return nil, "";
}
