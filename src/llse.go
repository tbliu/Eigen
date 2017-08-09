package src
/** Computes the Linear Least Squares Estimation of Ax=b 
* A is a matrix, b is a vector
*/

func llse(A *Matrix, b *Matrix) (*Matrix, string) {
    if (b.N != 1) {
        return nil, "ERROR: b must be a valid vector";
    }
    Atrans, e := transpose(A);
    if (e != "") {
        return nil, "ERROR: Invalid matrix";
    }

    inverse, e := inv(mulMatrices(Atrans, A));
    if (e != "") {
        return nil, "ERROR: Invalid matrix";
    }

    xhat := mulMatrices(inverse, mulMatrices(Atrans, b));
    round(xhat);
    return xhat, ""
}
