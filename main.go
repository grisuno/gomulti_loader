package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Definir flag para el archivo de shellcode
    shellcodeFile := flag.String("shellcode", "", "Path to shellcode file (hex format)")
    flag.Parse()
    
    // Validar que se proporcionó el archivo
    if *shellcodeFile == "" {
        fmt.Println("Error: -shellcode file path is required")
        fmt.Println("Usage: ./loader -shellcode /path/to/shellcode.txt")
        os.Exit(1)
    }
    
    executeLoader(*shellcodeFile)
}
