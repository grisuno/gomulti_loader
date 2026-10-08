# Concepts

Second-brain semantic layer: nouns map atomically to file sets (EXTRACTED); verbs aggregate structural edges (INFERRED).

| Concept | Files | Mentions | Top Files |
|---------|-------|----------|-----------|
| `loader` | 2 | 6 | `loader_linux.go`, `loader_windows.go` |
| `build` | 2 | 2 | `loader_linux.go`, `loader_windows.go` |
| `execute` | 2 | 2 | `loader_linux.go`, `loader_windows.go` |
| `file` | 2 | 2 | `loader_linux.go`, `loader_windows.go` |
| `read` | 2 | 2 | `loader_linux.go`, `loader_windows.go` |
| `shellcode` | 2 | 2 | `loader_linux.go`, `loader_windows.go` |

## Dialectic Prompts

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
