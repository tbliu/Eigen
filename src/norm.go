package src
/** Returns the normalized version of a vector v */

func norm(v *Matrix) (*Matrix, string) {
    normsquared, e := dot(v, v);
    if (e != "") {
        return nil, e;
    }

    factor, e2 := sqrt(normsquared);
    if (e2 != "") {
        return nil, e2;
    }

    unitVector := scaleMatrix(v, 1/factor);
    round(unitVector);
    return unitVector, "";
}
