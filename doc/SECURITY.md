# Security Considerations

## Threat Model

The primary security concern for cheat is **malicious cheatsheets** crafted by attackers and distributed through shared repositories (like community cheatsheets). 

### Attack Vectors

1. **Malicious Cheatsheet Content**
   - An attacker creates a cheatsheet with specially crafted content
   - The cheatsheet is shared via community repositories or other channels
   - An unsuspecting user downloads and views the cheatsheet
   - Potential impacts:
     - Path traversal via sheet names
     - Malformed YAML frontmatter causing crashes
     - Regex DoS via search patterns
     - Tag injection/manipulation

2. **NOT in Scope: Self-Attack**
   - Users attacking themselves by creating malicious content
   - Users intentionally running dangerous commands from cheatsheets
   - These are not considered security issues as users have full control

### Security Measures

1. **Path Traversal Protection**
   - Sheet names are validated to prevent directory traversal
   - See `internal/cheatpath/ValidateSheetName` 

2. **Input Sanitization**
   - YAML frontmatter parsing handles malformed input gracefully
   - Search functions protect against regex DoS attacks
   - Tag filtering prevents injection attacks

3. **Fuzz Testing**
   - Comprehensive fuzz tests for all input parsing functions
   - Run with `make test-fuzz` (quick) or `make test-fuzz-long` (thorough)
   - Tests specifically crafted around the threat model

### Security Best Practices

1. **For Users:**
   - Only use cheatsheet repositories from trusted sources
   - Review cheatsheet content before executing commands
   - Be cautious with cheatsheets containing complex regex patterns

2. **For Contributors:**
   - All input parsing code must be fuzz tested
   - Path operations must use `ValidateSheetName`
   - Never execute cheatsheet content as code
   - Avoid string concatenation for building commands

### Reporting Security Issues

If you discover a security vulnerability, please report it to:
- Open an issue at https://github.com/cheat/cheat/issues
- For sensitive issues, contact the maintainers directly

### Security Testing

Run security-focused tests:
```bash
# Quick fuzz tests (1 minute)
make test-fuzz

# Thorough fuzz tests (40 minutes)  
make test-fuzz-long

# Run specific fuzz test
go test -fuzz=FuzzValidateSheetName -fuzztime=60s ./internal/cheatpath
```