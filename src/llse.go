package src
/** Computes the Linear Least Squares Estimation of Ax=b 
* A is a matrix, b is a vector
*/

func llse(A *Matrix, b *Matrix) (*Matrix, string) {
    Atrans, e := transpose(A);
    if (e != "") {
        return nil, "ERROR: Invalid matrix";
    }

    inverse, e := inv(mulMatrices(Atrans, A));
    if (e != "") {
        return nil, "ERROR: Invalid matrix";
    }

    return mulMatrices(mulMatrices(inverse, Atrans), b), "";
}
