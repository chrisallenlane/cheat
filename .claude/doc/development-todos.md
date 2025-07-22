# Development TODOs - cheat

Last Updated: 2025-07-22

## High Priority Issues

1. **Fix gitdir.go Path Bounds Bug**
   - File: `internal/repo/gitdir.go`
   - Issue: `pos+5` could exceed string length causing panic
   - Impact: Potential crash when processing malformed git paths
   - Recommendation: Add bounds checking before slice access

## Code Quality

1. **Error Messages**
   - Make error messages more user-friendly
   - Add contextual help for common mistakes
   - Improve validation error descriptions

2. **Validation Consistency**
   - Standardize validation patterns across packages
   - Create common validation utilities

3. **Fuzz Testing**
   - Evaluate whether the project would benefit from fuzz testing
   - Consider fuzzing sheet parsing, search regex handling, and path validation
   - Could help discover edge cases in input handling

## Performance Optimizations

1. **Search Performance**
   - Consider parallelizing search across multiple cheatsheets
   - Profile with large collections (1000+ sheets)
   - Only if users report performance issues

## Recommendations for Next Steps

### Immediate (This Week)
1. **Fix the gitdir.go bug** - This is a crash bug that should be addressed immediately
2. **Add bounds checking** to prevent similar slice/string access panics:
   - `cmd_edit.go:112` - Check if `strings.Fields(conf.Editor)` returns non-empty slice
   - `main.go:82` & `cmd_init.go:44` - Check if `confpaths` slice is non-empty before accessing `[0]`

### Long Term (If Needed)
1. **Performance optimizations**
   - Only if performance becomes a user-reported issue
   - Parallelize search operations for large cheatsheet collections

2. **Automated Release Pipeline**
   - Implement formal CI-based release mechanism using GoReleaser
   - Automate binary building, checksums, and GitHub releases
   - Add automated testing and quality gates
   - Note: This is a significant project requiring CI/CD infrastructure setup

## Notes
- Focus on stability and reliability over new features
- Address crash bugs and defensive programming issues
