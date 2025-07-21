# Security Considerations for cheat

## Overview

`cheat` is a local command-line tool designed to be run by individual users on their own systems. It is **not** designed to be exposed as a network service or to handle untrusted input from external sources.

## Threat Model

### In Scope
- Protecting users from accidental misuse
- Preventing scripts from using `cheat` to access unexpected files
- Basic input validation to ensure operations stay within designated directories

### Out of Scope
- Protection against malicious local users (they already have shell access)
- Network security (cheat has no network functionality)
- Protection against malicious cheatsheet content (users control their own content)

## Security Features

### 1. Path Traversal Protection

**Status**: ✅ Implemented

Cheatsheet names are validated to prevent directory traversal attacks:

```bash
# These are blocked:
cheat --edit "../../../etc/passwd"    # Parent directory traversal
cheat --edit "/etc/passwd"             # Absolute paths
cheat --edit "~/sensitive"             # Home directory expansion
cheat --edit ".gitignore"              # Hidden files (not displayed)
cheat --edit "config/.env"             # Nested hidden files
cheat --rm ".."                        # Dangerous removals
```

See [ADR-001](adr/001-path-traversal-protection.md) for implementation details.

### 2. Configuration Path Validation

**Status**: ✅ Implemented

The configuration system validates cheatpaths to ensure they:
- Have non-empty names
- Have non-empty paths
- Don't have duplicate names

### 3. Git Operations

**Status**: ⚠️ Limited Protection

When cloning community cheatsheets:
- Only clones from configured URLs
- Uses go-git library (not shell execution)
- No automatic execution of repository hooks

**Recommendation**: Only configure trusted repository URLs.

## Best Practices for Users

### 1. Cheatsheet Sources

- Only use cheatsheets from trusted sources
- Review cheatsheets before executing their commands
- Be cautious with cheatsheets containing:
  - Shell commands with variables
  - Sudo/administrative commands
  - Network operations
  - File system modifications

### 2. Configuration Security

- Store your config file with appropriate permissions (e.g., 600)
- Use read-only cheatpaths for community/shared content
- Keep personal cheatsheets in a separate, writable cheatpath

### 3. Editor Configuration

- Use a trusted editor in your configuration
- Avoid editors that execute initialization scripts from the current directory
- Consider using a minimal editor for cheat (e.g., `vim -u NONE`)

## Reporting Security Issues

While `cheat` is a local tool with limited security impact, we still take security seriously. If you discover a security issue:

1. **Do NOT** open a public issue
2. Instead, please report it to the maintainers privately
3. Include:
   - Description of the issue
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

## Security Non-Goals

The following are explicitly **not** security goals for cheat:

1. **Sandboxing**: Cheatsheets can contain any commands the user could run directly
2. **Multi-user isolation**: Each user has their own cheat configuration
3. **Audit logging**: No tracking of what cheatsheets are viewed or edited
4. **Encryption**: Cheatsheets are stored as plain text files
5. **Access control**: Relies on filesystem permissions

## Development Guidelines

When contributing to cheat:

1. **Input validation**: Validate user input early and clearly
2. **Fail safely**: Reject suspicious input rather than trying to sanitize
3. **Clear errors**: Provide helpful error messages that don't expose sensitive info
4. **Minimal privileges**: Don't require or use elevated privileges
5. **No shell execution**: Use Go libraries rather than shelling out when possible

## Version History

| Version | Security Changes |
|---------|-----------------|
| 4.4.3   | Added path traversal protection for cheatsheet names |
| 4.0.0   | Switched from Python to Go (memory safety) |