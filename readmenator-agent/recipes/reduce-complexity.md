# Recipe: Reduce File Complexity

Target hotspot: `loader_windows.go`
(complexity 1.0, centrality 1.0)

1. Read dependents: `grep -n 'loader_windows.go' readmenator-agent/ARCHITECTURE.md`
2. Extract functions/classes into new files in the same subsystem
3. Update imports
4. Regenerate: `readmenator .`
