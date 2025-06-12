# MCP server project

Builds an MCP (Model Context Protocol) server using MCP-Go.

## Tech stack

- **MCP framework**: [MCP-Go](https://mcp-go.dev/servers) - Go implementation of MCP server
- **Repository**: https://github.com/mark3labs/mcp-go
- **Build tool**: [GoReleaser](https://goreleaser.com/) with GitHub Actions
- **Versioning**: [Release Please](https://github.com/googleapis/release-please-action) with googleapis/release-please-action
- **Package management**: GoReleaser config includes Homebrew formula generation
- **Task runner**: [Just](https://github.com/casey/just) for common commands

## Development workflow

### Build & release

- GoReleaser handles cross-platform builds via GitHub Actions
- Release Please automatically manages versioning and changelog generation
- Homebrew formula is automatically generated and updated

### Common commands

Use `just` to run common development tasks:

- Check `justfile` for available commands

## Resources

- [MCP-Go documentation](https://mcp-go.dev/servers)
- [MCP-Go GitHub](https://github.com/mark3labs/mcp-go)
- [GoReleaser documentation](https://goreleaser.com/)
- [Release Please documentation](https://github.com/google-github-actions/release-please-action)
- [Just documentation](https://github.com/casey/just)

## Commit standards

### Atomic commits

Break all code changes into small, atomic commits. Each commit must represent one logical change that makes sense on its own. If your commit message contains "and", split it into multiple commits.

### Conventional commits

All commit messages must follow the Conventional Commits format: `type: description`

Required commit types:

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `style:` - Code formatting
- `refactor:` - Code restructuring
- `test:` - Adding or updating tests
- `chore:` - Build process, tool changes

## Documentation standards

### Code formatting

In Markdown documentation, wrap all code-related items in backticks. This includes variable names, function names, file names, folder paths, configuration keys, command line tools, and API endpoints.

### Markdown formatting

After making changes to markdown files, run `prettier` for consistent formatting:

```bash
prettier --parser markdown --write <file>
```
