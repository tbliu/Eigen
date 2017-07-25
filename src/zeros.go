package src

// creates a matrix of size dim with all zeros. 
// If len(dim) == 1, create a square matrix of zeros
// Else create an MxN matrix
func zeros(dim ...int) (*Matrix, string) {
    if (len(dim) > 2 || len(dim) <= 0) {
        return nil, "ERROR: The number of parameters must be either one or two";
    }
    var rows int;
    var cols int;
    if (len(dim) == 1) {
        rows = dim[0];
        cols = dim[0];
    } else {
        rows = dim[0];
        cols = dim[1];
    }
    values := make([][]float64, rows);
    for i := 0; i < len(values); i++ {
        values[i] = make([]float64, cols);
    }
    mat, err := NewMatrix(values);
    if (err != "") {
        return nil, err;
    }
    return mat, "";
}
