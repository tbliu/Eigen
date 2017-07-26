package src

/** Returns a matrix whose columns??? Rows??? span the rowspace of A */

func row(m *Matrix) (*Matrix, string) {
    trans, _ := transpose(m);
    return col(trans);
}
