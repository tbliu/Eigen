package src

/** Returns a matrix B whose columns span the left null space of A */

func Lnull(m *Matrix) (*Matrix, string) {
    trans, _ := transpose(m);
    return null(trans);
}
