package main

import (
    "bufio"
    "fmt"
    "os"
    "./src"
    //"strconv"
)

const EXIT = "exit";
const PROMPT = "> ";
const CLEAR = "\x1b[3;J\x1b[H\x1b[2J"; // Clears terminal screen
const CLEAN = "clean";

func main() {
    fmt.Println("Welcome to Eigen! Type 'exit' to exit.");
    reader := bufio.NewScanner(os.Stdin);
    fmt.Print(PROMPT);
    for reader.Scan() {
        line := reader.Text();
        if (line == EXIT) {
            break;
        } else if (line == "clear") {
            os.Stdout.WriteString(CLEAR);
        } else if (line == CLEAN) {
            src.Clean();
        } else {
            result := src.Transact(line);
            if (result == "") {
                fmt.Print(result)
            } else {
                fmt.Println(result);
            }
        }
        fmt.Print(PROMPT);
    }
    fmt.Println();
}
