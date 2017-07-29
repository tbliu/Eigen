package src

/** Returns the rank (dimension of the column space) of A 
 !!!! THIS RETURNS AN INT NOT A MATRIX
*/

func rank(m *Matrix) int {
    col, _ := col(m);
    count := 0;
    for i := 0; i < len(col.cols); i++ {
        count += 1;
    }
    return count;
}
