package src

/** Returns the dimension of the null space of a matrix */
func nullity(m *Matrix) int {
    return m.N - rank(m);
}
