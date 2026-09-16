# Subsystem: root

## app.py
- Layer: utility
- Doc: _*_ coding: utf8 _*_
- Language: py

## install.sh
- Layer: utility
- Language: sh

## loader_linux.go
- Layer: utility
- Doc: go:build linux
- Language: go
- Symbols:
  - `executeLoader` (function, line 29) `func executeLoader(`
  - `readShellcodeFromFile` (function, line 49) `func readShellcodeFromFile(`

## loader_windows.go
- Layer: utility
- Doc: go:build windows
- Language: go
- Symbols:
  - `executeLoader` (function, line 29) `func executeLoader(`
  - `readShellcodeFromFile` (function, line 49) `func readShellcodeFromFile(`

## main.go
- Layer: utility
- Language: go
- Symbols:
  - `main` (function, line 9) `func main(`
