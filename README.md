# Eigen
Eigen is a simple CLI that performs matrix and vector operations, modular arithmetic, and basic arithmetic. 
Eigen is currently a work in progress, so feel free to report any bugs or let me know if you have any suggestions for the program!

## Install
1. [Install Go](https://golang.org/dl/)
2. `$ git clone https://github.com/tbliu/Eigen`
3. `$ cd Eigen`
4. `$ go run Main.go`

## Documentation 

### Arithmetic
1. Use `+,-,/,*` to perform basic arithmetic in Eigen as you would normally.
  * Eigen supports arithmetic for floating point digits as well as integers
  
### Variable assignment
1. Use `=` to assign a variable to a value.
2. Valid value types are: `int, float64, *Matrix`
3. Variable names follow traditional rules:
  * The first character cannot be a number
  * All other characters can be either numbers or letters
  * Variables are case sensitive 
4. `exit, clear, clean` are all reserved key words and cannot be used as variables

### Matrices
1. To declare a matrix:
    1. Use `[` to denote the beginning of the matrix.
    2. Elements within a row are separated by a `,`
    3. Rows are separated by `;`
    4. Use `]` to denote the end of the end of the matrix.
2. Example:
    * `[1, 2, 1; 4, 3, 1; 0, 2, 1]` creates the matrix: <br />
    ![alt text](https://github.com/tbliu/Eigen/blob/master/assets/images/matrix_example_3x3.jpg.png "Example matrix");
3. You can also assign matrices to variables: 
    * `x = [1, 2, 1; 4, 3, 1; 0, 2, 1]` creates the same matrix as above, and assigns it to the variable `x`.

### Vectors
Vectors do not have a special definition in Eigen. They are simply defined as ℝ<sup>Nx1</sup> matrices. Any references to the word "Vector" or the parameter "\*Vector" should just be understood as a matrix with one column.
    
### Matrix Operations
Command                              | Description                                                    | Return type
------------------------------------ | -------------                                                  | --------------
`col(A *Matrix)`                     | Returns a matrix B whose columns span the column space of A    | Matrix
`det(A *Matrix)`                     | Returns the determinant of a square matrix A                   | float64
`id(N int)`                          | Returns the NxN identity matrix                                | Matrix
`inv(A *Matrix)`                     | Returns the inverse of a square matrix A                       | Matrix
`Lnull(A *Matrix)`                   | Returns a matrix B whose columns span the left null space of A | Matrix
`null(A *Matrix)`                    | Returns a matrix B whose columns span the null space of A      | Matrix
`nullity(A *Matrix)`                 | Returns the dimension of the null space of A                   | int
`rank(A *Matrix)`                    | Returns the dimension of the column space of A                 | int
`row(A *Matrix)`                     | Returns a matrix B whose columns span the row space of A       | Matrix
`rref(A *Matrix)`                    | Returns the reduced row echelon form of a matrix A             | Matrix
`solve(A *Matrix, b *Matrix)`        | Returns the vector x that solves the system of equations Ax=b  | Vector 
`transpose(A *Matrix)`               | Returns the transpose of a matrix A                            | Matrix
`zeros(N int)`                       | Returns the NxN zero matrix                                    | Matrix
`zeros(N int, M int)`                | Returns the NxM zero matrix                                    | Matrix
`dot(v *Vector, w *Vector)`          | Returns the dot products of vector v and vector w              | float64
`orth(v *Vector, w *Vector)`         | Returns 1 if v and w are orthogonal. 0 otherwise               | int
`llse(A *Matrix, b *Vector)`         | Returns the linear least squares estimate of `x` in `Ax=b`     | Vector
`proj(v *Vector, w *Vector)`         | Returns the projection of the vector w onto the vector v       | Vector
`gs(A *Matrix)`                      | Orthonormalizes the columns of A                               | Matrix
`qr(A *Matrix, tag string)`          | Performs QR decomposition on A. If tag == 'q', Q. If tag == 'r', return R. Quotes are not necessary. | Matrix 
`norm(v *Vector)`                    | Returns a vector w, which is the normalized version of v       | Vector
`roll(v *Vector, N int)`             | Returns a vector w, whose values are the same as v's, shifted by N places | Vector
`xcorr(v *Vector, w *Vector)`        | Returns the cross-correlation of w with respect to v           | Vector
`autocorr(v *Vector)`                | Returns the auto-correlation of v                              | Vector

### Modular arithmetic operations
Command                              | Description                                                    | Return type
------------------------------------ | -------------                                                  | --------------
`mod(a int, m int)`                  | Returns a modulo m                                             | int
`modexp(x int, y int, m int)`        | Returns x to the power of y modulo m. Uses repeated squaring, so this function works well with large ints. | int
`modinv(a int, m int)`               | Returns the inverse of a modulo m.                             | int

### Miscellaneous Operations
Command                              | Description                                                    | Return type
------------------------------------ | -------------                                                  | --------------
`gcd(x int, y int)`                  | Returns the greatest common divisor of x and y                 | int
`sqrt(x float64)`                    | Returns the square root of a number up to 3 decimal points     | float64
`exit`                               | Exits the program                                              | None
`clear`                              | Clears the terminal window                                     | None
`clean`                              | Removes all variable assignments                               | None

### List of Known Bugs 
1. Incomplete arithmetic queries (ex: `3*`) cause Eigen to panic (Index out of range error)
2. There are some bugs with matrix operations when not using variables (ex: [3,2] \* [3,1;1,2] returns an ERROR message. But assigning A=[3,2] and B = [3,1;1,2] and doing A\*B works).
3. Changing color to white after a red error message might make text invisible to people with white terminal backgrounds. Fix so it gets default terminal color
4. When subtracting with negative numbers, Eigen sometimes processes a negative number as a variable (ex: -3-4 returns "ERROR: Invalid variable '-3'")
