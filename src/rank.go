package src

/** Returns the rank (dimension of the column space) of A */
func rank(m *Matrix) int {
    col, _ := col(m);
    if (col == nil) {
        return 0;
    }
    count := 0;
    for i := 0; i < len(col.cols); i++ {
        count += 1;
    }
    return count;
}
