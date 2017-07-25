package src

/* Creates the NxN identity matrix */

func id(N int) (*Matrix, string) {
    vals := make([][]float64, N);
    for i := 0; i < len(vals); i++ {
        arr := make([]float64, N);
        arr[i] = 1;
        vals[i] = arr;
    }
    matrix, err := NewMatrix(vals);
    if (err == "") {
        return matrix, "";
    }
    return nil, err;
}
