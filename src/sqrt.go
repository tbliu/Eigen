package src
/** Calculates the square root of a number */

import "math"

func sqrt(x float64) (float64, string) {
    res := math.Sqrt(x);
    if (math.IsNaN(res) || math.IsInf(res, 0)) {
        return 0.0, "ERROR: Cannot take the square root";
    }
    return float64(int(res * 1000)) / 1000, "";
}
