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

func main() {
    fmt.Println("Welcome to Calculator! Type 'exit' to exit");
    reader := bufio.NewScanner(os.Stdin);
    fmt.Print(PROMPT);
    //op := src.NewCalculator();
    for reader.Scan() {
        line := reader.Text();
        if (line == EXIT) {
            break;
        }
        src.Read(line);
        fmt.Println(src.Read(line));
        fmt.Print(PROMPT);
        line = reader.Text();
    }
}
