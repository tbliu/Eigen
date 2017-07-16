package src

type Matrix struct{
    numRows int;
    numCols int;
    rows [][]float64;
    cols [][]float64;
}

func NewMatrix(values [][]float64) (*Matrix, string) {
    if (!validate(values)) {
        return nil, "ERROR: Dimensions of matrix must match.";
    }
    m := new(Matrix);
    m.rows = values;
    m.cols = valsToCols(values);
    m.numRows = len(m.rows);
    m.numCols = len(m.cols);
    return m, "";
}

func validate(values [][]float64) bool {
    if (len(values) <= 0) {
        return false;
    }
    numElements := len(values[0]);
    for i := 1; i < len(values); i++ {
        if (numElements != len(values[i])) {
            return false;
        }
    }
    return true;
}

func valsToCols(values [][]float64) [][]float64 {
    cols := make([][]float64, len(values[0]));
    for i := range cols {
        cols[i] = make([]float64, len(values));
    }
    for i := 0; i < len(values[0]); i++ {
        for j := 0; j < len(values); j++ {
            cols[i][j] = values[j][i];
        }
    }
    return cols;
}
