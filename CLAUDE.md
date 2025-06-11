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