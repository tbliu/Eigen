package src
/** proj(v, w) returns the projection of w onto v */
func proj(v *Matrix, w *Matrix) (*Matrix, string) {
    num, e := dot(v, w);
    if (e != "") {
        return nil, e;
    }

    denom, e2 := dot(v, v);
    if (e2 != "") {
        return nil, e2;
    }
    if (denom == 0) {
        return nil, "ERROR: Projection is not defined";
    }

    return scaleMatrix(v, num/denom), "";
}
