# Development TODOs - cheat

Last Updated: 2025-07-22

## High Priority Issues

1. **Fix gitdir.go Path Bounds Bug**
   - File: `internal/repo/gitdir.go`
   - Issue: `pos+5` could exceed string length causing panic
   - Impact: Potential crash when processing malformed git paths
   - Recommendation: Add bounds checking before slice access

## Recommendations for Next Steps

### Immediate (This Week)
1. **Fix the gitdir.go bug** - This is a crash bug that should be addressed immediately
2. **Add bounds checking** to prevent similar slice/string access panics:
   - `cmd_edit.go:112` - Check if `strings.Fields(conf.Editor)` returns non-empty slice
   - `main.go:82` & `cmd_init.go:44` - Check if `confpaths` slice is non-empty before accessing `[0]`

### Long Term (If Needed)
1. **Automated Release Pipeline**
   - Implement formal CI-based release mechanism using GoReleaser
   - Automate binary building, checksums, and GitHub releases
   - Add automated testing and quality gates
   - Note: This is a significant project requiring CI/CD infrastructure setup
