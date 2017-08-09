package src
/** Return 0 if vectors v and w are not orthogonal. Else return 1 */

func orth(v *Matrix, w *Matrix) (int, string) {
    dotproduct, e := dot(v, w);
    if (e != "") {
        return 0, "ERROR: Invalid parameters";
    }
    if (dotproduct == 0) {
        return 1, "";
    }
    return 0, "";
}
