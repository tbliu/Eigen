package src

/** Returns a diagonal matrix whose values along the diagonals are the eigenvalues of m 
* The diagonal matrix D is given by D = (P^-1)(A)(P)
*/
func diag(m *Matrix) (*Matrix, string) {
    P, err := eig(m);
    if (err != "") {
        return nil, err;
    }
    Pinv, _ := inv(P);
    return mulMatrices(mulMatrices(Pinv, m), P), "";
}
