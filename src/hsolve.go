package src

/** Returns the vector that solves the homogeneous equation Ax=0 */
func hsolve(m *Matrix) (*Matrix, string) {
    aug, _ := zeros(m.M, 1);
    return solve(m, aug);
}
