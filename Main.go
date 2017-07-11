package main

import (
    "bufio"
    "fmt"
    "os"
)

const EXIT = "exit";
const PROMPT = "> ";

func main() {
    fmt.Println("Welcome to Calculator! Type 'exit' to exit");
    reader := bufio.NewScanner(os.Stdin);
    fmt.Print(PROMPT);

    for reader.Scan() {
        line := reader.Text();
        if (line == EXIT) {
            break;
        }

        fmt.Print(PROMPT);
        line = reader.Text();
    }
}
