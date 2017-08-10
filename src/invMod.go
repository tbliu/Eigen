package src

/** Finds the inverse a number A modulo B 
* Reference: http://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
*/

func modinv(A int, B int) (int, string) {
    divisor, _ := gcd(A, B);
    if (divisor != 1) {
        return 0, "ERROR: There is no unique inverse";
    }
    m := B // original mod
    q := 0;
    t := 0;
    x0 := 0;
    x1 := 1;
    for ; A > 1; {
        q = A / B;
        t = B;
        B = A % B;
        A = t;
        t = x0;
        x0 = x1 - q * x0;
        x1 = t;
    }

    for ; x1 < 0; {
        x1 += m;
    }
    return x1, "";
}
