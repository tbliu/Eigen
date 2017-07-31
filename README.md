# Eigen
Eigen is a simple CLI that performs matrix and vector operations, modular arithmetic, and basic arithmetic. 
Eigen is currently a work in progress, so feel free to report any bugs or let me know if you have any suggestions for the program!

## Documentation
### Basic Commands
1. `exit` exits the interface.
2. `clear` clears the terminal window.
3. `clean` removes all variables you have declared so far. 

### Arithmetic
1. Use `+,-,/,*` to perform basic arithmetic in Eigen as you would normally.
  * Eigen supports arithmetic for floating point digits as well as integers
  * Whitespace does not matter: `2+2` is the same as `2 + 2`
  
### Variable assignment
1. Use `=` to assign a variable to a value.
2. Valid value types are: `int, float, matrix`
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
    * `[1, 2; -3, 1]` creates the matrix: <br />
    ![alt text](https://github.com/tbliu/Eigen/blob/master/assets/images/matrix_example_3x3.jpg.png "Example matrix");
3. You can also assign matrices to variables: 
    * `x = [1, 2; -3, 1]` creates the same matrix as above, and assigns it to the variable `x`.
    
### Matrix Operations
* col(A)
* det(A)
* id(N)
* inv(A)
* Lnull(A)
* null(A)
* nullity(A)
* rank(A)
* row(A)
* rref(A)
* solve(A, b)
* transpose(A)
* zeros(N), zeros(N, M)

### Modular arithmetic operations
Coming soon!

## Install
1. [Install Go](https://golang.org/dl/)
2. `$ git clone https://github.com/tbliu/Eigen`
3. `$ go run Main.go`
