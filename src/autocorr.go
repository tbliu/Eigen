package src

/** Returns the autocorrelation of Vector x */

func autocorr(x *Matrix) (*Matrix, string) {
    return xcorr(x,x);
}
