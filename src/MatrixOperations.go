package src

import (
    "strconv"
    "strings"
    //"fmt"
)

/** Defines basic operations and helper functions for matrices */

/************** Begin parser for matrix **************/
func ApplyMatrixOperation(query string) (*Matrix, string) {
    first, operator, second, argErr := findArgs(query);
    if (argErr != "") {
        return nil, argErr;
    }
    f1, f := IsMatrix(first);
    s1, s := IsMatrix(second);
    if (f && s) {
        if (operator == "+") {
            x := addMatrices(f1, s1);
            return x, Print(x);
        } else if (operator == "~" || operator == "-") {
            x := subtractMatrices(f1, s1);
            return x, Print(x);
        } else if (operator == "*") {
            x := mulMatrices(f1, s1);
            return x, Print(x);
        } else {
            return nil, "ERROR: Cannot divide matrices";
        }
    } else if (f && !s) {
        secondValue, err := strconv.ParseFloat(second, 64);
        if (err != nil) {
            return nil, "ERROR: Malformed operand";
        }
        if (operator == "+") {
            x := addConstant(f1, secondValue);
            return x, Print(x);
        } else if (operator == "~" || operator == "-") {
            x := subtractConstant(f1, secondValue);
            return x, Print(x);
        } else {
            x := scaleMatrix(f1, secondValue);
            return x, Print(x);
        }
    } else if (!f && s) {
        firstValue, err := strconv.ParseFloat(first, 64);
        if (err != nil) {
            return nil, "ERROR: Malformed operand";
        }
        if (operator == "+") {
            x := addConstant(s1, firstValue);
            return x, Print(x);
        } else if (operator == "~" || operator == "-") {
            x := subtractConstant(s1, firstValue);
            return x, Print(x);
        } else {
            x := scaleMatrix(s1, firstValue);
            return x, Print(x);
        }
    }
    return nil, "ERROR: Malformed query";
}


func ApplyMultipleMatrixOperations(query string) (*Matrix, string) {
    firstOperator := strings.IndexAny(query, "+~/*");
    if (CountAny(query, "+", "~", "-", "*", "/") == 1) {
        return ApplyMatrixOperation(query);
    }
    firstArgument := query[0:firstOperator+1];
    query = query[firstOperator+1:len(query)];
    secondArgIndex := strings.IndexAny(query, "+~/*");
    secondArgument := query[0:secondArgIndex];
    query = query[secondArgIndex:len(query)];
    totalArgs := []string{firstArgument, secondArgument};
    mat, str := ApplyMatrixOperation(strings.Join(totalArgs, ""));
    if (mat == nil) {
        return mat, str;
    }
    newArgs := []string{matrixToString(mat), query};
    return ApplyMultipleMatrixOperations(strings.Join(newArgs, ""));
}

/************** Begin basic functionality **************/

func Print(m *Matrix) string {
    if (m == nil) {
        return "ERROR: Invalid matrix";
    }
    str := "";
    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            str += strconv.FormatFloat(m.rows[i][j], 'f', -1, 64);
            str += " ";
        }
        str += "\n";
    }
    return str;
}

// swaps row a and row b of a matrix
func swap(m *Matrix, a int, b int) {
    row1 := make([]float64, len(m.rows[a]));
    row2 := make([]float64, len(m.rows[b]));
    copy(row1, m.rows[a]);
    copy(row2, m.rows[b]);
    m.rows[b] = row1;
    m.rows[a] = row2;
}

// scales row i by a scale of 'scalar'
func scaleRow(row []float64, scalar float64) []float64 {
    newRow := make([]float64, len(row));
    for j := 0; j < len(newRow); j++ {
        newRow[j] = row[j] * scalar;
    }
    return newRow;
}

// adds row a and b together and puts the result into row c (which must be either a or b)
func addRows(a []float64, b []float64) []float64 {
    if (len(a) != len(b)) {
        return nil;
    }
    newRow := make([]float64, len(a));
    for i := 0; i < len(a); i++ {
        newRow[i] = a[i] + b[i];
    }
    return newRow;
}

func copyMatrix(m *Matrix) *Matrix {
    vals := make([][]float64, len(m.rows));
    for i := 0; i < len(vals); i++ {
        arr := make([]float64, len(m.rows[i]));
        copy(arr, m.rows[i]);
        vals[i] = arr;
    }
    n, _ := NewMatrix(vals);
    return n;
}

// scales every element in m by a factor of a
func scaleMatrix(m *Matrix, a float64) *Matrix {
    scaled := copyMatrix(m);
    for i := 0; i < len(scaled.rows); i++ {
        for j := 0; j < len(scaled.rows[i]); j++ {
            scaled.rows[i][j] *= a;
        }
    }
    scaled.cols = valsToCols(scaled.rows);
    return scaled;
}

// adds a to every element in m
func addConstant(m *Matrix, a float64) *Matrix {
    added := copyMatrix(m);
    for i := 0; i < len(added.rows); i++ {
        for j := 0; j < len(added.rows[i]); j++ {
            added.rows[i][j] += a;
        }
    }
    added.cols = valsToCols(added.rows);
    return added;
}

func subtractConstant(m *Matrix, a float64) *Matrix {
    return addConstant(m, -1*a);
}

// adds two matrices together
func addMatrices(m *Matrix, n *Matrix) *Matrix {
    if (m.M != n.M || m.N != n.N) { // matrix dimensions must match
        return nil;
    }
    toReturn := copyMatrix(m);
    for i := 0; i < len(toReturn.rows); i++ {
        for j := 0; j < len(toReturn.rows[i]); j++ {
            toReturn.rows[i][j] += n.rows[i][j];
        }
    }
    toReturn.cols = valsToCols(toReturn.rows);
    return toReturn;
}

func subtractMatrices(m *Matrix, n *Matrix) *Matrix {
    if (m.M != n.M || m.N != n.N) {
        return nil;
    }
    toReturn := copyMatrix(m);
    for i := 0; i < len(toReturn.rows); i++ {
        for j := 0; j < len(toReturn.rows[i]); j++ {
            toReturn.rows[i][j] -= n.rows[i][j];
        }
    }
    toReturn.cols = valsToCols(toReturn.rows);
    return toReturn;
}

// multiplies two matrices together
func mulMatrices(m *Matrix, n *Matrix) *Matrix {
    if (m.N != n.M) { // num cols of m must equal num rows of n
        return nil;
    }
    toReturn := zeros(m.M, n.N);
    for i := 0; i < len(toReturn.rows); i++ {
        for j := 0; j < len(toReturn.rows[i]); j++ {
            toReturn.rows[i][j] = computeEntry(m.rows[i], n.cols[j])
        }
    }
    toReturn.cols = valsToCols(toReturn.rows);
    return toReturn;
}

// Helper function for mulMatrices. Returns entry (i,j) in matrix
// by combining the ith row in m and the jth column in n
func computeEntry(i []float64, j []float64) float64 {
    entry := 0.0;
    for k := 0; k < len(i); k++ {
        entry += i[k] * j[k];
    }
    return entry;
}
