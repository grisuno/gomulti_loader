# gomulti_loader
go run . -shellcode shellcode_linux.txt 

## Multi-Platform Shellcode Loader
## Author: grisuno
## Language: Go (with CGO)
## Platforms: Linux & Windows (64-bit)

# Overview
This project is a multi-platform shellcode loader written in Go using CGO to interface with native system calls. It supports both Linux and Windows operating systems and can execute raw shellcode from a file formatted with \x## byte notation (e.g., \x48\x31\xc0).

The loader reads shellcode from a text file, allocates executable memory, copies the shellcode into that memory, and executes it.

# Features
✅ Cross-platform support:
Linux: Uses mmap() for executable memory allocation.
Windows: Uses VirtualAlloc() with PAGE_EXECUTE_READWRITE.
📂 Flexible input: Parses shellcode in \x## format from plaintext files.
🔧 Simple CLI interface: Accepts path to shellcode file via command-line flag.
⚙️ Build automation: Makefile-style build instructions included for both platforms.
Supported Architectures
amd64 (x86_64) only
OS: Linux or Windows
Build Instructions
🔹 For Linux:
bash´´´
GOOS=linux GOARCH=amd64 go build -o loader_linux
´´´
🔹 For Windows:
Ensure you have mingw-w64 installed (e.g., x86_64-w64-mingw32-gcc):

bash´´´
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o loader_windows.exe
´´´
💡 Note: CGO is required for Windows due to use of kernel32.dll functions. 

🧹 Clean Build Artifacts:
bash´´´
rm -f loader_linux loader_windows.exe
´´´
Usage
bash'''
# On Linux
./loader_linux -shellcode /path/to/shellcode.txt

# On Windows
loader_windows.exe -shellcode C:\path\to\shellcode.txt
´´´
Example Shellcode File Format:
text'''
\x6a\x29\x58\x99\x6a\x02\x5f\x6a\x01\x5e\x0f\x05...
´´´
Any text file containing shellcode in \x## format is supported. The parser extracts all valid \x## sequences regardless of formatting or line breaks. 

How It Works
Reads the shellcode file.
Parses all \x## hex byte values into a byte array.
Allocates executable memory using OS-specific APIs:
Linux → mmap()
Windows → VirtualAlloc()
Copies shellcode into allocated memory.
Executes the shellcode by calling it as a function.
⚠️ Warning: This tool is intended for educational, research, or authorized security testing purposes only. 

## Example Shellcode Included
Two example payloads are embedded in the source:

Linux (64-bit): Reverse TCP shell (connects to IP:port)
Windows (64-bit): Reflective DLL injection / reverse meterpreter-style payload (via ws2_32.dll and cmd execution)
These are provided for testing and demonstration.

## Security Notes
This program requires no external dependencies beyond standard system libraries.
The use of mmap and VirtualAlloc with EXECUTE permissions may trigger AV/EDR detection.
Shellcode must be properly encoded and null-free depending on delivery context.
## Author
👤 grisuno
Security Researcher & LazyOwn Red Team Developer

Disclaimer
📌 This tool is designed for educational and ethical use only. Unauthorized use of this software to exploit systems without permission is illegal and unethical. The author assumes no liability for misuse.

Use responsibly and in compliance with all applicable laws and regulations.

License

GPLV3

![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)
