# Development TODOs - cheat

Last Updated: 2025-07-21

## Testing Infrastructure

1. **Test Coverage**
   - Set coverage targets (e.g., 80% minimum)
   - Add coverage reporting to CI/CD
   - Generate coverage badges or reports

2. **Integration Tests**
   - Add end-to-end tests for CLI commands
   - Test complete workflows (create, edit, search, etc.)

3. **Performance Tests**
   - Add benchmarks for sheet search operations
   - Profile performance with large cheatsheet collections
   - Benchmark regex compilation and caching

## Code Quality

1. **Refactor Large Functions**
   - Break down command functions (cmd_*.go) into smaller units
   - Extract common patterns into helper functions
   - Improve cyclomatic complexity

2. **Error Messages**
   - Make error messages more user-friendly
   - Add contextual help for common mistakes
   - Improve validation error descriptions

3. **Validation Consistency**
   - Standardize validation patterns across packages
   - Create common validation utilities

## Development Process

1. **Pre-commit Hooks**
   - Automate `make fmt`, `make lint`, `make vet`
   - Add test execution to pre-commit flow

2. **CI/CD Pipeline**
   - Make build configuration visible in repository
   - Add automated testing and quality checks
   - Consider GitHub Actions for releases

## Performance Optimizations

1. **Regex Optimization**
   - Cache compiled regex patterns
   - Optimize search operations for large collections

2. **Platform Code**
   - Better isolate OS-specific functionality
   - Reduce platform-specific branching
