package src

/** Computes the determinant of a matrix 
 *  Matrix must be square.
*/

func det(m *Matrix) (int, string) {
    if (m.N != m.M) {
        return 0, "ERROR: Matrix must be square";
    }
    return 0, "";
}
