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

func main() {
    fmt.Println("Welcome to Calculator! Type 'exit' to exit. Type 'clear' to clear the screen.");
    reader := bufio.NewScanner(os.Stdin);
    fmt.Print(PROMPT);
    for reader.Scan() {
        line := reader.Text();
        if (line == EXIT) {
            break;
        } else if (line == "clear") {
            os.Stdout.WriteString(CLEAR);
        } else {
            fmt.Println(src.Read(line));
        }
        fmt.Print(PROMPT);
    }
    fmt.Println();
}
