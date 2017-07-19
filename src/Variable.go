package src

import (
    "strconv"
)

/** Defines variable structs */
type Variable struct {
    class string;
    intVal int;
    floatVal float64;
    matrix *Matrix;
}

func NewVariable(class string, intVal int, floatVal float64, matrix *Matrix) Variable {
    v := new(Variable);
    v.class = class;
    v.intVal = intVal;
    v.floatVal = floatVal;
    v.matrix = matrix;
    return *v;
}

func Get(v Variable) string {
    if (v.class == "matrix") {
        return Print(v.matrix);
    } else if (v.class == "int") {
        return strconv.Itoa(v.intVal);
    } else {
        return strconv.FormatFloat(v.floatVal, 'f', 2, 64);
    }
}

func ToString(v Variable) string {
    if (v.class == "int") {
        return strconv.Itoa(v.intVal);
    } else if (v.class == "float") {
        return strconv.FormatFloat(v.floatVal, 'f', 2, 64);
    } else {
        return Print(v.matrix);
    }
}

func Defined(name string) bool {
    _, matched := Variables[name];
    return matched;
}
