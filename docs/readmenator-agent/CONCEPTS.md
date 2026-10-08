# Concepts

Nouns map atomically to file sets (EXTRACTED); verbs aggregate structural edges (INFERRED).

- `loader` | files=2 | mentions=6 | `loader_linux.go`, `loader_windows.go`
- `build` | files=2 | mentions=2 | `loader_linux.go`, `loader_windows.go`
- `execute` | files=2 | mentions=2 | `loader_linux.go`, `loader_windows.go`
- `file` | files=2 | mentions=2 | `loader_linux.go`, `loader_windows.go`
- `read` | files=2 | mentions=2 | `loader_linux.go`, `loader_windows.go`
- `shellcode` | files=2 | mentions=2 | `loader_linux.go`, `loader_windows.go`

## Dialectic

- Thesis: `build` centralizes 2 files; Antithesis: `execute` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `build` centralizes 2 files; Antithesis: `file` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `build` centralizes 2 files; Antithesis: `loader` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `build` centralizes 2 files; Antithesis: `read` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `build` centralizes 2 files; Antithesis: `shellcode` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `execute` centralizes 2 files; Antithesis: `file` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `execute` centralizes 2 files; Antithesis: `loader` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `execute` centralizes 2 files; Antithesis: `read` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `execute` centralizes 2 files; Antithesis: `shellcode` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `file` centralizes 2 files; Antithesis: `loader` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
