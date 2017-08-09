# Eigen
Eigen is a simple CLI that performs matrix and vector operations, modular arithmetic, and basic arithmetic. 
Eigen is currently a work in progress, so feel free to report any bugs or let me know if you have any suggestions for the program!

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
`solve(A *Matrix, b *Matrix)`        | Returns the vector x that solves the system of equations Ax=b  | Matrix (vector) 
`transpose(A *Matrix)`               | Returns the transpose of a matrix A                            | Matrix
`zeros(N)`                           | Returns the NxN zero matrix                                    | Matrix
`zeros(N, M)`                        | Returns the NxM zero matrix                                    | Matrix

### Modular arithmetic operations
Coming soon!

### Miscellaneous Operations
Command                              | Description                                                    | Return type
------------------------------------ | -------------                                                  | --------------
`gcd(x int, y int)`                  | Returns the greatest common divisor of x and y                 | int
`exit`                               | Exits the program                                              | None
`clear`                              | Clears the terminal window                                     | None
`clean`                              | Removes all variable assignments                               | None

## Install
1. [Install Go](https://golang.org/dl/)
2. `$ git clone https://github.com/tbliu/Eigen`
3. `$ cd Eigen`
4. `$ go run Main.go`
