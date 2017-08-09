package src
/** Returns the vector V shifted by n digits
  * ex: roll([1;2;3;4], 3) --> [2;3;4;1]
*/

func roll(v *Matrix, n int) (*Matrix, string) {
    if (v.N != 1) {
        return nil, "ERROR: First parameter must be a matrix";
    }

    newCol := make([]float64, v.M);
    for i := 0; i < v.M; i++ {
        newIndex := (i + n) % v.M;
        newCol[newIndex] = v.rows[i][0];
    }
    val := make([][]float64, 1);
    val[0] = newCol;
    rolled, _ := NewMatrix(val);
    return transpose(rolled);
}
