# MCP Server Project

This project builds an MCP (Model Context Protocol) server using MCP-Go.

## Tech Stack

- **MCP Framework**: [MCP-Go](https://mcp-go.dev/servers) - Go implementation of MCP server
- **Repository**: https://github.com/mark3labs/mcp-go
- **Build Tool**: [GoReleaser](https://goreleaser.com/) with GitHub Actions
- **Versioning**: [Release Please](https://github.com/googleapis/release-please-action) with googleapis/release-please-action
- **Package Management**: GoReleaser config includes Homebrew formula generation
- **Task Runner**: [Just](https://github.com/casey/just) for common commands

## Development Workflow

### Build & Release

- GoReleaser handles cross-platform builds via GitHub Actions
- Release Please automatically manages versioning and changelog generation
- Homebrew formula is automatically generated and updated

### Common Commands

Use `just` to run common development tasks:

- Check `justfile` for available commands

## Resources

- [MCP-Go Documentation](https://mcp-go.dev/servers)
- [MCP-Go GitHub](https://github.com/mark3labs/mcp-go)
- [GoReleaser Documentation](https://goreleaser.com/)
- [Release Please Documentation](https://github.com/google-github-actions/release-please-action)
- [Just Documentation](https://github.com/casey/just)

## Commit Standards

### Atomic Commits

All code changes MUST be broken down into small, atomic commits. Each commit MUST represent one logical change that makes sense on its own. If your commit message contains "and", it SHOULD be split into multiple commits.

### Conventional Commits

All commit messages MUST follow the Conventional Commits format: `type: description`

Required commit types:

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `style:` - Code formatting
- `refactor:` - Code restructuring
- `test:` - Adding or updating tests
- `chore:` - Build process, tool changes

## Documentation Standards

### Code Formatting

In Markdown documentation, all code-related items MUST be wrapped in backticks. This includes variable names, function names, file names, folder paths, configuration keys, command line tools, and API endpoints.

### Markdown Formatting

After making changes to markdown files, MUST run `prettier` to ensure consistent formatting:

```bash
prettier --parser markdown --write <file>
```
