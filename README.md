
## Avaible commands

```
# Configure AI provider and API key
go-commit config set-provider <provider-name> --api-key <your-api-key>

# Set detail level for commit messages
go-commit config set-detail [brief|normal|detailed]

# Generate commit from all staged files
go-commit commit generate

# Generate commit based on specific files (even if not staged)
go-commit commit analyze --files <file1> <file2>

# Generate commit from staged files but only analyzing specific files
go-commit commit generate --only <file1> <file2>

# Cherry-pick specific files to analyze (interactive mode)
go-commit commit pick
```