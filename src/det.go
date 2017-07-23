package src

func det(m *Matrix) (int, string) {
    if (m.N != m.M) {
        return 0, "ERROR: Matrix must be square";
    }
    return 0, ""; 
}
