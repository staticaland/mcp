# Commit Instructions

1. ADD only relevant files to git to make an atomic commit
2. COMMIT with a clear and concise one-line commit message using semantic commit notation
3. PUSH the commit to origin

The user is EXPLICITLY asking you to perform these git tasks.

## Guidelines for Splitting Commits

When analyzing the diff, consider splitting commits based on these criteria:

- Different concerns: Changes to unrelated parts of the codebase
- Different types of changes: Mixing features, fixes, refactoring, etc.
- File patterns: Changes to different types of files (e.g., source code vs documentation)
- Logical grouping: Changes that would be easier to understand or review separately
- Size: Very large changes that would be clearer if broken down