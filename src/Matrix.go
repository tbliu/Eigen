package src

import (
    "strings"
    "strconv"
)

/** Defines the matrix struct  */
type Matrix struct{
    M int; // defines number of rows (dimension in R^M)
    N int; // defines number of cols (dimension in R^N)
    rows [][]float64;
    cols [][]float64;
}

func NewMatrix(values [][]float64) (*Matrix, string) {
    //fmt.Println(values);
    if (!validate(values)) {
        return nil, "ERROR: Dimensions of matrix must match.";
    }
    m := new(Matrix);
    m.rows = values;
    m.cols = valsToCols(values);
    m.M = len(m.rows);
    m.N = len(m.cols);
    return m, "";
}

func validate(values [][]float64) bool {
    if (values == nil || len(values) <= 0) {
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

func queryToValues(query string) [][]float64 {
    query = strings.Replace(query, "[", "", -1);
    query = strings.Replace(query, "]", "", -1);
    rows := strings.Split(query, ";");
    values := make([][]float64, len(rows));
    for i := 0; i < len(values); i++ {
        row := rowToValues(rows[i]);
        if (row == nil) {
            return nil
        }
        values[i] = row;
    }
    return values;
}

func rowToValues(row string) []float64 {
    indivNums := strings.Split(row, ",");
    values := make([]float64, len(indivNums));
    for i := 0; i < len(values); i++ {
        val, err := strconv.ParseFloat(string(indivNums[i]), 64);
        if (err != nil) {
            return nil;
        }
        values[i] = val;
    }
    return values;
}

// Turns rows into columns
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
