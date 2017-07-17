package src

import (
    "strconv"
)

/** Defines operations for matrices */

func Print(m *Matrix) string {
    str := "";
    for i := 0; i < len(m.rows); i++ {
        for j := 0; j < len(m.rows[i]); j++ {
            str += strconv.FormatFloat(m.rows[i][j], 'f', 1, 64);
            str += " ";
        }
        str += "\n";
    }
    return str;
}
