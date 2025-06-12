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

## Commit Structure Requirements

### Atomic Commit Composition

**REQ-001**: Developers MUST decompose all changes into atomic commits where each commit represents the smallest meaningful unit of change that maintains system integrity.

**REQ-002**: Each commit MUST be self-contained such that the codebase remains in a functional state after applying the commit in isolation.

**REQ-003**: Developers MUST NOT combine logically distinct changes within a single commit, regardless of their perceived relationship or temporal proximity.

**REQ-004**: When a proposed commit message contains conjunctions (specifically "and", "or", "also", "plus"), the developer MUST evaluate whether the commit violates atomicity requirements and decompose accordingly.

### Commit Message Standards

**REQ-005**: All commit messages MUST conform to the Conventional Commits specification version 1.0.0 as published at https://www.conventionalcommits.org/.

**REQ-006**: Commit messages MUST follow the structure: `<type>[optional scope]: <description>` where:
- `type` is drawn from the approved type vocabulary
- `scope` is optional and project-specific
- `description` is a concise summary in imperative mood

**REQ-007**: Commit types MUST be selected from the following approved vocabulary:
- `feat`: new feature for the user
- `fix`: bug fix for the user
- `docs`: changes to documentation
- `style`: formatting changes that do not affect code meaning
- `refactor`: code change that neither fixes a bug nor adds a feature
- `test`: adding missing tests or correcting existing tests
- `chore`: changes to build process or auxiliary tools

## Documentation Formatting Requirements

### Code Reference Markup

**REQ-008**: Within Markdown documentation, all code-related elements MUST be enclosed in backticks (`) to apply monospace formatting.

**REQ-009**: Code-related elements subject to markup requirements include but are not limited to:
- Variable names
- Function names
- Class names
- File names and paths
- Configuration keys
- Command line arguments
- API endpoints
- Database table and column names

**REQ-010**: Developers SHOULD use triple backticks (```) with language specification for multi-line code blocks to enable syntax highlighting.
